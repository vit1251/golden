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

	type commandName struct {
		id CommandID
		name string
	}

	commandNames := []commandName{
		{M_NUL, "M_NUL"},
		{M_ADR, "M_ADR"},
		{M_PWD, "M_PWD"},
		{M_FILE, "M_FILE"},
		{M_OK, "M_OK"},
		{M_EOB, "M_EOB"},
		{M_GOT, "M_GOT"},
		{M_GET,"M_GET"},
	}

	for _, commandName := range commandNames {
		if commandName.id == self {
			return commandName.name
		}
	}

	return fmt.Sprintf("M_UNKOWN[%d]", self)

}
