package main

import (
	"github.com/gin-gonic/gin"
	_"github.com/mattn/go-sqlite3"
	"database/sql"
	"log"
)

func main() {

	db, err := sql.Open("sqlite3", "./temphu.db")

	if err != nil {
		log.Fatal("db error" + err.Error())
	}

	r := gin.Default()
	r.Static("/web", "./web")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/temp", func(c *gin.Context) {
		temps, err := db.Query("SELECT `temp` FROM `temphu` ORDER BY DATE(`date`) DESC LIMIT 1")

		if err != nil {
			println(err.Error())
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		temps.Next()

		var temp int

		err = temps.Scan(&temp)

		c.JSON(200, gin.H{
			"temp": temp,
		})
	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}