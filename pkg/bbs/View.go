package bbs

type View interface {
	Render(cs *ConnState)
	ProcessEvent(cs *ConnState, event *TerminalEvent)
}
