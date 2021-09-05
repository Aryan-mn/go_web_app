package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func writeToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler)http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}


// sessionLoad loads and saves the session on every request
func sessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}