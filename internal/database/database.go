package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mr-celos/Atlas/internal/config"
)

func Connect(cfg config.Config) (*pgxpool.Pool, error) { // connects to the database and returns a connection pool
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	pool, poolErr := pgxpool.New(context.Background(), connString)

	if poolErr != nil {
		return pool, poolErr
	}

	pingErr := pool.Ping(context.Background())
	if pingErr != nil {
		return pool, pingErr
	}

	return pool, nil
}
