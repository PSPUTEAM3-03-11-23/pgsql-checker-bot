package service

import (
	"fmt"
	"hackathon-tg-bot/internal/app/mapper"
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/model/response"
	"hackathon-tg-bot/pkg/errorHandler"
	"net/http"
)

type AlertRepository interface {
	GetAll() (*[]entity.Alert, error)
	Get(int) (*entity.Alert, error)
}
type AlertService struct {
	alertRepo AlertRepository
}

func NewAlertService(alertRepo AlertRepository) *AlertService {
	return &AlertService{alertRepo: alertRepo}
}

func (d *AlertService) GetAll() (*[]response.Alert, *errorHandler.HttpErr) {
	alertResponses := &[]response.Alert{}

	alert, err := d.alertRepo.GetAll()
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	alertResponses = mapper.AlertToAlertResponses(alert)

	return alertResponses, nil
}

func (d *AlertService) Get(id int) (*response.Alert, *errorHandler.HttpErr) {
	alertResponse := &response.Alert{}
	alert, err := d.alertRepo.Get(id)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	if alert == nil {
		return nil, errorHandler.New(fmt.Sprintf("Alert with id %d does not exists", id), http.StatusNotFound)
	}

	alertResponse = mapper.AlertToAlertResponse(alert)

	return alertResponse, nil
}
