package sensors

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stianeikeland/go-rpio"
	"github.com/d2r2/go-dht"
	"database/sql"
	"log"
	"strconv"
)

type tempHuHandler struct {
	Pin int
	ioIsWorking bool
}

func (t tempHuHandler) GetTempHu() gin.HandlerFunc {

	return func(c *gin.Context) {
		if !t.ioIsWorking {
			c.JSON(500, gin.H{
				"message": "error",
				"error": "gpio is not working",
			})
			return
		}
		temp, hum, _, err := dht.ReadDHTxxWithRetry(dht.DHT11, 9, false, 2)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "error",
				"error" : err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"temp": temp,
			"hum" : hum,
		})
	}
}

func (t tempHuHandler) GetTempList() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("sqlite3", "./temphu.db")

		if err != nil {
			c.JSON(500, gin.H{
				"message": "error",
				"error": err.Error(),
			})
			return
		}

		rows, err := db.Query("SELECT date, temp FROM temphu WHERE DATETIME(date) BETWEEN DATETIME('now', '-1 day') AND DATETIME('now');")

		if err != nil {
			log.Fatal(err.Error())
		}
		var temps = "[["
		for rows.Next() {
			var date string
			var temp int
			err := rows.Scan(&date, &temp)

			if err != nil {
				//TODO error
			}

			temps += "[\""+date+"\","+strconv.Itoa(temp)+"],"
		}

		l := len(temps)
		temps = temps[:l-1]

		temps += "]]"
		//c.JSON(200, temps)
		c.Data(200, "application/json", []byte(temps))
	}
}

func (t tempHuHandler) GetHuList() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("sqlite3", "./temphu.db")

		if err != nil {
			c.JSON(500, gin.H{
				"message": "error",
				"error": err.Error(),
			})
			return
		}

		rows, err := db.Query("SELECT date, hu FROM temphu WHERE DATE(date) BETWEEN DATETIME('now', '-1 day') AND DATETIME('now');")

		if err != nil {
			log.Fatal(err.Error())
		}
		var hus = "[["
		for rows.Next() {
			var date string
			var hu int
			err := rows.Scan(&date, &hu)

			if err != nil {
				//TODO error
			}

			hus += "[\""+date+"\","+strconv.Itoa(hu)+"],"
		}

		l := len(hus)
		hus = hus[:l-1]

		hus += "]]"
		//c.JSON(200, hus)
		c.Data(200, "application/json", []byte(hus))
	}
}

func TempHuHandler(pin int) tempHuHandler {
	err := rpio.Open()
	handler := tempHuHandler{
		pin,
		err == nil,
	}

	return handler
}
