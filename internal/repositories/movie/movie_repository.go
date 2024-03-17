package movieRepository

import (
	"context"
	"database/sql"
	"filmLib/internal/domain/movie"
	"filmLib/pkg/repeatable"
	"time"

	"github.com/google/uuid"
)

type MovieRepository interface {
	AddMovie(ctx context.Context, movie movie.Movie) (uuid.UUID, error)
	AllMovies(ctx context.Context) ([]movie.Movie, error)
	PatchMovie(ctx context.Context, movieId uuid.UUID, title, description string,
		date time.Time, rate int, actorsIn []uuid.UUID) error
	DeleteMovie(ctx context.Context, movieId uuid.UUID) error
}

type moviePostgres struct {
	db     *sql.DB
}

// TODO: check that all actors exists
func (m *moviePostgres) AddMovie(ctx context.Context, movie movie.Movie) (uuid.UUID, error) {
	const op = "MovieRepository.AddMovie"

	_, err := m.db.Exec("INSERT INTO films VALUES($1, $2, $3, $4, $5, $6)",
		movie.Id, movie.Title, movie.Description, movie.Date, movie.Rate, movie.ActorsIn)

	return movie.Id, repeatable.WrapError(op, err)
}

func (m *moviePostgres) AllMovies(ctx context.Context) ([]movie.Movie, error) {
	const op = "MovieRepository.AllMovies"

	var movies []movie.Movie
	var movie movie.Movie

	rows, err := m.db.Query("SELECT uuid, title, description, date, rate FROM films")

	if err != nil {
		return nil, repeatable.WrapError(op, err)
	}

	for rows.Next() {
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.Date,
			&movie.Rate)

		if err != nil {
			return nil, repeatable.WrapError(op, err)
		}

		movies = append(movies, movie)

	}

	return movies, nil
}

// TODO: check that all actors exists
func (m *moviePostgres) PatchMovie(ctx context.Context, movieId uuid.UUID, title, description string,
	date time.Time, rate int, actorsIn []uuid.UUID) error {

	const op = "MovieRepository.PatchMovie"

	_, err := m.db.Exec("UPDATE films SET title=$1, description=$2, date=$3, rate=$4, actorsIn=$5 WHERE id = $6",
		title, description, date, rate, actorsIn, movieId)

	if err != nil {
		return repeatable.WrapError(op, err)
	}

	return err
}

func (m *moviePostgres) DeleteMovie(ctx context.Context, movieId uuid.UUID) error {
	const op = "MovieRepository.DeleteMovie"

	_, err := m.db.Exec("DELETE FROM films WHERE id=$1", movieId)

	if err != nil {
		return repeatable.WrapError(op, err)
	}

	return err
}

func NewMovieRepository(db *sql.DB) MovieRepository {
	return &moviePostgres{
		db:     db,
	}
}
