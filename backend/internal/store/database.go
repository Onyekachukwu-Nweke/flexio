package store

import (
	"database/sql"
	"flexio-api/config"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Open(cfg *config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBName,
		cfg.DBPassword,
		cfg.DBSSLMode)

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	fmt.Println("Connected to Database...")
	return db, nil
}
