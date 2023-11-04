package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"hackathon-tg-bot/internal/pkg/app"
	"hackathon-tg-bot/pkg/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Init Инициализация сервиса
func Init() {
	config.Init("development")
}

func main() {
	Init()
	r := gin.Default()

	a := app.New(r)
	srv := a.GetServer()

	go func() {
		err := a.Run(srv)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Stopping http server")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
	}
}
