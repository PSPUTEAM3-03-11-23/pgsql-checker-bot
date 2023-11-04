package LocationValidator

import (
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/pkg/errorHandler"
	"net/http"
)

func ValidateLocation(location *entity.Location) *errorHandler.HttpErr {
	if location.Latitude == nil || *location.Latitude < -90 || *location.Latitude > 90 {
		return errorHandler.New("invalid latitude", http.StatusBadRequest)
	}

	if location.Longitude == nil || *location.Longitude < -180 || *location.Longitude > 180 {
		return errorHandler.New("invalid longitude", http.StatusBadRequest)
	}
	return nil
}
