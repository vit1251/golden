package packet

import (
	"bytes"
	"fmt"
	"log"
)

type MessageBodyParser struct {
}

func NewMessageBodyParser() *MessageBodyParser {
	return new(MessageBodyParser)
}

const (
//	CR = "\x0D"
	LF = "\x0A"
)

func (self MessageBodyParser) Parse(content []byte) (*MessageBody, error) {

	messageBody := NewMessageBody()

	/* Remove "soft" linefeed */
	parts := bytes.Split(content, []byte(LF))
	newContent := bytes.Join(parts, []byte{})

	/* Save RAW packet */
	messageBody.SetPacket(newContent)

	/* Split by "hard" line split */
	rows := bytes.Split(newContent, []byte(CR))

	/* Parse AREA */
	if len(rows) > 0 {

		row := rows[0]

		if bytes.HasPrefix(row, []byte{'A', 'R', 'E', 'A', ':'}) {

			/* Set AREA value */
			areaName := string(row[5:])
			messageBody.SetArea(areaName)

			/* Remove AREA */
			rows = rows[1:]
		}

	}

	/* Process message body */
	var msgBody bool = true
	var msgUUE bool = false
	var attach *MessageBodyAttach

	for _, row := range rows {

		if bytes.HasPrefix(row, []byte{'\x01'}) {
			/* Process control paragraph */
			k := NewKludge()
			k.Set(row)
			messageBody.AddKludge(*k)
		} else if msgUUE {

			if  bytes.HasPrefix(row, []byte("end")) {
				//
				size := attach.Len()
				messageBody.AddAttach(*attach)
				attach = nil
				//
				messageBody.AddLine([]byte(fmt.Sprintf(" --- UUE ( size = %d ) --- ", size)))
				//
				msgUUE = false
				//
			} else {
				attach.WriteLine(row)
			}

		} else if msgBody {

			if bytes.HasPrefix(row, []byte{'b', 'e', 'g', 'i', 'n'}) {
				//
				parts := bytes.Split(row, []byte{' '})
				if len(parts) == 3 {
					msgUUE = true
					attach = NewMessageBodyAttach()
					attach.SetPermission(string(parts[1]))
					attach.SetName(string(parts[2]))
				} else {
					/* Process message body */
					messageBody.AddLine(row)
				}

			} else if bytes.HasPrefix(row, []byte{' ', '*', ' ', 'O', 'r', 'i', 'g', 'i', 'n', ':'}) {
				messageBody.AddLine(row)
				messageBody.SetOrigin(row[10:])
				msgBody = false
				msgUUE = false
			} else {
				/* Process message body */
				messageBody.AddLine(row)
			}

		} else {
			log.Printf("MessageBodyParser: Parse: wrong state: row = %s", row)
		}

	}

	return messageBody, nil

}