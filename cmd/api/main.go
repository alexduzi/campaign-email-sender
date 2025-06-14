package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type product struct {
	ID   int
	Name string
}

type myHandler struct{}

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("myHandler"))
}

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("before")
		next.ServeHTTP(w, r)
		println("after")
	})
}

func main() {
	r := chi.NewRouter()

	// m := myHandler{}
	// r.Handle("/handler", m)

	r.Use(myMiddleware)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

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

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {

		obj := map[string]string{"message": "success!"}

		render.JSON(w, r, obj)
	})

	r.Post("/product", func(w http.ResponseWriter, r *http.Request) {
		var product product

		render.DecodeJSON(r.Body, &product)

		product.ID = 5

		render.JSON(w, r, product)
	})

	http.ListenAndServe(":3000", r)
}
