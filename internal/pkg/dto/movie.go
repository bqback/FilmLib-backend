package dto

type NewMovie struct {
	Name      string
	Gender    string
	BirthDate string
	Actors    []string
}

type UpdatedMovie struct {
	ID        uint64
	Name      string
	Gender    string
	BirthDate string
	Actors    []string
}

type MovieInfo struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type MovieID struct {
	Value uint64
}
