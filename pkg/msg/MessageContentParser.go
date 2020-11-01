package msg

import (
	"bytes"
	"github.com/vit1251/golden/pkg/packet"
)

type MessageContentParser struct {
}

func NewMessageContentParser() *MessageContentParser {
	return new(MessageContentParser)
}

const (
	CR = "\x0D"
	LF = "\x0A"
)

func (self MessageContentParser) Parse(content []byte) (*MessageContent, error) {

	mc := NewMessageContent()

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
			k := packet.NewKludge()
			k.Set(row)
			mc.AddKludge(*k)
		}
		if bytes.HasPrefix(row, []byte{' ', '*', ' ', 'O', 'r', 'i', 'g', 'i', 'n', ':'}) {
			msgBody = false
		}
	}

	return mc, nil

}
