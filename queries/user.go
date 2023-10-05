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

func (q *UserQueries) DeleteUser(id uuid.UUID) bool {
	query := `DELETE FROM Users WHERE id=$1 LIMITE=1`
	res, err := q.Exec(query, id)
	if err == nil {
		count, err := res.RowsAffected()
		if err == nil {
			// delete only current user
			return count == 1
		}
	}

	return false
}

func (q *UserQueries) CheckUserEmail(email string, u *models.User) bool {

	query := `SELECT id, name, email, password FROM users WHERE email=$1 LIMIT=1`
	rows, err := q.Queryx(query, email)
	if err != nil {
		return false
	}

	for rows.Next() {
		err = rows.StructScan(&u)
		if err != nil {
			return false
		}
	}

	return true
}
