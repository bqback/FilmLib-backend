package service

import (
	"filmlib/internal/auth"
	"filmlib/internal/service/actor"
	authService "filmlib/internal/service/auth"
	"filmlib/internal/service/movie"
	"filmlib/internal/service/search"
	"filmlib/internal/storage"
)

type Services struct {
	Actor  IActorService
	Movie  IMovieService
	Search ISearchService
	Auth   IAuthService
}

func NewServices(storages *storage.Storages, manager *auth.AuthManager) *Services {
	return &Services{
		Actor:  actor.NewActorService(storages.Actor),
		Movie:  movie.NewMovieService(storages.Movie),
		Search: search.NewSearchService(storages.Actor, storages.Movie),
		Auth:   authService.NewAuthService(storages.Auth, manager),
	}
}
