package handler

import (
	"github.com/gin-gonic/gin"
	"hackathon-tg-bot/internal/app/model/response"
	"hackathon-tg-bot/internal/app/validator"
	"hackathon-tg-bot/pkg/errorHandler"
	"net/http"
)

type AlertService interface {
	GetAll() (*[]response.Alert, *errorHandler.HttpErr)
	Get(id int) (*response.Alert, *errorHandler.HttpErr)
}

type AlertHandler struct {
	alertService AlertService
}

func NewAlertHandler(alertService AlertService) *AlertHandler {
	return &AlertHandler{alertService: alertService}
}

func (d *AlertHandler) GetAll(c *gin.Context) {
	alert, httpErr := d.alertService.GetAll()
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

	alert, httpErr := d.alertService.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	c.JSON(http.StatusOK, alert)
}
