package main

import (
	// "context"
	"database/sql"
	// oasFilmLib "filmLib/filmlibAPI"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "filmlib"
	password = "1111"
	dbname   = "filmlib"
)

type M struct {
	Id uuid.UUID
	Title string
	Description string
	Date time.Time
	Rate int
	ActorsList []uuid.UUID
}

func main() {
	// service := &filmlibService{}

	// server, err := oasFilmLib.NewServer(service, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Fatal(http.ListenAndServe(":8080", server))
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Println(err)
	}

// 	_, err = db.Exec("INSERT INTO films VALUES($1, $2, $3, $4, $5, $6)", "550e8400-e29b-41d4-a716-446655440000", "Django freedom",
// "desc", "2007-10-11", 10, "{550e8400-e29b-41d4-a716-446655440000, 550e8400-e29b-41d4-a716-446655440000}")
// 	if err != nil {
// 		log.Println(err)
// 	}

	var m M
	res := db.QueryRow("SELECT * FROM films")
	if err != nil {
		log.Println(err)
	}
	res.Scan(&m.Id, &m.Title, &m.Description, &m.Date, &m.Rate, &m.ActorsList)
	fmt.Println(m)
}

// type filmlibService struct {
// 	movies []oasFilmLib.Movie
// 	actors []oasFilmLib.Actor
// }

// func (f *filmlibService) AddActor(ctx context.Context, req *oasFilmLib.Actor) error {
// 	f.actors = append(f.actors, *req)
// 	log.Println("Successfully add actor")
// 	log.Println(f.actors)

// 	return nil
// }

// func (f *filmlibService) AddMovie(ctx context.Context, req *oasFilmLib.Movie) error {
// 	log.Println("Successfully add movie")
// 	return nil
// }

// func (f *filmlibService) DeleteActorByID(ctx context.Context,
// 	params oasFilmLib.DeleteActorByIDParams) error {
// 	return nil
// }

// func (f *filmlibService) GetAllActors(ctx context.Context) (oasFilmLib.Actors, error) {
// 	return f.actors, nil
// }

// func (f *filmlibService) GetMovies(ctx context.Context, params oasFilmLib.GetMoviesParams) (oasFilmLib.Movies, error) {
// 	return f.movies, nil
// }

// func (f *filmlibService) MoviesMovieIDDelete(ctx context.Context,
// 	params oasFilmLib.MoviesMovieIDDeleteParams) error {
// 	return nil
// }

// func (f *filmlibService) MoviesMovieIDPatch(ctx context.Context,
// 	params oasFilmLib.MoviesMovieIDPatchParams) error {
// 	return nil
// }

// func (f *filmlibService) PatchActorById(ctx context.Context, params oasFilmLib.PatchActorByIdParams) error {
// 	return nil
// }

// func (f *filmlibService) NewError(ctx context.Context, err error) *oasFilmLib.ErrorStatusCode {
// 	return &oasFilmLib.ErrorStatusCode{StatusCode: 503}
// }
