package main

import (
	"os"

	"github.com/AndreyDodonov/machclub_back"
	"github.com/AndreyDodonov/machclub_back/pkg/handler"
	"github.com/AndreyDodonov/machclub_back/pkg/repository"
	"github.com/AndreyDodonov/machclub_back/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
//	logrus.SetFormatter(new(logrus.JSONFormatter))

	gin.SetMode(gin.ReleaseMode) // release mode. For debug mode comment this line

	if err := initConfig(); err != nil {
		logrus.Fatalf("error in initialization configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db. Error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(apiserver.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while starting http server: %s", err.Error())
	}
}

// init configuration ...
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
