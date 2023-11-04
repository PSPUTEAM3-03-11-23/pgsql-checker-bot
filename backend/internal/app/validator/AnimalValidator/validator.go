package AnimalValidator

import (
	"fmt"
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/model/input"
	"hackathon-tg-bot/internal/app/model/response"
	"hackathon-tg-bot/pkg/errorHandler"
	"net/http"
)

func ValidateLifeStatus(lifeStatus string) *errorHandler.HttpErr {
	if lifeStatus != entity.Alive && lifeStatus != entity.Dead {
		return errorHandler.New(fmt.Sprintf("lifeStatus must be in [%s, %s]", entity.Alive, entity.Dead), http.StatusBadRequest)
	}
	return nil
}

func ValidateGender(gender string) *errorHandler.HttpErr {
	if gender != entity.Male && gender != entity.Female && gender != entity.Other {
		return errorHandler.New(fmt.Sprintf("gender must be in [%s, %s, %s]", entity.Male, entity.Female, entity.Other), http.StatusBadRequest)
	}
	return nil
}

func validateAnimalInput(input *input.Animal) *errorHandler.HttpErr {
	if input.Weight == nil {
		return errorHandler.New("weight is missing", http.StatusBadRequest)
	}
	if *input.Weight <= 0 {
		return errorHandler.New("weight must be greater than 0", http.StatusBadRequest)
	}

	if input.Length == nil {
		return errorHandler.New("length is missing", http.StatusBadRequest)
	}
	if *input.Length <= 0 {
		return errorHandler.New("length must be greater than 0", http.StatusBadRequest)
	}

	if input.Height == nil {
		return errorHandler.New("height is missing", http.StatusBadRequest)
	}
	if *input.Height <= 0 {
		return errorHandler.New("height must be greater than 0", http.StatusBadRequest)
	}

	if input.Gender == nil {
		return errorHandler.New("gender is missing", http.StatusBadRequest)
	}
	err := ValidateGender(*input.Gender)
	if err != nil {
		return err
	}

	if input.ChipperId == nil {
		return errorHandler.New("chipper_id is missing", http.StatusBadRequest)
	}
	if *input.ChipperId <= 0 {
		return errorHandler.New("chipper_id must be greater than 0", http.StatusBadRequest)
	}

	if input.ChippingLocationId == nil {
		return errorHandler.New("chipping_location_id is missing", http.StatusBadRequest)
	}
	if *input.ChippingLocationId <= 0 {
		return errorHandler.New("chipping_location_id must be greater than 0", http.StatusBadRequest)
	}
	return nil
}

func ValidateAnimalCreateInput(input *input.Animal) *errorHandler.HttpErr {
	if input.AnimalTypeIds == nil || len(input.AnimalTypeIds) == 0 {
		return errorHandler.New("animal types are empty", http.StatusBadRequest)
	}

	animalTypeIds := map[int]bool{}
	for _, animalTypeId := range input.AnimalTypeIds {
		if animalTypeId <= 0 {
			return errorHandler.New("animal type id must be greater than 0", http.StatusBadRequest)
		}
		if animalTypeIds[animalTypeId] {
			return errorHandler.New("duplicated animal type id", http.StatusConflict)
		}
		animalTypeIds[animalTypeId] = true
	}

	httpErr := validateAnimalInput(input)
	if httpErr != nil {
		return httpErr
	}

	return nil
}

func ValidateAnimalUpdateInput(input *input.Animal, oldAnimal *response.Animal) *errorHandler.HttpErr {
	httpErr := validateAnimalInput(input)
	if httpErr != nil {
		return httpErr
	}

	if *input.LifeStatus == entity.Alive && oldAnimal.LifeStatus == entity.Dead {
		return errorHandler.New("cant set status Alive to Dead animal", http.StatusBadRequest)
	}
	return nil
}

func ValidateAnimalTypeUpdateInput(input *input.AnimalTypeUpdate) *errorHandler.HttpErr {
	if input.OldTypeId == nil {
		return errorHandler.New("oldTypeId is missing", http.StatusBadRequest)
	}
	if *input.OldTypeId <= 0 {
		return errorHandler.New("oldTypeId must be greater than 0", http.StatusBadRequest)
	}
	if input.NewTypeId == nil {
		return errorHandler.New("newTypeId is missing", http.StatusBadRequest)
	}
	if *input.NewTypeId <= 0 {
		return errorHandler.New("newTypeId must be greater than 0", http.StatusBadRequest)
	}
	return nil
}

func ValidateAnimalLocationPointUpdate(input *input.AnimalLocationPointUpdate) *errorHandler.HttpErr {
	if input.VisitedLocationPointId == nil {
		return errorHandler.New("visitedLocationPointId is missing", http.StatusBadRequest)
	}
	if *input.VisitedLocationPointId <= 0 {
		return errorHandler.New("visitedLocationPointId must be greater than 0", http.StatusBadRequest)
	}
	if input.LocationPointId == nil {
		return errorHandler.New("locationPointId is missing", http.StatusBadRequest)
	}
	if *input.LocationPointId <= 0 {
		return errorHandler.New("locationPointId must be greater than 0", http.StatusBadRequest)
	}
	return nil
}
