package handlers

import "web-app/pkg/config"

// Repo is the variable for handlers can use Repository
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App config.AppConfig
}

// NewRepo sets config to Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: *a,
	}
}

// NewHandlers sets configured Repository to variable Repo
func NewHandlers(r *Repository) {
	Repo = r
}
