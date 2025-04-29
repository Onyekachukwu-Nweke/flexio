package app

import (
	"database/sql"
	"flexio-api/config"
	"flexio-api/internal/api"
	"flexio-api/internal/store"
	"flexio-api/migrations"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication(cfg *config.Config) (*Application, error) {
	pgDB, err := store.Open(cfg)
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	// stores
	workoutStore := store.NewPostgresWorkoutStore(pgDB)

	// handlers
	workoutHandler := api.NewWorkoutHandler(workoutStore)

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}
