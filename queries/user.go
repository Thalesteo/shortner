package queries

import (
	"github.com/Thalesteo/trypgx/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

func (q *UserQueries) GetUser(id uuid.UUID) (models.User, error) {
	user := models.User{}

	query := `SELECT id, name, email, created_at FROM users WHERE id=$1`
	err := q.Get(&user, query, id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (q *UserQueries) CreateUser(u *models.User) error {
	query := `INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`

	_, err := q.Exec(query, u.ID, u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (q *UserQueries) CheckUserEmail(email string) models.User {
	user := models.User{}
	/*
		check if table users has entry with email
		if has -> return user
		else -> nil
	*/
	return user
}
