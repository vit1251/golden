package ui2

type Screen struct {
	t *Terminal
}

func NewScreen(t *Terminal) *Screen {
	s := new(Screen)
	s.t = t
	return s
}

func (s *Screen) WriteString(str string) {
	s.t.Write([]byte(str))
}

func (s *Screen) WriteStringXY(x int, y int, str string) {
	s.t.GotoXY(x, y)
	s.t.Write([]byte(str))
}

func (s *Screen) DrawLine(str string) {

	for i := 0; i < s.t.Width; i++ {
		s.t.Write([]byte(str))
	}

}

func (s *Screen) DrawLineY(y int, patt string) {
	s.t.GotoXY(0, y)
	var row string
	for i := 0; i < s.t.Width; i++ {
		row += patt
	}
	s.t.Write([]byte(row))
}
