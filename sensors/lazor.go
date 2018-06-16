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
	if !l.ioIsWorking {
		return func(c *gin.Context) {
			c.JSON(500, gin.H{
				"message": "gpio is not working",
			})
		}
	}
	return func(c *gin.Context) {
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
