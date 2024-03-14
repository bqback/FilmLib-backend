package entities

import "time"

type Actor struct {
	ID        uint64
	Name      string
	Gender    string
	BirthDate time.Time
	Movies    []Movie
}

type Movie struct {
	ID          uint64
	Title       string
	Description string
	ReleaseDate time.Time
	Rating      float32
	Actors      []Actor
}
