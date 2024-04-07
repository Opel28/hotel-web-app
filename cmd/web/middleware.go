package main

import (
	"context"
	"log"
	"net/http"
)

//func WriteToConsole(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("Hello World")
//		next.ServeHTTP(w, r)
//	})
//}

// SessionLoad is loading session to the middleware
func SessionLoad(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := app.Session.Get(r, "session-name")
		if err != nil {
			// Deletes old or invalid cookies
			http.SetCookie(w, &http.Cookie{Name: "session-name", MaxAge: -1, Path: "/"})
			log.Fatal("Error getting session: " + err.Error())
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "session", session))
		h.ServeHTTP(w, r)
	})
}
