package entities

import (
	"filmlib/internal/pkg/dto"
	"time"
)

type Actor struct {
	ID        uint64          `json:"id"`
	Name      string          `json:"name"`
	Gender    string          `json:"gender"`
	BirthDate time.Time       `json:"dob" db:"dob"`
	Movies    []dto.MovieInfo `json:"actor_movies" db:"-"`
}

type Movie struct {
	ID          uint64          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ReleaseDate time.Time       `json:"release_date" db:"release_date"`
	Rating      float32         `json:"rating"`
	Actors      []dto.ActorInfo `json:"movie_actors" db:"-"`
}

type SearchResult struct {
	Actors []Actor `json:"actors"`
	Movies []Movie `json:"movies"`
}

type Role struct {
	RoleName string
	IsAdmin  bool
}
