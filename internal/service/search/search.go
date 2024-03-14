package search

import "filmlib/internal/storage"

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
