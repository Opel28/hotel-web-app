package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"os"
	"web-app/pkg/config"
	"web-app/pkg/handlers"
	"web-app/pkg/render"
)

const (
	portNum = ":8080"
)

// TODO: Add encryptionKey
var (
	app   config.AppConfig
	store = sessions.NewCookieStore([]byte(os.Getenv("HOTEL_SESSION_KEY")))
)

func main() {
	// set true when in production
	app.InProduction = false

	//key := securecookie.GenerateRandomKey(32)
	//keyString := hex.EncodeToString(key)
	//fmt.Println(keyString)
	// sessions setup and saving to configs
	store.Options = &sessions.Options{MaxAge: 20 * 60 * 60, HttpOnly: true, Secure: app.InProduction, SameSite: http.SameSiteLaxMode}
	app.Session = store

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
