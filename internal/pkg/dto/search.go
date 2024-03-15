package dto

import "filmlib/internal/pkg/entities"

type SearchResult struct {
	Actors []*entities.Actor
	Movies []*entities.Movie
}
