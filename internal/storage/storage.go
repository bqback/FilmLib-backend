package storage

import (
	"filmlib/internal/storage/postgresql"

	"github.com/jmoiron/sqlx"
)

type Storages struct {
	Actor IActorStorage
	Movie IMovieStorage
}

func NewPostgresStorages(db *sqlx.DB) *Storages {
	return &Storages{
		Actor: postgresql.NewActorStorage(db),
		Movie: postgresql.NewMovieStorage(db),
	}
}
