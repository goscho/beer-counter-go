package data

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func Connect() (*pgxpool.Pool, error) {

	connStr := "postgresql://postgres:example@localhost:5432/test"
	pool, err := pgxpool.Connect(context.Background(), connStr)

	if err != nil {
		log.Println("Connecting DB failed ", err)
		return nil, err
	}

	if pool.Ping(context.Background()) != nil {
		log.Println("DB ping failed ", err)
		return nil, err
	}

	log.Print("DB is reachable")
	return pool, nil
}
