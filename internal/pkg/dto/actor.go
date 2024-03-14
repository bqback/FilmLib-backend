package dto

type NewActor struct {
	Name      string
	Gender    string
	BirthDate string
}

type UpdatedActor struct {
	ID        uint64
	Name      string
	Gender    string
	BirthDate string
}

type ActorID struct {
	Value uint64
}
