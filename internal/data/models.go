package data

import (
	"database/sql"
	"log"
	"time"
)

type Token struct {
	ID string `json:"id"`
}

type Person struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Token *Token `json:"token,omitempty"`
}

type Item struct {
	id   string
	name string
}

type Transaction struct {
	id        string
	person    Person
	item      Item
	timeStamp time.Time
}

type PersonModel struct {
	DB *sql.DB
}

func (pm PersonModel) FindAll() ([]Person, error) {
	rows, err := pm.DB.Query("SELECT ID, Name, Token FROM beer_counter.person")
	if err != nil {
		log.Println(err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	var people = make([]Person, 0)

	for rows.Next() {
		var (
			id      string
			name    string
			tokenId sql.NullString
		)

		err := rows.Scan(&id, &name, &tokenId)
		if err != nil {
			log.Println("error reading result ", err)
			return nil, err
		}

		person := Person{ID: id, Name: name}

		if tokenId.Valid {
			person.Token = &Token{ID: tokenId.String}
		}

		people = append(people, person)
	}
	return people, nil
}
