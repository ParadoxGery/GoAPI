package main

import (
	"github.com/gin-gonic/gin"
	"GoAPI/sensors"
	"GoAPI/io"
	"database/sql"
	"time"
	"github.com/d2r2/go-dht"
)

func main() {
	go collectData()

	r := gin.Default()
	r.Static("/web", "./web")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//PEW PEW
	r.PUT("/pew", sensors.LazorHandler(10).Pew(true))
	r.DELETE("/pew", sensors.LazorHandler(10).Pew(false))

	//IO
	r.PUT("/io/:pin", io.IoHandler().IoHigh())
	r.DELETE("/io/:pin", io.IoHandler().IoLow())

	//TEMPHU
	r.GET("/temp", sensors.TempHuHandler(9).GetTempHu())

	r.GET("/tempdata", sensors.TempHuHandler(9).GetTempList())
	r.GET("/hudata", sensors.TempHuHandler(9).GetHuList())

	//API REPORT
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.Default().Routes())
	})

	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}

func collectData() {
	db, err := sql.Open("sqlite3", "./temphu.db")

	if err != nil {
		return
	}
	for {
		temp, hu, _ , err := dht.ReadDHTxxWithRetry(dht.DHT11, 9, false, 5)

		if err != nil {
			continue
		}

		stmt, err := db.Prepare("INSERT INTO temphu(date, temp, hu) VALUES (DATETIME('now'), ?, ?)")

		if err != nil {
			continue
		}

		stmt.Exec(temp, hu)

		time.Sleep(1.8e+12)
	}
}