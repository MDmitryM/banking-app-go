package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/MDmitryM/banking-app-go/pkg/handler"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Hello from banking app!")
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("Can't init configs, %s", err.Error())
	}

	hander := handler.NewHandler()
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
