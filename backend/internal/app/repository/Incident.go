package repository

import (
	"database/sql"
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/storage/postgres"
)

type IncidentRepository struct {
	db                 *sql.DB
	databaseReposotiry *DatabaseRepository
}

func NewIncidentRepository(s *postgres.Storage, databaseRepository *DatabaseRepository) *IncidentRepository {
	return &IncidentRepository{db: s.GetDb(), databaseReposotiry: databaseRepository}
}

func (a *IncidentRepository) Get(id int) (*entity.Incident, error) {
	row := a.db.QueryRow("select id, db_id, error, date from incidents WHERE id = $1", id)

	var d entity.Incident
	var dbId int
	switch err := row.Scan(&d.Id, &dbId, &d.Error, &d.Date); err {
	case sql.ErrNoRows:
		return nil, err
	}
	d.Db, _ = a.databaseReposotiry.Get(dbId)
	return &d, nil
}
