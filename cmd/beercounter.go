package main

import (
	"beer-counter-go/internal/data"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Env struct {
	persons data.PersonModel
}

func main() {

	dbConnection, err := data.Connect()

	defer func(dbConnection *sql.DB) {
		err := dbConnection.Close()
		if err != nil {
			log.Println(err)
		}
	}(dbConnection)

	env := &Env{
		persons: data.PersonModel{DB: dbConnection},
	}

	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	engine.GET("/person", env.personsAll)

	err = engine.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal("Problem attaching to webserver ", err)
	}
}

func (env *Env) personsAll(c *gin.Context) {
	people, err := env.persons.FindAll()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, people)
}
