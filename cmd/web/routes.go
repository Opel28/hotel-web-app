package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"net/http"
	"os"
	"web-app/pkg/config"
	"web-app/pkg/handlers"
)

func routes(app config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Use(csrf.Protect([]byte(os.Getenv("HOTEL_CSRF_KEY")), csrf.Secure(app.InProduction), csrf.HttpOnly(true), csrf.SameSite(csrf.SameSiteLaxMode)))
	//mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)

	fileServer := http.FileServer(http.Dir("./html-source/"))
	mux.Handle("/html-source/*", http.StripPrefix("/html-source", fileServer))

	return mux
}
