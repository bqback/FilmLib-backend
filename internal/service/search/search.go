package search

import (
	"context"
	"filmlib/internal/pkg/entities"
	"filmlib/internal/storage"
)

type SearchService struct {
	as storage.IActorStorage
	ms storage.IMovieStorage
}

func NewSearchService(actorStorage storage.IActorStorage, movieStorage storage.IMovieStorage) *SearchService {
	return &SearchService{
		as: actorStorage,
		ms: movieStorage,
	}
}

func (s *SearchService) FindMovies(ctx context.Context, search_term string) ([]*entities.Movie, error) {
	movies, err := s.ms.FindByString(ctx, search_term)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
