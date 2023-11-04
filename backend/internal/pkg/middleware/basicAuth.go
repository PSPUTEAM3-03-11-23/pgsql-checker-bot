package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/model/input"
	"hackathon-tg-bot/internal/app/repository"
	"hackathon-tg-bot/internal/app/storage/postgres"
)

func DecodeCredentials(c *gin.Context) (string, string, bool) {
	r := c.Request
	return r.BasicAuth()
}

func GetAccountByCreds(c *gin.Context) (*entity.User, error) {
	login, password, ok := DecodeCredentials(c)
	if !ok {
		return nil, errors.New("")
	}

	loginInput := &input.Login{
		Email:    &login,
		Password: &password,
	}
	storage, _ := postgres.Get()
	userRepo := repository.NewUserRepository(storage)
	user, _ := userRepo.GetByCreds(loginInput)
	return user, nil
}

// BasicAuth middleware для basic auth
func BasicAuth(c *gin.Context) {
	//user, err := GetAccountByCreds(c)
	//
	//if err != nil || user == nil {
	//	c.AbortWithStatus(http.StatusUnauthorized)
	//	c.Next()
	//	return
	//}
	//
	//c.Set("user", user)
	c.Next()
}
