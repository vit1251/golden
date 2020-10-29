package stream

import "fmt"

type CommandID byte

const (
	M_NUL  CommandID = 0
	M_ADR  CommandID = 1
	M_PWD  CommandID = 2
	M_FILE CommandID = 3
	M_OK   CommandID = 4
	M_EOB  CommandID = 5
	M_GOT  CommandID = 6
	M_ERR  CommandID = 7
	M_BSY  CommandID = 8
	M_GET  CommandID = 9
	M_SKIP CommandID = 10
)

func (self CommandID) String() string {
	switch self {
	case M_NUL:
		return "M_NUL"
	case M_ADR:
		return "M_ADR"
	case M_PWD:
		return "M_PWD"
	case M_FILE:
		return "M_FILE"
	case M_OK:
		return "M_OK"
	case M_EOB:
		return "M_EOB"
	}
	return fmt.Sprintf("M_UNKOWN[%d]", self)
}