package main

import (
	"github.com/gin-gonic/gin"
	"GoAPI/sensors"
)

func main() {
	r := gin.Default()
	r.Static("/web", "./web")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.PUT("/pew", sensors.LazorHandler(10).Pew(true))
	r.DELETE("/pew", sensors.LazorHandler(10).Pew(false))

	r.GET("/temp", sensors.TempHuHandler(9).GetTempHu())

	r.GET("/temps", func(c *gin.Context) {
		c.JSON(200, [][]int{ {1, 2}, {3, 5}, {5, 13}, {1, 42}, {3, 45}, {5,53} })
	})
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.Default().Routes())
	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}