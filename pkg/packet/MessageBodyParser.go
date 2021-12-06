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

const (
	AREA_KLUDGE = "AREA:"
)

func (self MessageBodyParser) parseArea(messageBody *MessageBody, rows [][]byte) {
	rowCount := len(rows)
	if rowCount > 0 {
		startLine := rows[0]
		if bytes.HasPrefix(startLine, []byte(AREA_KLUDGE)) {
			/* Set AREA value */
			areaStartIndex := len(AREA_KLUDGE)
			areaName := startLine[areaStartIndex:]
			newAreaName := string(areaName)
			messageBody.SetArea(newAreaName)
		}
	}
}

const (
	MSG_BODY_PARSE_BODY         = 1
	MSG_BODY_PARSE_UUE          = 2
	MSG_BODY_PARSE_ROUTE_INFO   = 3
)

func checkNumber(str []byte) bool {
	var byteCount = len(str)
	var digitCount = 0
	for _, b := range str {
		if b == '0' || b == '1' || b == '2' || b == '3' || b == '4' {
			digitCount = digitCount + 1
		}
		if b == '5' || b == '6' || b == '7' || b == '8' || b == '9' {
			digitCount = digitCount + 1
		}
	}
	return byteCount == digitCount
}

func (self MessageBodyParser) Parse(content []byte) (*MessageBody, error) {

	messageBody := NewMessageBody()

	/* Remove "soft" linefeed */
	parts := bytes.Split(content, []byte(LF))
	newContent := bytes.Join(parts, []byte{})

	/* Save RAW packet */
	messageBody.SetPacket(newContent)

	/* Split by "hard" line split */
	rows := bytes.Split(newContent, []byte(CR))

	/* Parse AREA name */
	self.parseArea(messageBody, rows)

	/* Process message body */
	var attach *MessageBodyAttach
	var parserMode = MSG_BODY_PARSE_BODY

	for _, row := range rows {

		log.Printf("Parse: %d %s", parserMode, row)

		if parserMode == MSG_BODY_PARSE_BODY {
			if bytes.HasPrefix(row, []byte{'\x01'}) {
				k := NewKludge()
				k.Set(row)
				messageBody.AddKludge(*k)
			} else if bytes.HasPrefix(row, []byte("begin")) {
				parts := bytes.Split(row, []byte{' '})
				partCount := len(parts)
				if partCount == 3 && checkNumber(parts[1]) {
                                        //
					attach = NewMessageBodyAttach()
					attach.SetPermission(string(parts[1]))
					attach.SetName(string(parts[2]))
					//
					parserMode = MSG_BODY_PARSE_UUE
				} else {
					messageBody.AddLine(row)
				}
			} else if bytes.HasPrefix(row, []byte(" * Origin:")) {
				parserMode = MSG_BODY_PARSE_ROUTE_INFO
				messageBody.AddLine(row)
				messageBody.SetOrigin(row[10:])
			} else {
				messageBody.AddLine(row)
			}

		} else if parserMode == MSG_BODY_PARSE_UUE {
			if  bytes.HasPrefix(row, []byte("end")) {
				size := attach.Len()
				messageBody.AddAttach(*attach)
				attach = nil
				messageBody.AddLine([]byte(fmt.Sprintf(" --- UUE ( size = %d ) --- ", size)))
				parserMode = MSG_BODY_PARSE_BODY
			} else {
				attach.WriteLine(row)
			}
		} else {
			log.Printf("MessageBodyParser: Parse: wrong state: row = %s", row)
		}
	}

	return messageBody, nil

}
