package main

import (
	"log"
	"user/internal/pkg/app"
)

func main() {
	app := app.NewApp()
	log.Println("App2 started at port 8888")
	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
