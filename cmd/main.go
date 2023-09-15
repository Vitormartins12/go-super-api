package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Started super api")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/home", HomeHandler)

	http.ListenAndServe(":8085", r)
}

// Endpoints
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
