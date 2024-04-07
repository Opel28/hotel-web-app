package handlers

import (
	"net/http"
	"web-app/pkg/models"
	"web-app/pkg/render"
)

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home", &models.TemplateData{})
}
