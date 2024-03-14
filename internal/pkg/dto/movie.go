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

type MovieID struct {
	Value uint64
}
