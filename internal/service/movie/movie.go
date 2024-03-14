package movie

import "filmlib/internal/storage"

type MovieService struct {
	ms storage.IMovieStorage
}

func NewMovieService(movieStorage storage.IMovieStorage) *MovieService {
	return &MovieService{
		ms: movieStorage,
	}
}
