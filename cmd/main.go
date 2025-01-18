package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/MDmitryM/banking-app-go/pkg/handler"
	"github.com/MDmitryM/banking-app-go/pkg/repository"
	"github.com/MDmitryM/banking-app-go/pkg/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title						banking app API
// @version					    1.0
// @description				    API server for banking application
// @host						localhost:8080
// @BasePath					/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	fmt.Println("Hello from banking app!")
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("Can't init configs, %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("cant load .env file, err - %s", err.Error())
	}

	mongodbUri := "mongodb://" + viper.GetString("dev_db.username") + ":" +
		os.Getenv("MONGO_PASSWORD") + "@" +
		viper.GetString("dev_db.host") + ":" +
		viper.GetString("dev_db.port")

	mongoConfig := repository.MongoConfig{
		URI:      mongodbUri,
		Database: viper.GetString("dev_db.database"),
		Timeout:  viper.GetDuration("dev_db.timeout"),
	}

	db, err := repository.NewMongoDB(mongoConfig)
	if err != nil {
		logrus.Fatalf("error while creating DB, err - %s", err.Error())
	}
	defer db.Close(context.Background())

	repository := repository.NewRepository(db)
	service := service.NewService(repository)

	hander := handler.NewHandler(service)
	echo := echo.New()
	hander.SetupRouts(echo)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	//Starting server
	go func() {
		if err := echo.Start(":" + viper.GetString("port")); err != nil && err != http.ErrServerClosed {
			logrus.Fatal("shutting down the server, error - ", err.Error())
		}
	}()
	//Wait for interrupt signal to gracefully shut down the server with timeout of 10 seconds
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := echo.Shutdown(ctx); err != nil {
		logrus.Fatal(err)
	}
	logrus.Println("server is gracefully shut downed")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
