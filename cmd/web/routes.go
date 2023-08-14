package main

import (
	"log"

	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/mux"

	"github.com/rs/cors"
)

func (a *application) routes() http.Handler {
	
	mux := mux.NewRouter()

	mux.Use(a.LoadSession)
	if a.debug{
		mux.Use(loggingMiddleware)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		a.session.Put(r.Context(), "test", "TuanNV")
		err := a.render(w, r, "home", nil)
		if err != nil {
			log.Fatal(err)
		}
	}).Methods("GET")

	mux.HandleFunc("/cooking", func(w http.ResponseWriter, r *http.Request) {

		vars := make(jet.VarMap)
		vars.Set("testSession", a.session.GetString(r.Context(), "test"))
		err := a.render(w, r, "cooking", vars)
		if err != nil {
			log.Fatal(err)
		}
	}).Methods("GET")

	fileServer := http.FileServer(http.Dir("../../public"))
	mux.Handle("/public/*", http.StripPrefix("/public", fileServer))
	
	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:80"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Use CORS middleware for all request
	handler := c.Handler(mux)

	return handler
}