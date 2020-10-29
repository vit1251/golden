package bbs

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"go.uber.org/dig"
	"log"
)

type ConnState struct {
	t          *Terminal      /* Terminal       */
	scr        *Screen        /* Screen         */
	activeView Widget         /* Area widget    */
	container  *dig.Container /* DI container   */
	activeArea string         /* Active area    */
}

func NewConnState(c *dig.Container) *ConnState {
	cs := new(ConnState)
	cs.container = c
	return cs
}

func (cs *ConnState) SetSession(sess ssh.Session) {
	cs.activeView = NewWelcomeView()
	cs.t = NewTerminal(sess)
	cs.scr = NewScreen(cs.t)
}

func (cs *ConnState) ProcessConnection() {
	log.Printf("Session start.")
	var processing bool = true
	for processing {
		select {

		case rr := <-cs.t.remoteRedraw:
			fmt.Printf("Redraw: rr = %v", rr)
			cs.activeView.Render(cs)
			break

		case winEvent := <-cs.t.winEvent:
			fmt.Printf("change window: %v\n", winEvent)
			cs.t.Width = winEvent.Width
			cs.t.Height = winEvent.Height
			cs.t.Redraw()
			break

		case event := <-cs.t.Event:
			if event.Type == TerminalError {
				processing = false
			} else {
				cs.activeView.ProcessEvent(cs, &event)
				cs.t.Redraw()
			}
			break

		}
	}
	log.Printf("Session complete.")
}
