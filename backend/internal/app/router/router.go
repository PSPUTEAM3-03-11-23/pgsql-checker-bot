package router

import (
	"github.com/gin-gonic/gin"
	"hackathon-tg-bot/internal/app/handler"
	"hackathon-tg-bot/internal/app/repository"
	"hackathon-tg-bot/internal/app/service"
	"hackathon-tg-bot/internal/app/storage/postgres"
	"hackathon-tg-bot/internal/pkg/middleware"
)

// InitRoutes Инициализация путей эндпоинтов, сервисов и репозиториев
func InitRoutes(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")
	storage, err := postgres.New()
	if err != nil {
		panic("cant connect to DB")
	}
	databaseRepo := repository.NewDatabaseRepository(storage)
	databaseService := service.NewDatabaseService(databaseRepo)
	databaseHandler := handler.NewDatabaseHandler(databaseService)
	databasesGroup := api.Group("database")
	{
		databasesGroup.GET("/", middleware.BasicAuth, databaseHandler.GetAll)
		databasesGroup.GET("/:id", middleware.BasicAuth, databaseHandler.Get)
		databasesGroup.POST("/", middleware.BasicAuth, databaseHandler.Create)
	}

	return r
}
