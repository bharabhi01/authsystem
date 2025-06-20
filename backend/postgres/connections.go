package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func InitDB() error {
	var err error
	pool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	return nil
}

func GetPool() *pgxpool.Pool {
	return pool
}

func CloseDB() {
	if pool != nil {
		pool.Close()
	}
}
