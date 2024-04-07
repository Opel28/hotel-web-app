package main

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"web-app/pkg/config"
	"web-app/pkg/handlers"
	"web-app/pkg/render"
)

const (
	portNum = ":8080"
)

var (
	app     config.AppConfig
	key     = securecookie.GenerateRandomKey(32)
	session *sessions.CookieStore
)

func main() {
	// set true when in production
	app.InProduction = false

	// sessions setup and saving to configs
	session = sessions.NewCookieStore(key)
	session.Options = &sessions.Options{MaxAge: 20 * 60 * 60, HttpOnly: true, Secure: app.InProduction, SameSite: http.SameSiteLaxMode}
	app.Session = session

	// using template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache ", err)
	}
	app.TemplateCache = tc
	app.UseCache = true

	//using repository
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.SetConfig(&app)

	//Starting server
	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(app),
	}

	fmt.Println("Listening on port " + portNum)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
