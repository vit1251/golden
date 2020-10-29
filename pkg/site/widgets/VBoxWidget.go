package widgets

import "net/http"

type VBoxWidget struct {
	Items []IWidget
}

func NewVBoxWidget() *VBoxWidget {
	vBox := new(VBoxWidget)
	return vBox
}

func (self *VBoxWidget) Render(w http.ResponseWriter) error {
	for _, item := range self.Items {
		item.Render(w)
	}
	return nil
}

func (self *VBoxWidget) Add(widget IWidget) *VBoxWidget {
	self.Items = append(self.Items, widget)
	return self
}
