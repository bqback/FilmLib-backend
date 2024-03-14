package postgresql

import "database/sql"

type PgMovieStorage struct {
	db *sql.DB
}

func NewMovieStorage(db *sql.DB) *PgMovieStorage {
	return &PgMovieStorage{
		db: db,
	}
}
