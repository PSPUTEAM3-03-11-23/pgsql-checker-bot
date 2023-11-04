package mapper

import (
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/model/response"
)

func DatabaseToDatabaseResponse(database *entity.Database) *response.Database {
	r := &response.Database{
		Id:       database.Id,
		Host:     database.Host,
		Port:     database.Port,
		Username: database.Username,
		Password: database.Password,
		DbName:   database.DbName,
		Schema:   database.Schema,
		Title:    database.Title,
	}

	return r
}

func DatabaseToDatabaseResponses(databases *[]entity.Database) *[]response.Database {
	rs := make([]response.Database, 0)

	for _, database := range *databases {
		rs = append(rs, *DatabaseToDatabaseResponse(&database))
	}

	return &rs
}
