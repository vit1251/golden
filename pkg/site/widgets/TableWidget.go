package widgets

import (
	"fmt"
	"io"
)

type TableHeaderWidget struct {
}

type TableCelWidget struct {
	Widget  IWidget
	Content string
	class   string
}

type TableRowWidget struct {
	class string
	Cells []*TableCelWidget
	title string
}

func NewTableCellWidget() *TableCelWidget {
	cw := new(TableCelWidget)
	return cw
}

func (self *TableCelWidget) SetWidget(w IWidget) *TableCelWidget {
	self.Widget = w
	return self
}

func (self *TableCelWidget) SetClass(s string) *TableCelWidget {
	self.class = s
	return self
}

func (self *TableRowWidget) AddCell(cell *TableCelWidget) *TableRowWidget {
	self.Cells = append(self.Cells, cell)
	return self
}

func (self *TableRowWidget) SetClass(s string) *TableRowWidget {
	self.class = s
	return self
}

func (self *TableRowWidget) SetTitle(title string) *TableRowWidget {
	self.title = title
	return self
}

type TableWidget struct {
	headers TableHeaderWidget
	rows    []*TableRowWidget
	class   string
}

func NewTableWidget() *TableWidget {
	w := new(TableWidget)
	return w
}

func NewTableRowWidget() *TableRowWidget {
	row := new(TableRowWidget)
	return row
}

func (self *TableWidget) Render(w io.Writer) error {
	fmt.Fprintf(w, "<table class=\"%s\">\n", self.class)
	for _, row := range self.rows {
		fmt.Fprintf(w, "<tr class=\"%s\" title=\"%s\">\n", row.class, row.title) // TODO - maybe escape title ...
		for _, cel := range row.Cells {
			fmt.Fprintf(w, "<td class=\"%s\">\n", cel.class)
			if cel.Widget != nil {
				cel.Widget.Render(w)
			} else {
				fmt.Fprintf(w, cel.Content)
			}
			fmt.Fprintf(w, "</td>\n")
		}
		fmt.Fprintf(w, "</tr>\n")
	}
	fmt.Fprintf(w, "</table>\n")
	return nil
}

func (self *TableWidget) AddRow(row *TableRowWidget) *TableWidget {
	self.rows = append(self.rows, row)
	return self
}

func (self *TableWidget) SetClass(s string) *TableWidget {
	self.class = s
	return self
}
