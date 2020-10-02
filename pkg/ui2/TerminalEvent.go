package ui2

const (
	TerminalKey   = 1
	TerminalError = 2
)

type TerminalEvent struct {
	Type   int
	Key    string
}

func NewTerminalEvent() *TerminalEvent {
	return new(TerminalEvent)
}
