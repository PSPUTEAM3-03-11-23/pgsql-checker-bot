package repository

import (
	"database/sql"
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/storage/postgres"
)

type AlertRepository struct {
	db                 *sql.DB
	incidentRepository *IncidentRepository
}

func NewAlertRepository(s *postgres.Storage, incidentRepository *IncidentRepository) *AlertRepository {
	return &AlertRepository{db: s.GetDb(),
		incidentRepository: incidentRepository,
	}
}

func (a *AlertRepository) GetAll(user *entity.User) (*[]entity.Alert, error) {
	rows, err := a.db.Query("select id, incident_id, is_sent from alerts WHERE user_id = $1", user.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var alerts []entity.Alert
	var incidentId int
	for rows.Next() {
		d := entity.Alert{}
		err := rows.Scan(&d.Id, &incidentId, &d.IsSent)
		if err != nil {
			continue
		}
		d.Incident, _ = a.incidentRepository.Get(incidentId)

		alerts = append(alerts, d)
	}
	return &alerts, nil
}

func (a *AlertRepository) Get(id int, user *entity.User) (*entity.Alert, error) {
	row := a.db.QueryRow("select id, incident_id, is_sent from alerts WHERE id = $1 AND user_id = $2", id, user.Id)

	var d entity.Alert
	var incidentId int

	switch err := row.Scan(&d.Id, &incidentId, &d.IsSent); err {
	case sql.ErrNoRows:
		return nil, err
	}
	d.Incident, _ = a.incidentRepository.Get(incidentId)

	return &d, nil
}
