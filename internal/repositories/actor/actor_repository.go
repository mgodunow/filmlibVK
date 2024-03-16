package actorRepository

import (
	"context"
	"database/sql"
	"filmLib/internal/domain/actor"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

type ActorReposostory interface {
	AllActors(ctx context.Context) ([]actor.Actor, error)
	AddActor(ctx context.Context, actor actor.Actor) (uuid.UUID, error)
	UpdateActor(ctx context.Context, actorId uuid.UUID, actorDTO ActorDTO) error
	DeleteActor(ctx context.Context, actorid uuid.UUID) error
}

type ActorDTO struct {
	Name     string
	Gender   string
	Birthday time.Time
}

type actorPostgres struct {
	db     *sql.DB
	logger *log.Logger
	mutex  *sync.Mutex
}

func (a *actorPostgres) AllActors(ctx context.Context) ([]actor.Actor, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
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
	a.mutex.Lock()
	defer a.mutex.Unlock()

	_, err := a.db.Exec("INSERT INTO actors VALUES($1,$2,$3,$4)", actor.Id, actor.Name,
		actor.Gender, actor.Birthday)

	return actor.Id, err
}

func (a *actorPostgres) UpdateActor(ctx context.Context, actorId uuid.UUID, actorDTO ActorDTO) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	_, err := a.db.Exec("UPDATE actors SET name=$1, gender=$2, birthday=$3",
	actorDTO.Name, actorDTO.Gender, actorDTO.Birthday)
	return err
}

func (a *actorPostgres) DeleteActor(ctx context.Context, actorId uuid.UUID) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	_, err := a.db.Exec("DELETE FROM actors WHERE id=$1", actorId)

	return err
}

func NewActorRepository(db *sql.DB, logger *log.Logger) ActorReposostory {
	return &actorPostgres{
		db: db,
		logger: logger,
	}
}