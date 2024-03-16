package movie

import (
	"filmLib/internal/domain/actor"
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	Id          uuid.UUID
	Title       string
	Description string
	Date        time.Time
	Rate        int
	ActorsIn    []actor.Actor
}

func NewMovie(id uuid.UUID, title, description string,
	date time.Time, rate int,
	actorsIn []actor.Actor) Movie {
	return Movie{
		Id:          id,
		Title:       title,
		Description: description,
		Date:        date,
		Rate:        rate,
		ActorsIn:    actorsIn,
	}
}
