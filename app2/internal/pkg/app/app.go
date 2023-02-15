package app

import (
	"app2/internal/app/config"
	"app2/internal/app/handler"
	"app2/internal/app/repository"
	"app2/internal/app/server"
	"app2/internal/app/service"
	"log"
)

type App struct {
	repository *repository.Repository
	service    *service.Service
	handler    *handler.Handler
	config     config.MongoConfig
}

func NewApp() *App {
	var app App
	app.config = config.LoadMongoConfig()
	mongo := repository.NewMongoDB(app.config)
	app.repository = repository.NewRepository(mongo)
	app.service = service.NewService(app.repository)
	app.handler = handler.NewHandler(app.service)
	return &app
}

func (app *App) Run() error {
	s := server.ServerUp(app.handler.InitRoutes())
	if err := server.Start(s); err != nil {
		log.Println(err.Error() + " server: Start")
		return err
	}
	return nil
}
