package main

import (
	"app1/internal/pkg/app"
	"log"
)

func main() {
	app := app.NewApp()
	log.Println("App1 started at port 8887")
	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
