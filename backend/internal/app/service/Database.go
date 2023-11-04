package service

import (
	"fmt"
	"hackathon-tg-bot/internal/app/mapper"
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/model/input"
	"hackathon-tg-bot/internal/app/model/response"
	"hackathon-tg-bot/pkg/errorHandler"
	"net/http"
)

type DatabaseRepository interface {
	GetAll() (*[]entity.Database, error)
	Get(int) (*entity.Database, error)
	Create(databaseInput *input.Database) (*entity.Database, error)
}
type DatabaseService struct {
	databaseRepo DatabaseRepository
}

func NewDatabaseService(databaseRepo DatabaseRepository) *DatabaseService {
	return &DatabaseService{databaseRepo: databaseRepo}
}

func (d *DatabaseService) GetAll() (*[]response.Database, *errorHandler.HttpErr) {
	databaseResponses := &[]response.Database{}

	database, err := d.databaseRepo.GetAll()
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	databaseResponses = mapper.DatabaseToDatabaseResponses(database)

	return databaseResponses, nil
}

func (d *DatabaseService) Get(id int) (*response.Database, *errorHandler.HttpErr) {
	databaseResponse := &response.Database{}
	database, err := d.databaseRepo.Get(id)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	if database == nil {
		return nil, errorHandler.New(fmt.Sprintf("Database with id %d does not exists", id), http.StatusNotFound)
	}

	databaseResponse = mapper.DatabaseToDatabaseResponse(database)

	return databaseResponse, nil
}

func (d *DatabaseService) Create(databaseInput *input.Database) (*response.Database, *errorHandler.HttpErr) {
	databaseResponse := &response.Database{}

	database, err := d.databaseRepo.Create(databaseInput)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	databaseResponse = mapper.DatabaseToDatabaseResponse(database)

	return databaseResponse, nil
}
