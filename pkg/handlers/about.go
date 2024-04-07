package handlers

import (
	"net/http"
	"web-app/pkg/models"
	"web-app/pkg/render"
)

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["hello"] = "Hello World"
	render.RenderTemplate(w, "about", &models.TemplateData{
		StringMap: stringMap,
	})
}
