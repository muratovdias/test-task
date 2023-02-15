package server

import "net/http"

func ServerUp(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":8887",
		Handler: handler,
	}
}

func Start(server *http.Server) error {
	return server.ListenAndServe()
}
