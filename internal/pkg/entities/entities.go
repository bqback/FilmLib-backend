package entities

import (
	"filmlib/internal/pkg/dto"
	"time"
)

type Actor struct {
	ID        uint64          `json:"id"`
	Name      string          `json:"name"`
	Gender    string          `json:"gender"`
	BirthDate time.Time       `json:"dob"`
	Movies    []dto.MovieInfo `json:"actor_movies"`
}

type Movie struct {
	ID          uint64          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ReleaseDate time.Time       `json:"release_date"`
	Rating      float32         `json:"rating"`
	Actors      []dto.ActorInfo `json:"movie_actors"`
}

type SearchResult struct {
	Actors []*Actor
	Movies []*Movie
}
