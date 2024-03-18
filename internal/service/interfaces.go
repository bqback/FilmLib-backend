package service

import (
	"context"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/pkg/entities"
)

type IActorService interface {
	GetActors(context.Context) ([]*entities.Actor, error)
	Create(context.Context, dto.NewActor) (*entities.Actor, error)
	Read(context.Context, dto.ActorID) (*entities.Actor, error)
	Update(context.Context, dto.UpdatedActor) (*entities.Actor, error)
	Delete(context.Context, dto.ActorID) error
}

type IMovieService interface {
	GetMovies(context.Context, dto.SortOptions) ([]*entities.Movie, error)
	Create(context.Context, dto.NewMovie) (*entities.Movie, error)
	Read(context.Context, dto.MovieID) (*entities.Movie, error)
	// Update(context.Context, dto.UpdatedMovie) (*entities.Movie, error)
	Delete(context.Context, dto.MovieID) error
}

type ISearchService interface {
	FindMovies(context.Context, string) ([]*entities.Movie, error)
}
