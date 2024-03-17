package movieRepository

import (
	"context"
	"database/sql"
	"filmLib/internal/domain/movie"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
	logger *logrus.Logger
}

// TODO: check that all actors exists
func (m *moviePostgres) AddMovie(ctx context.Context, movie movie.Movie) (uuid.UUID, error) {
	_, err := m.db.Exec("INSERT INTO films VALUES($1, $2, $3, $4, $5, $6)",
		movie.Id, movie.Title, movie.Description, movie.Date, movie.Rate, movie.ActorsIn)

	return movie.Id, err
}

func (m *moviePostgres) AllMovies(ctx context.Context) ([]movie.Movie, error) {
	var movies []movie.Movie
	var movie movie.Movie

	rows, err := m.db.Query("SELECT uuid, title, description, date, rate FROM films")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.Date,
			&movie.Rate)

		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)

	}

	return movies, nil
}

// TODO: check that all actors exists
func (m *moviePostgres) PatchMovie(ctx context.Context, movieId uuid.UUID, title, description string,
	date time.Time, rate int, actorsIn []uuid.UUID) error {

	_, err := m.db.Exec("UPDATE films SET title=$1, description=$2, date=$3, rate=$4, actorsIn=$5 WHERE id = $6",
		title, description, date, rate, actorsIn, movieId)
	return err
}

func (m *moviePostgres) DeleteMovie(ctx context.Context, movieId uuid.UUID) error {
	_, err := m.db.Exec("DELETE FROM films WHERE id=$1", movieId)
	return err
}

func NewMovieRepository(db *sql.DB, logger *logrus.Logger) MovieRepository {
	return &moviePostgres{
		db:     db,
		logger: logger,
	}
}
