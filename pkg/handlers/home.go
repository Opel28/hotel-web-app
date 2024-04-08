package handlers

import (
	"fmt"
	"net/http"
	"web-app/pkg/models"
	"web-app/pkg/render"
)

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {

	//session := r.Context().Value("session").(*sessions.Session)
	session, err := m.App.Session.Get(r, "session-name")
	if err != nil {
		http.Error(w, "Unable to get session!", http.StatusInternalServerError)
		fmt.Println(err)
	}

	session.Values["remote-ip"] = r.RemoteAddr

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Unable to save session!", http.StatusInternalServerError)
	}

	fmt.Println(r.RemoteAddr)

	render.RenderTemplate(w, "home", &models.TemplateData{})
}
