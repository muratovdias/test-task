package app

import (
	"log"
	"net/http"
	"user/internal/app/config"
	"user/internal/app/handler"
	"user/internal/app/repository"
	"user/internal/app/service"

	"github.com/go-chi/chi"
)

type App struct {
	repository *repository.Repository
	service    *service.Service
	handler    *handler.Handler
	config     config.Config
}

func NewApp() *App {
	var app App
	app.config = config.LoadMongoConfig()
	mongo := repository.NewMongoDB(app.config)
	app.repository = repository.NewRepository(mongo)
	app.service = service.NewService(app.repository)
	app.handler = handler.NewHandler(app.service, app.config)
	return &app
}

func ServerUp(router *chi.Mux) *http.Server {
	return &http.Server{
		Addr:    ":8888",
		Handler: router,
	}
}

func (app *App) Run() error {
	server := ServerUp(app.handler.InitRoutes())
	if err := server.ListenAndServe(); err != nil {
		log.Println(err.Error() + " server: Start")
		return err
	}
	return nil
}
