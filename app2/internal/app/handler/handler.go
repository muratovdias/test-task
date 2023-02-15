package handler

import (
	"app2/internal/app/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Handler struct {
	service *service.Service
	router  *chi.Mux
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
		router:  chi.NewRouter(),
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	h.router.Use(middleware.DefaultLogger)
	h.router.Post("/create-user", h.createUser)
	h.router.Get("/get-user/{email}", h.getUser)
	return h.router
}
