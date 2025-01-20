package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), "postgresql://postgres:passw0rd@localhost:5432/tixmaster")
	if err != nil {
		panic(err)
	}

	return dbpool
}
