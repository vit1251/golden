package ui2

import "fmt"

type TextWidget struct {
	lines []string
}

func NewTextWidget() *TextWidget {
	return new(TextWidget)
}

func (tw *TextWidget) AddLine(s string) {
	tw.lines = append(tw.lines, s)
}

func (tw *TextWidget) Render(s *Screen) {
	for _, l := range tw.lines {
		s.WriteString(fmt.Sprintf("%s\n", l))
	}
}

func (tw *TextWidget) RenderXY(x int, y int, s *Screen) {
	for i, l := range tw.lines {
		s.WriteStringXY(x, y+i, l)
	}
}
