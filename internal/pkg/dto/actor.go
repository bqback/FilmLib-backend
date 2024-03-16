package dto

import "time"

type NewActor struct {
	Name      string    `json:"name"`
	Gender    string    `json:"gender" validate:"oneof=male female other"`
	BirthDate time.Time `json:"dob" validate:"datetime"`
}

type UpdatedActor struct {
	ID        uint64    `json:"-"`
	Name      string    `json:"name,omitempty" validate:"required_without_all=Gender BirthDate"`
	Gender    string    `json:"gender,omitempty" validate:"oneof=male female other,required_without_all=Name BirthDate"`
	BirthDate time.Time `json:"dob,omitempty"  validate:"datetime,required_without_all=Name Gender"`
}

type ActorInfo struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type ActorID struct {
	Value uint64
}
