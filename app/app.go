package app

import (
	"database/sql"
	"filmLib/config"
	"filmLib/filmlibAPI"
	actorRepository "filmLib/internal/repositories/actor"
	movieRepository "filmLib/internal/repositories/movie"
	"filmLib/internal/service"
	"fmt"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

type App struct {
	cfg config.Config
	logger *logrus.Logger
	server *filmlibAPI.Server
}

func NewApp() *App {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Can't load config: %s", err)
	}
	logger := logrus.New()

	db, err := dbConnect(cfg)
	if err != nil {
		log.Fatalf("Can't connect to db: %s", err)
		
	}

	movieRepo := movieRepository.NewMovieRepository(db, logger)
	actorRepo := actorRepository.NewActorRepository(db, logger)

	service := service.NewService(movieRepo, actorRepo, logger)

	server, err := filmlibAPI.NewServer(service, nil)
	
	if err != nil {
		log.Fatalf("Can't create server: %s", err)
	}

	return &App{
		cfg: cfg,
		logger: logger,
		server: server,
	}
}

func dbConnect(cfg config.Config) (*sql.DB, error) {
	connectStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	return sql.Open("postgres", connectStr)
}

func (a *App) ListenAndServe() {
	log.Printf("Server runs on port: %d", a.cfg.AppPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", a.cfg.DBPort), a.server); err != nil {
		log.Fatal(err)
	}
}