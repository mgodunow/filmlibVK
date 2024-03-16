package actor

import (
	"time"

	"github.com/google/uuid"
)

type Actor struct {
	Id uuid.UUID
	Name     string
	Gender   string
	Birthday time.Time
}

func NewActor(id uuid.UUID, name, gender string, birthday time.Time) Actor{
	return Actor{
		Id: id,
		Name: name,
		Gender: gender,
		Birthday: birthday,
	}
}