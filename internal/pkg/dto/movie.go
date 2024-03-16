package dto

type NewMovie struct {
	Name        string
	Gender      string
	ReleaseDate string `json:"release_date" db:"release"`
	Actors      []string
}

type UpdatedMovie struct {
	ID          uint64
	Name        string
	Gender      string
	ReleaseDate string `json:"release_date" db:"release"`
	Actors      []string
}

type MovieInfo struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type MovieID struct {
	Value uint64
}
