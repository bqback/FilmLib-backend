package postgresql

import "database/sql"

type PgActorStorage struct {
	db *sql.DB
}

func NewActorStorage(db *sql.DB) *PgActorStorage {
	return &PgActorStorage{
		db: db,
	}
}
