package main

import (
	"log"

	"github.com/AndreyDodonov/machclub_back"
	"github.com/AndreyDodonov/machclub_back/pkg/handler"
	"github.com/AndreyDodonov/machclub_back/pkg/repository"
	"github.com/AndreyDodonov/machclub_back/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(apiserver.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while starting http server: %s", err.Error())
	}
}
