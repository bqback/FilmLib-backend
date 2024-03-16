package movie

import (
	"context"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/pkg/entities"
	"filmlib/internal/storage"
)

type MovieService struct {
	ms storage.IMovieStorage
}

func NewMovieService(movieStorage storage.IMovieStorage) *MovieService {
	return &MovieService{
		ms: movieStorage,
	}
}

func (s *MovieService) Create(ctx context.Context, info dto.NewMovie) (*entities.Movie, error) {
	movie, err := s.ms.Create(ctx, info)
	if err != nil {
		return nil, err
	}
	actors, err := s.ms.GetMovieActors(ctx, dto.MovieID{Value: movie.ID})
	if err != nil {
		return nil, err
	}
	movie.Actors = actors
	return movie, nil
}
