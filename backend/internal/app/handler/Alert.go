package handler

import (
	"github.com/gin-gonic/gin"
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/model/response"
	"hackathon-tg-bot/internal/app/validator"
	"hackathon-tg-bot/pkg/errorHandler"
	"net/http"
)

type AlertService interface {
	GetAll(user *entity.User) (*[]response.Alert, *errorHandler.HttpErr)
	Get(id int, user *entity.User) (*response.Alert, *errorHandler.HttpErr)
}

type AlertHandler struct {
	alertService AlertService
}

func NewAlertHandler(alertService AlertService) *AlertHandler {
	return &AlertHandler{alertService: alertService}
}

func (d *AlertHandler) GetAll(c *gin.Context) {
	userAny, _ := c.Get("user")
	user := userAny.(*entity.User)

	alert, httpErr := d.alertService.GetAll(user)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	c.JSON(http.StatusOK, alert)
}

func (d *AlertHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	userAny, _ := c.Get("user")
	user := userAny.(*entity.User)

	alert, httpErr := d.alertService.Get(id, user)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	c.JSON(http.StatusOK, alert)
}
