package dto

import "time"

type NewMovie struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date" db:"release"`
	Rating      float32   `json:"rating"`
	Actors      []uint64  `json:"movie_actors" db:"-"`
}

type UpdatedMovie struct {
	ID          uint64    `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date" db:"release"`
	Rating      float32   `json:"rating"`
	Actors      []uint64  `json:"movie_actors" db:"-"`
}

type MovieInfo struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type MovieID struct {
	Value uint64
}
