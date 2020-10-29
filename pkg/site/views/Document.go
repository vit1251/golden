package views

import (
	"github.com/vit1251/golden/pkg/site/widgets"
	"html/template"
	"net/http"
)

type Document struct {
	Layout    string
	Page      string
	params    map[string]interface{}
}

func NewDocument() *Document {
	doc := new(Document)
	doc.params = make(map[string]interface{})
	doc.InitMenu()
	return doc
}

func (self *Document) SetLayout(path string) error {
	self.Layout = path
	return nil
}

func (self *Document) SetPage(path string) error {
	self.Page = path
	return nil
}

func (self *Document) SetParam(name string, value interface{}) error {
	self.params[name] =  value
	return nil
}

//func (self *Document) SetWidget(w IWidget) error {
//}

func (self *Document) Render(w http.ResponseWriter) error {
	var tmpl *template.Template
	tmpl, err1 := template.ParseFiles(self.Layout, self.Page)
	if err1 != nil {
		return err1
	}
	err2 := tmpl.ExecuteTemplate(w, "layout", self.params)
	if err2 != nil {
		return err2
	}
	return nil
}

func (self *Document) InitMenu() {
	mmw := widgets.NewMainMenuWidget()
	menus := mmw.Init()
	self.SetParam("Menus", menus)
}
