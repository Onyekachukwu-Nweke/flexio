package store

import (
	"database/sql"
	"flexio-api/config"
	"fmt"
	"github.com/pressly/goose/v3"
	"io/fs"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Open(cfg *config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBName,
		cfg.DBPassword)

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	fmt.Println("Connected to Database...")
	return db, nil
}

func MigrateFS(db *sql.DB, migrationFS fs.FS, dir string) error {
	goose.SetBaseFS(migrationFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()
	return Migrate(db, dir)
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("db migrate error: goose %w", err)
	}

	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("goose up error: %w", err)
	}
	return nil
}
