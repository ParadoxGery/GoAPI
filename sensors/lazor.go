package sensors

import (
	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio"
)

type lazorHandler struct {
	Pin int
	ioIsWorking bool
}

func (l lazorHandler) Pew(on bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !l.ioIsWorking {
			c.JSON(500, gin.H{
				"message": "gpio is not working",
			})

			return
		}
		pin := rpio.Pin(10)
		pin.Output()
		if on {
			pin.High()
		} else {
			pin.Low()
		}
	}
}

func LazorHandler(pin int) lazorHandler {
	err := rpio.Open()
	handler := lazorHandler{
		pin,
		err == nil,
	}

	return handler
}
