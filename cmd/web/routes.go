package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/gorilla/securecookie"
	"net/http"
	"web-app/pkg/config"
	"web-app/pkg/handlers"
)

func routes(app config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	key := securecookie.GenerateRandomKey(32)
	mux.Use(csrf.Protect(key, csrf.Secure(app.InProduction), csrf.HttpOnly(true), csrf.SameSite(csrf.SameSiteLaxMode)))
	//mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)

	return mux
}
