package ui2

type Widget interface {
	Render(cs *ConnState)
	ProcessEvent(cs *ConnState, event *TerminalEvent)
}
