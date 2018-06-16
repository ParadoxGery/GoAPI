package main

import (
	"github.com/gin-gonic/gin"
	"GoAPI/sensors"
	"GoAPI/io"
)

func main() {
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

	//TEMPDATA
	r.GET("/temps", func(c *gin.Context) {
		//c.JSON(200, [][][]int{ { {1, 2}, {3, 5}, {5, 13}, {1, 42}, {3, 45}, {5,53} } })
		c.JSON(200, [][]gin.H{ { {"a": 2}, {"v": 5}, {"x": 13}, {"2": 42}, {"aa": 45}, {"5":53} } })
	})

	r.GET("/tempdata", sensors.TempHuHandler(9).GetTempList())
	r.GET("/hudata", sensors.TempHuHandler(9).GetHuList())

	//API REPORT
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.Default().Routes())
	})

	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}