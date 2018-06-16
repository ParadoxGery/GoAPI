package sensors

import (
	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio"
	"github.com/d2r2/go-dht"
)

type tempHuHandler struct {
	Pin int
	ioIsWorking bool

}

func (t tempHuHandler) GetTempHu() gin.HandlerFunc {
	if !t.ioIsWorking {
		return func(c *gin.Context) {
			c.JSON(500, gin.H{
				"message": "gpio is not working",
			})
		}
	}

	return func(c *gin.Context) {
		temp, hum, _, err := dht.ReadDHTxxWithRetry(dht.DHT11, 9, false, 2)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"temp": temp,
			"hum" : hum,
		})
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
