package main

import (
	"net/http"

	"github.com/elkcityhazard/gocial-links/internal/handlers"
)

func routes() http.Handler {

	srv := http.NewServeMux()

	srv.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// do something
			handlers.HandleGetHome(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	srv.HandleFunc("/link", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		default:
			http.NotFound(w, r)
		}
	})

	return redirectTrailingSlash(srv)

}
