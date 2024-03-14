package storage

import (
	"database/sql"
	"filmlib/internal/storage/postgresql"
)

type Storages struct {
	Actor IActorStorage
	Movie IMovieStorage
}

func NewPostgresStorages(db *sql.DB) *Storages {
	return &Storages{
		Actor: postgresql.NewActorStorage(db),
		Movie: postgresql.NewMovieStorage(db),
	}
}
