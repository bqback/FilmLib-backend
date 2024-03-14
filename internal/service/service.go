package service

import (
	"filmlib/internal/service/actor"
	"filmlib/internal/service/movie"
	"filmlib/internal/service/search"
	"filmlib/internal/storage"
)

type Services struct {
	Actor  IActorService
	Movie  IMovieService
	Search ISearchService
}

func NewServices(storages *storage.Storages) *Services {
	return &Services{
		Actor:  actor.NewActorService(storages.Actor),
		Movie:  movie.NewMovieService(storages.Movie),
		Search: search.NewSearchService(storages.Actor, storages.Movie),
	}
}
