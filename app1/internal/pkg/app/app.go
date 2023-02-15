package app

import (
	"app1/internal/app/handler"
	"app1/internal/app/server"
	"app1/internal/app/service"
	"net/http"
)

type App struct {
	handler *handler.Handler
	service *service.Service
	server  *http.Server
	routes  http.Handler
}

func NewApp() *App {
	var app App
	app.service = service.NewService()
	app.handler = handler.NewHandler(app.service)
	app.routes = app.handler.InitRoutes()
	return &app
}

func (app *App) Run() error {
	app.server = server.ServerUp(app.routes)
	if err := server.Start(app.server); err != nil {
		return err
	}
	return nil
}
