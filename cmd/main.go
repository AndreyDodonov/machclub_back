package main

import (
	"log"

	"github.com/AndreyDodonov/machclub_back"
	"github.com/AndreyDodonov/machclub_back/pkg/handler"
	"github.com/AndreyDodonov/machclub_back/pkg/repository"
	"github.com/AndreyDodonov/machclub_back/pkg/service"

	"github.com/spf13/viper"
	_ "github.com/lib/pq"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error in initialization configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db. Error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(apiserver.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while starting http server: %s", err.Error())
	}
}

// init configuration ...
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}