package main

import (
	"log"
	"github.com/AndreyDodonov/machclub_back"
	"github.com/AndreyDodonov/machclub_back/internal/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(machclubback.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while starting http server: %s", err.Error())
	}
}
