package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

func ServerUp(router *chi.Mux) *http.Server {
	return &http.Server{
		Addr:    ":8888",
		Handler: router,
	}
}

func Start(server *http.Server) error {
	return server.ListenAndServe()
}
