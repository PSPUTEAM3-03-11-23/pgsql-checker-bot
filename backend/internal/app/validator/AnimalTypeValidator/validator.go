package AnimalTypeValidator

import (
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/validator"
	"hackathon-tg-bot/pkg/errorHandler"
	"net/http"
)

func ValidateAnimalType(animalType *entity.AnimalType) *errorHandler.HttpErr {
	if validator.IsStringEmpty(animalType.Type) {
		return errorHandler.New("type is empty", http.StatusBadRequest)
	}
	return nil
}
