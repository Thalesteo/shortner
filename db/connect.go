package db

import (
	"github.com/Thalesteo/trypgx/queries"
)

type Queries struct {
	*queries.UserQueries
}

func OpenConnection() (*Queries, error) {
	db, err := PostgreSqlConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		UserQueries: &queries.UserQueries{DB: db},
	}, nil
}
