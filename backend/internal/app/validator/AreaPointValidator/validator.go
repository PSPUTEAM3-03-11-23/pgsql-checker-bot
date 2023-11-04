package AreaPointValidator

import (
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/pkg/errorHandler"
	"net/http"
)

func ValidateAreaPoint(areaPoint *entity.AreaPoint) *errorHandler.HttpErr {
	if areaPoint.Latitude == nil || *areaPoint.Latitude < -90 || *areaPoint.Latitude > 90 {
		return errorHandler.New("invalid latitude", http.StatusBadRequest)
	}

	if areaPoint.Longitude == nil || *areaPoint.Longitude < -180 || *areaPoint.Longitude > 180 {
		return errorHandler.New("invalid longitude", http.StatusBadRequest)
	}
	return nil
}
