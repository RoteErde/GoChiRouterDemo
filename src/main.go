package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

/**
Building up tiny examples of Chi routing

*/
func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	//REST API for obtaining user
	r.Route("/person", func(r chi.Router) {
		r.Route("/{UserID}", func(r chi.Router) {
			r.Get("/", GetPerson)
		})

	})

	http.ListenAndServe(":3000", r)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(chi.URLParam(r, "UserID")))
}
