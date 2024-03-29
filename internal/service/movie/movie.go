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

func (s *MovieService) Read(ctx context.Context, id dto.MovieID) (*entities.Movie, error) {
	movie, err := s.ms.Read(ctx, id)
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

func (s *MovieService) Update(ctx context.Context, info dto.UpdatedMovie) (*entities.Movie, error) {
	movie, err := s.ms.Update(ctx, info)
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

func (s *MovieService) Delete(ctx context.Context, id dto.MovieID) error {
	return s.ms.Delete(ctx, id)
}

func (s *MovieService) GetMovies(ctx context.Context, opts dto.SortOptions) ([]*entities.Movie, error) {
	movies, err := s.ms.GetMovies(ctx, opts)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
