package io

import (
	"github.com/stianeikeland/go-rpio"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ioHandler struct {
	isIoWorking bool
}

func (i ioHandler) IoHigh() gin.HandlerFunc {

	return func(c *gin.Context) {
		if !i.isIoWorking {
			c.JSON(500,  gin.H{
				"message": "error",
				"error": "gpio is not working",
			})
			return
		}
		pinNum, err := strconv.Atoi(c.Param("pin"))

		if err != nil {
			c.JSON(500, gin.H{
				"message": "error",
				"error": err.Error(),
			})

			return
		}

		pin := rpio.Pin(pinNum)
		pin.Output()
		pin.High()
		c.JSON(200, gin.H{
			"message": "success",
			"pin": pinNum,
		})
	}
}

func (i ioHandler) IoLow() gin.HandlerFunc {

	return func(c *gin.Context) {
		if !i.isIoWorking {
			c.JSON(500,  gin.H{
				"message": "error",
				"error": "gpio is not working",
			})
			return
		}
		pinNum, err := strconv.Atoi(c.Param("pin"))

		if err != nil {
			c.JSON(500, gin.H{
				"message": "error",
				"error": err.Error(),
			})

			return
		}

		pin := rpio.Pin(pinNum)
		pin.Output()
		pin.Low()
		c.JSON(200, gin.H{
			"message": "success",
			"pin": pinNum,
		})
	}
}

func IoHandler() ioHandler {
	err := rpio.Open()

	handler := ioHandler{
		err == nil,
	}

	return handler
}