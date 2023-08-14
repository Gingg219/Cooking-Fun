package main

import (
	"log"
	"net/http"
)

func (a *application) LoadSession(next http.Handler) http.Handler {
	return a.session.LoadAndSave(next)
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
        log.Printf("Request URL: %s", r.RequestURI)
        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}