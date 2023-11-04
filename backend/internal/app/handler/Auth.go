package handler

import (
	"github.com/gin-gonic/gin"
	"hackathon-tg-bot/internal/app/model/input"
	"hackathon-tg-bot/internal/app/model/response"
	"hackathon-tg-bot/pkg/errorHandler"
	"net/http"
)

type AuthService interface {
	Login(login *input.Login) (*response.User, *errorHandler.HttpErr)
}

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (d *AuthHandler) Login(c *gin.Context) {
	loginInput := &input.Login{}
	err := c.BindJSON(&loginInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	user, httpErr := d.authService.Login(loginInput)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
