package handler

import (
	"fmt"
	style2 "github.com/vit1251/golden/internal/site/style"
	"net/http"
)

type StyleHandler struct {
}

func NewStyleHandler() *StyleHandler {
	return &StyleHandler{}
}

func (h *StyleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	css1 := style2.NewCSSStyleSheet()
	rule1 := style2.NewCSSRule()

	// Message preview box
	//    rule1.SetSelectorText(".message-preview")
	//    rule1.Set("border", "1px solid yellow")

	css1.InsertRule(rule1)

	content := css1.String()

	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))
	w.Header().Set("Content-Type", " text/css; charset=utf-8")
	w.WriteHeader(200)

	w.Write([]byte(content))
}
