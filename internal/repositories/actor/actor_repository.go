package actorRepository

import (
	"context"
	"database/sql"
	"filmLib/internal/domain/actor"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ActorReposostory interface {
	AllActors(ctx context.Context) ([]actor.Actor, error)
	AddActor(ctx context.Context, actor actor.Actor) (uuid.UUID, error)
	UpdateActor(ctx context.Context, actorId uuid.UUID, name, gender string, birthday time.Time) error
	DeleteActor(ctx context.Context, actorid uuid.UUID) error
}

type actorPostgres struct {
	db     *sql.DB
	logger *logrus.Logger
}

func (a *actorPostgres) AllActors(ctx context.Context) ([]actor.Actor, error) {
	rows, err := a.db.Query("SELECT * FROM actors")
	if err != nil {
		return nil, err
	}
	var actors []actor.Actor
	var actor actor.Actor

	for rows.Next() {
		err := rows.Scan(&actor.Id, &actor.Name, &actor.Gender, &actor.Birthday)
		if err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}

	return actors, nil

}

func (a *actorPostgres) AddActor(ctx context.Context, actor actor.Actor) (uuid.UUID, error) {
	_, err := a.db.Exec("INSERT INTO actors VALUES($1,$2,$3,$4)", actor.Id, actor.Name,
		actor.Gender, actor.Birthday)

	return actor.Id, err
}

func (a *actorPostgres) UpdateActor(ctx context.Context, actorId uuid.UUID,
	name, gender string, birthday time.Time) error {

	_, err := a.db.Exec("UPDATE actors SET name=$1, gender=$2, birthday=$3",
		name, gender, birthday)
	return err
}

func (a *actorPostgres) DeleteActor(ctx context.Context, actorId uuid.UUID) error {
	_, err := a.db.Exec("DELETE FROM actors WHERE id=$1", actorId)

	return err
}

func NewActorRepository(db *sql.DB, logger *logrus.Logger) ActorReposostory {
	return &actorPostgres{
		db:     db,
		logger: logger,
	}
}
