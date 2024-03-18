package dto

import (
	"encoding/json"
	"errors"
	"time"
)

type NewMovie struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date" db:"release_date"`
	Rating      float32   `json:"rating"`
	Actors      []uint64  `json:"movie_actors" db:"-"`
}

type ExpectedMovieUpdate struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date" db:"release_date"`
	Rating      float32   `json:"rating"`
	Actors      []uint64  `json:"movie_actors" db:"-"`
}

type UpdatedMovie struct {
	ID     uint64
	Values map[string]interface{}
}

type MovieInfo struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type MovieInfoList []MovieInfo

func (m *MovieInfoList) Scan(value interface{}) error {
	switch value := value.(type) {
	case []byte:
		return json.Unmarshal(value, &m)
	case nil:
		return nil
	default:
		return errors.New("failed asserting value type")
	}
}

type GetAllMovie struct {
	ID          uint64
	Title       string
	Description string
	ReleaseDate time.Time `db:"release_date"`
	Rating      float32
	Actors      ActorInfoList `json:"actors" db:"actors"`
}

type MovieID struct {
	Value uint64
}
