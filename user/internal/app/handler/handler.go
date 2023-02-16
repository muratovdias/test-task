package handler

import (
	"user/internal/app/config"
	"user/internal/app/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Handler struct {
	service *service.Service
	config  config.Config
}

func NewHandler(service *service.Service, conf config.Config) *Handler {
	return &Handler{
		service: service,
		config:  conf,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.DefaultLogger)
	router.Post("/create-user", h.createUser)
	router.Get("/get-user/{email}", h.getUser)
	return router
}
