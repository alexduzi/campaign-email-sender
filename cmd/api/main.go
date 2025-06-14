package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		product := r.URL.Query().Get("product")
		id := r.URL.Query().Get("id")

		if product != "" && id != "" {
			w.Write([]byte("Hello world -> " + product + " id: " + id))
		} else {
			w.Write([]byte("Hello world"))
		}
	})

	r.Get("/{productName}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "productName")

		w.Write([]byte("Hello world -> " + param))
	})

	http.ListenAndServe(":3000", r)
}
