package repository

import (
	"database/sql"
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/model/input"
	"hackathon-tg-bot/internal/app/storage/postgres"
)

type DatabaseRepository struct {
	db *sql.DB
}

func NewDatabaseRepository(s *postgres.Storage) *DatabaseRepository {
	return &DatabaseRepository{db: s.GetDb()}
}

func (a *DatabaseRepository) GetAll() (*[]entity.Database, error) {
	rows, err := a.db.Query("select id, host, port, username, password, db_name, schema, title from dbs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var databases []entity.Database

	for rows.Next() {
		d := entity.Database{}
		err := rows.Scan(&d.Id, &d.Host, &d.Port, &d.Username, &d.Password, &d.DbName, &d.Schema, &d.Title)
		if err != nil {
			continue
		}
		databases = append(databases, d)
	}
	return &databases, nil
}

func (a *DatabaseRepository) Get(id int) (*entity.Database, error) {
	row := a.db.QueryRow("select id, host, port, username, password, db_name, schema, title from dbs WHERE id = $1", id)

	var d entity.Database

	switch err := row.Scan(&d.Id, &d.Host, &d.Port, &d.Username, &d.Password, &d.DbName, &d.Schema, &d.Title); err {
	case sql.ErrNoRows:
		return nil, err
	}
	return &d, nil
}

func (a *DatabaseRepository) Create(databaseInput *input.Database) (*entity.Database, error) {
	var id int
	_ = a.db.QueryRow("insert into dbs(host, port, username, password, db_name, schema, title) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		databaseInput.Host, databaseInput.Port, databaseInput.Username, databaseInput.Password, databaseInput.DbName, databaseInput.Schema, databaseInput.Title).
		Scan(&id)

	e, err := a.Get(id)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (a *DatabaseRepository) Delete(id int) (bool, error) {
	res, err := a.db.Exec("DELETE FROM dbs WHERE id=$1", id)
	if err == nil {
		count, err := res.RowsAffected()
		if err == nil {
			if count > 0 {
				return true, nil
			}
		}

	}

	return false, nil
}
