package main

import (
	"beer-counter-go/internal/data"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	dbConnection, err := data.Connect()
	defer dbConnection.Close()

	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	engine.GET("/db-user", func(c *gin.Context) {
		var user string
		query := "select usename from pg_user"
		err := dbConnection.QueryRow(context.Background(), query).Scan(&user)

		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"db-user": user,
		})
	})

	err = engine.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal("Problem attaching to webserver ", err)
	}
}
