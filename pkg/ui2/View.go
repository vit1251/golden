package ui2

type View interface {
	Render(cs *ConnState)
	ProcessEvent(cs *ConnState, event *TerminalEvent)
}
