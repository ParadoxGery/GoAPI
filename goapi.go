package main

import (
	"github.com/gin-gonic/gin"
	_"github.com/mattn/go-sqlite3"
	"database/sql"
	"log"
	"github.com/stianeikeland/go-rpio"
	"github.com/d2r2/go-dht"
)

func main() {

	db, err := sql.Open("sqlite3", "./temphu.db")

	if err != nil {
		log.Fatal("db error " + err.Error())
	}

	err = rpio.Open()

	if err != nil {
		log.Fatal("gpio error " + err.Error())
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
	r.PUT("/pew", func(c *gin.Context) {
		pin := rpio.Pin(10)
		pin.Output()
		pin.High()
	})
	r.DELETE("/pew", func(c *gin.Context) {
		pin := rpio.Pin(10)
		pin.Output()
		pin.Low()
	})
	r.GET("/temps", func(c *gin.Context) {
		temp, hum, ret, err := dht.ReadDHTxxWithRetry(dht.DHT11, 6, true, 5)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"temp": temp,
			"hum" : hum,
			"tries" : ret,
		})
	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}