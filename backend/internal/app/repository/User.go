package repository

import (
	"database/sql"
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/model/input"
	"hackathon-tg-bot/internal/app/storage/postgres"
)

type UserRepository struct {
	db                 *sql.DB
	incidentRepository *IncidentRepository
}

func NewUserRepository(s *postgres.Storage) *UserRepository {
	return &UserRepository{db: s.GetDb()}
}

func (a *UserRepository) GetByCreds(loginInput *input.Login) (*entity.User, error) {
	row := a.db.QueryRow("select id, name, is_deactivated, email from users WHERE email = $1 AND password = $2", loginInput.Email, loginInput.Password)

	var u entity.User

	switch err := row.Scan(&u.Id, &u.Name, &u.IsDeactivated, &u.Email); err {
	case sql.ErrNoRows:
		return nil, err
	}
	return &u, nil
}
