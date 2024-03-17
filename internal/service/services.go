package service

import (
	"context"
	"filmLib/filmlibAPI"
	"filmLib/internal/domain/movie"
	"filmLib/internal/domain/actor"
	"log"
	"time"

	"github.com/google/uuid"
)

type MoviesRepository interface {
	AddMovie(ctx context.Context, movie movie.Movie) (uuid.UUID, error)
	AllMovies(ctx context.Context) ([]movie.Movie, error)
	PatchMovie(ctx context.Context, movieId uuid.UUID, title, description string,
		date time.Time, rate int, actorsIn []uuid.UUID) error
	DeleteMovie(ctx context.Context, movieId uuid.UUID) error
}

type ActorsRepository interface {
	AllActors(ctx context.Context) ([]actor.Actor, error)
	AddActor(ctx context.Context, actor actor.Actor) (uuid.UUID, error)
	UpdateActor(ctx context.Context, actorId uuid.UUID, name, gender string, birthday time.Time) error
	DeleteActor(ctx context.Context, actorid uuid.UUID) error
}

type filmlibService struct {
	moviesRepo MoviesRepository
	actorsRepo ActorsRepository
	logger     *log.Logger
}

func (f *filmlibService) AddActor(ctx context.Context, req *filmlibAPI.Actor) error {
	panic("Implement me!")
}

func (f *filmlibService) AddMovie(ctx context.Context, req *filmlibAPI.Movie) error {
	panic("Implement me!")
}

func (f *filmlibService) DeleteActorByID(ctx context.Context,
	params filmlibAPI.DeleteActorByIDParams) error {

	panic("Implement me!")
}

func (f *filmlibService) GetAllActors(ctx context.Context) (filmlibAPI.Actors, error) {
	panic("Implement me!")
}

func (f *filmlibService) GetMovies(ctx context.Context, params filmlibAPI.GetMoviesParams) (filmlibAPI.Movies, error) {
	panic("Implement me!")
}

func (f *filmlibService) MoviesMovieIDDelete(ctx context.Context,
	params filmlibAPI.MoviesMovieIDDeleteParams) error {
	panic("Implement me!")
}

func (f *filmlibService) MoviesMovieIDPatch(ctx context.Context,
	params filmlibAPI.MoviesMovieIDPatchParams) error {
	panic("Implement me!")
}

func (f *filmlibService) PatchActorById(ctx context.Context, params filmlibAPI.PatchActorByIdParams) error {
	panic("Implement me!")
}

func (f *filmlibService) NewError(ctx context.Context, err error) *filmlibAPI.ErrorStatusCode {
	panic("Implement me!")
}

func NewService(moviesRepo MoviesRepository, actorsRepo ActorsRepository, l *log.Logger) *filmlibService {
	return &filmlibService{
		moviesRepo: moviesRepo,
		actorsRepo: actorsRepo,
		logger: l,
	}
}