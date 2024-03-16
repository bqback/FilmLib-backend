package storage

import (
	"context"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/pkg/entities"
)

type IActorStorage interface {
	// GetAll(context.Context) ([]*entities.Actor, error)
	Create(context.Context, dto.NewActor) (*entities.Actor, error)
	Read(context.Context, dto.ActorID) (*entities.Actor, error)
	// Update(context.Context, dto.UpdatedActor) (*entities.Actor, error)
	// Delete(context.Context, dto.ActorID) error
	GetActorMovies(context.Context, dto.ActorID) ([]dto.MovieInfo, error)
	// FindByName(context.Context, string) ([]*entites.Actor, error)
}

type IMovieStorage interface {
	// GetMovies(context.Context, dto.SortOptions) ([]*entities.Movie, error)
	// Create(context.Context, dto.NewMovie) (*entities.Movie, error)
	// Read(context.Context, dto.MovieID) (*entities.Movie, error)
	// Update(context.Context, dto.UpdatedMovie) (*entities.Movie, error)
	// Delete(context.Context, dto.MovieID) error
	// GetMovieActors(context.Context, dto.MovieID) ([]*dto.ActorInfo, error)
	// FindByTitle(context.Context, string) ([]*entites.Movie, error)
}
