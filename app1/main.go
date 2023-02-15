package main

import (
	"app1/internal/handler"
	"log"
	"net/http"
)

func main() {
	log.Println("App1 started at port 8887")
	if err := http.ListenAndServe(":8887", handler.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}
