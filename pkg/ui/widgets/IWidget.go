package widgets

import "net/http"

type IWidget interface {
	Render(w http.ResponseWriter) error
}
