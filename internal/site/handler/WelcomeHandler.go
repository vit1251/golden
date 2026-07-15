package handler

import (
    "net/http"
    "github.com/vit1251/golden/pkg/registry"
    "github.com/vit1251/golden/internal/site/views"
)

type WelcomeHandler struct {
    registry *registry.Container
}

func NewWelcomeHandler(registry *registry.Container) *WelcomeHandler {
    return &WelcomeHandler{
	registry: registry,
    }
}

func (h *WelcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    err := views.Page("Golden Point", views.WelcomeView()).Render(w)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
