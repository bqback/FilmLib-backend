package dto

import (
	"encoding/json"
	"errors"
	"time"
)

type NewActor struct {
	Name      string    `json:"name"`
	Gender    string    `json:"gender" validate:"oneof=male female other"`
	BirthDate time.Time `json:"dob" validate:"datetime"`
}

type ExpectedActorUpdate struct {
	Name      *string    `json:"name,omitempty" validate:"required_without_all=Gender BirthDate"`
	Gender    *string    `json:"gender,omitempty" validate:"oneof=male female other,required_without_all=Name BirthDate"`
	BirthDate *time.Time `json:"dob,omitempty"  validate:"datetime,required_without_all=Name Gender"`
}

type UpdatedActor struct {
	ID     uint64
	Values map[string]interface{}
	// Name      *string    `json:"name,omitempty" validate:"required_without_all=Gender BirthDate"`
	// Gender    *string    `json:"gender,omitempty" validate:"oneof=male female other,required_without_all=Name BirthDate"`
	// BirthDate *time.Time `json:"dob,omitempty"  validate:"datetime,required_without_all=Name Gender"`
}

type ActorInfo struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type ActorInfoList []ActorInfo

func (m *ActorInfoList) Scan(value interface{}) error {
	switch value := value.(type) {
	case []byte:
		return json.Unmarshal(value, &m)
	case nil:
		return nil
	default:
		return errors.New("failed asserting value type")
	}
}

type GetAllActor struct {
	ID        uint64
	Name      string
	Gender    string
	BirthDate time.Time     `db:"dob"`
	Movies    MovieInfoList `json:"movies" db:"movies"`
}

type ActorID struct {
	Value uint64
}
