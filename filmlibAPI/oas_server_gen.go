// Code generated by ogen, DO NOT EDIT.

package filmlibAPI

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AddActor implements addActor operation.
	//
	// Добавление актёра.
	//
	// POST /actors
	AddActor(ctx context.Context, req *Actor) error
	// AddMovie implements addMovie operation.
	//
	// Добавление нового фильма.
	//
	// POST /movies
	AddMovie(ctx context.Context, req *Movie) error
	// DeleteActorByID implements deleteActorByID operation.
	//
	// Удаление актёра по ID.
	//
	// DELETE /actors/{actor_id}
	DeleteActorByID(ctx context.Context, params DeleteActorByIDParams) error
	// GetAllActors implements getAllActors operation.
	//
	// Получение списка актёров.
	//
	// GET /actors
	GetAllActors(ctx context.Context) (Actors, error)
	// GetMovies implements getMovies operation.
	//
	// Получение отсортированного списка фильмов.
	//
	// GET /movies
	GetMovies(ctx context.Context, params GetMoviesParams) (Movies, error)
	// MoviesMovieIDDelete implements DELETE /movies/{movie_id} operation.
	//
	// Удаление фильма по ID.
	//
	// DELETE /movies/{movie_id}
	MoviesMovieIDDelete(ctx context.Context, params MoviesMovieIDDeleteParams) error
	// MoviesMovieIDPatch implements PATCH /movies/{movie_id} operation.
	//
	// Изменение информации о фильме.
	//
	// PATCH /movies/{movie_id}
	MoviesMovieIDPatch(ctx context.Context, params MoviesMovieIDPatchParams) error
	// PatchActorById implements patchActorById operation.
	//
	// Изменение информации об актёре.
	//
	// PATCH /actors/{actor_id}
	PatchActorById(ctx context.Context, params PatchActorByIdParams) error
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
