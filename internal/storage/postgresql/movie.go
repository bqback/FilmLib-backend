package postgresql

import "github.com/jmoiron/sqlx"

type PgMovieStorage struct {
	db *sqlx.DB
}

func NewMovieStorage(db *sqlx.DB) *PgMovieStorage {
	return &PgMovieStorage{
		db: db,
	}
}
