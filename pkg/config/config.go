package config

import (
	"github.com/gorilla/sessions"
	"html/template"
	"log"
)

// AppConfig holds application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *sessions.CookieStore
}
