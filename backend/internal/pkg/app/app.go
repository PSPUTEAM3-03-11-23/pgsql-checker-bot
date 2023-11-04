package app

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"hackathon-tg-bot/internal/app/router"
	"hackathon-tg-bot/pkg/config"
	"net/http"
)

type App struct {
	router *gin.Engine
}

// New Конструктор приложения
func New(r *gin.Engine) *App {
	return &App{router: router.InitRoutes(r)}
}

func (a *App) GetServer() *http.Server {
	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%s", config.GetConfig().GetString("server.address"),
			config.GetConfig().GetString("server.port")),
		Handler: a.router,
	}
	return srv
}

func (a *App) Run(srv *http.Server) error {
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err

	}
	return nil
}
