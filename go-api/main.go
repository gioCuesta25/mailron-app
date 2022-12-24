package main

import (
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/Giovanni-Romana-Cuesta/go-api/utils"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/documents/match", MatchPhrase)

	http.ListenAndServe(":4000", r)
}

func MatchPhrase(w http.ResponseWriter, r *http.Request) {

	term := r.URL.Query().Get("term")
	from := r.URL.Query().Get("from")

	resp := utils.MakeSearchRequest(term, from)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	data := utils.ParseResponse(body)

	w.Write(data)

}
