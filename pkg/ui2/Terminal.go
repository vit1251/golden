package ui2

import (
	"bytes"
	"fmt"
	"github.com/gliderlabs/ssh"
	"io"
)

//Format text
const (
	RESET       = 0
	BRIGHT 		= 1
	DIM			= 2
	UNDERSCORE	= 3
	BLINK		= 4
	REVERSE		= 5
	HIDDEN		= 6
)

//Foreground Colours (text)
const (
	F_BLACK 	= 30
	F_RED		= 31
	F_GREEN		= 32
	F_YELLOW	= 33
	F_BLUE		= 34
	F_MAGENTA 	= 35
	F_CYAN		= 36
	F_WHITE		= 37
)

//Background Colours
const (
	B_BLACK 	= 40
	B_RED		= 41
	B_GREEN		= 42
	B_YELLOW	= 44
	B_BLUE		= 44
	B_MAGENTA 	= 45
	B_CYAN		= 46
	B_WHITE		= 47
)

type TerminalRedraw bool

type Terminal struct {
	sess           ssh.Session
	Event          chan TerminalEvent
	winEvent       <-chan ssh.Window
	remoteRedraw   chan TerminalRedraw
	pty            ssh.Pty
	hasPty         bool
	Width          int
	Height         int
}

func NewTerminal(sess ssh.Session) *Terminal {
	t := new(Terminal)
	t.sess = sess
	t.Width = 80
	t.Height = 25
	t.Event = make(chan TerminalEvent)
	t.remoteRedraw = make(chan TerminalRedraw)
	t.pty, t.winEvent, t.hasPty = sess.Pty()
	go t.processEvent()
	return t
}



func (t *Terminal) makeKeySym(in []byte) string {
	fmt.Printf("size = %d data = %v\n", len(in), in)
	if bytes.Equal(in, []byte{13}) {
		return "ENTER"
	} else if bytes.Equal(in, []byte{27, 91, 50, 126}) {
		return "INSERT"
	} else if bytes.Equal(in, []byte{27, 91, 51, 126}) {
		return "DELETE"
	} else if bytes.Equal(in, []byte{27, 91, 68}) {
		return "LEFT"
	} else if bytes.Equal(in, []byte{27, 91, 67}) {
		return "RIGHT"
	} else if bytes.Equal(in, []byte{27, 91, 72}) {
		return "HOME"
	} else if bytes.Equal(in, []byte{27, 91, 70}) {
		return "END"
	} else if bytes.Equal(in, []byte{27, 91, 65}) {
		return "UP"
	} else if bytes.Equal(in, []byte{27, 91, 66}) {
		return "DOWN"
	} else if bytes.Equal(in, []byte{27, 91, 54, 126}) {
		return "PG_DOWN"
	} else if bytes.Equal(in, []byte{27, 91, 53, 126}) {
		return "PG_UP"
	} else if bytes.Equal(in, []byte{27}) {
		return "ESC"
	} else {
		return "unknown"
	}
}

func (t *Terminal) Redraw() {
	go func() {
		msg := new(TerminalRedraw)
		t.remoteRedraw <- *msg
	}()
}

func (t *Terminal) processEvent() {

	data := make([]byte, 64)

	for {
		size, err := t.sess.Read(data)
		if err == nil {
			fmt.Printf("size = %d data = %v\n", size, data)
			te := NewTerminalEvent()
			te.Type = TerminalKey
			te.Key = t.makeKeySym(data[:size])
			t.Event <- *te
		} else {
			fmt.Printf("err = %v\n", err)
			te := NewTerminalEvent()
			te.Type = TerminalError
			t.Event <- *te
			break
		}
	}

}

func (t *Terminal) Write(data []byte) error {
	_, err := t.sess.Write(data)
	return err
}

func (t *Terminal) GetSize() (int, int) {
	return t.Width, t.Height
}

func (t *Terminal) cursorhome() {
	io.WriteString(t.sess, fmt.Sprintf("\x1B[H"))
}

// Clear screen from cursor down
func (t *Terminal) cleareos() {
	io.WriteString(t.sess, "\x1B[J")
}

// Clear entire screen
func (t *Terminal) ED2() {
	io.WriteString(t.sess, "\x1B[2J")
}

func (t *Terminal) GotoXY(x int, y int) {
	io.WriteString(t.sess, fmt.Sprintf("\x1B[%d;%dH", y, x))
}

func (t *Terminal) ResetAttr() {
	io.WriteString(t.sess, "\x1B[0m")
}

func (t *Terminal) SetAttr(attr int) {
	io.WriteString(t.sess, fmt.Sprintf("\x1B[%dm", attr))
}
