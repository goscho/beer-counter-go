package data

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

func Connect() (*sql.DB, error) {

	connStr := "postgresql://postgres:example@localhost:5432/test"
	conn, err := sql.Open("pgx", connStr)

	if err != nil {
		log.Println("Connecting DB failed ", err)
		return nil, err
	}

	if conn.Ping() != nil {
		log.Println("DB ping failed ", err)
		return nil, err
	}

	log.Print("DB is reachable")
	return conn, nil
}
