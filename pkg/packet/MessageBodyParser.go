package packet

import "bytes"

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

	mc := NewMessageBody()

	/* Remove "soft" linefeed */
	parts := bytes.Split(content, []byte(LF))
	newContent := bytes.Join(parts, []byte{})

	/* Split by "hard" line split */
	rows := bytes.Split(newContent, []byte(CR))

	/* Parse AREA */
	if len(rows) > 0 {

		row := rows[0]

		if bytes.HasPrefix(row, []byte{'A', 'R', 'E', 'A', ':'}) {

			/* Set AREA value */
			areaName := string(row[5:])
			mc.SetArea(areaName)

			/* Remove AREA */
			rows = rows[1:]
		}

	}

	/* Process message body */
	var msgBody bool = true
	for _, row := range rows {
		if msgBody && !bytes.HasPrefix(row, []byte{'\x01'}) {
			mc.AddLine(row)
		}
		if bytes.HasPrefix(row, []byte{'\x01'}) {
			k := NewKludge()
			k.Set(row)
			mc.AddKludge(*k)
		}
		if bytes.HasPrefix(row, []byte{' ', '*', ' ', 'O', 'r', 'i', 'g', 'i', 'n', ':'}) {
			msgBody = false
		}
	}

	return mc, nil

}