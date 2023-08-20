package main

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/rs/cors"
)

func (a *application) routes() http.Handler {

	mux := mux.NewRouter()

	mux.Use(a.LoadSession)
	if a.debug {
		mux.Use(loggingMiddleware)
	}
	//User Route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		err := a.render(w, r, "home", nil)
		if err != nil {
			log.Fatal(err)
		}
	}).Methods("GET")

	//Admin
	mux.HandleFunc("/admin", a.getAllIngredient).Methods("GET")
	mux.HandleFunc("/admin/recipe/store", a.storeRecipe).Methods("POST")

	//Static files
	fs := http.StripPrefix("/public/",  http.FileServer(http.Dir("../../public")))
	mux.PathPrefix("/public/").Handler(fs).Methods("GET")

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
