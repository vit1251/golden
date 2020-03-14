package views

import (
	"github.com/vit1251/golden/pkg/ui/widgets"
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
	var menus []*widgets.MenuAction

	if menuAction := widgets.NewMenuAction(); menuAction != nil {
		menuAction.Link = "/"
		menuAction.Label = "Home"
		menus = append(menus, menuAction)
	}

	if menuAction := widgets.NewMenuAction(); menuAction != nil {
		menuAction.Link = "/netmail"
		menuAction.Label = "Netmail"
		menus = append(menus, menuAction)
	}

	if menuAction := widgets.NewMenuAction(); menuAction != nil {
		menuAction.Link = "/echo"
		menuAction.Label = "Echomail"
		menus = append(menus, menuAction)
	}

	if menuAction := widgets.NewMenuAction(); menuAction != nil {
		menuAction.Link = "/file"
		menuAction.Label = "Filebox"
		menus = append(menus, menuAction)
	}

	if menuAction := widgets.NewMenuAction(); menuAction != nil {
		menuAction.Link = "/stat"
		menuAction.Label = "Statistics"
		menus = append(menus, menuAction)
	}

	if menuAction := widgets.NewMenuAction(); menuAction != nil {
		menuAction.Link = "/service"
		menuAction.Label = "Service"
		menus = append(menus, menuAction)
	}

	if menuAction := widgets.NewMenuAction(); menuAction != nil {
		menuAction.Link = "/setup"
		menuAction.Label = "Setup"
		menus = append(menus, menuAction)
	}

	self.SetParam("Menus", menus)
}
