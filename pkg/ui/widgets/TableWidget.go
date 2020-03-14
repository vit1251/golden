package widgets

import "net/http"

type TableHeaderWidget struct {
}

type TableRowWidget struct {
}

type TableWidget struct {
	headers    TableHeaderWidget
	rows    []*TableRowWidget
}

func NewTableWidget() *TableWidget {
	w := new(TableWidget)
	return w
}

func (self *TableWidget) CreateHeaderRow() {

}

func (self *TableWidget) CreateRow() *TableRowWidget {
	row := new(TableRowWidget)
	self.rows = append(self.rows, row)
	return row
}

func (self *TableWidget) Render(w http.ResponseWriter) {
}
