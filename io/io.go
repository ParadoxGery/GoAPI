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
	if !i.isIoWorking {
		return func(c *gin.Context) {
			c.JSON(500,  gin.H{
				"message": "error",
				"error": "gpio is not working",
			})
		}
	}

	return func(c *gin.Context) {
		pinNum, err := strconv.Atoi(c.Param("pin"))

		if err != nil {
			c.JSON(500, gin.H{
				"message": "error",
				"error": err.Error(),
			})
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
	if !i.isIoWorking {
		return func(c *gin.Context) {
			c.JSON(500,  gin.H{
				"message": "error",
				"error": "gpio is not working",
			})
		}
	}

	return func(c *gin.Context) {
		pinNum, err := strconv.Atoi(c.Param("pin"))

		if err != nil {
			c.JSON(500, gin.H{
				"message": "error",
				"error": err.Error(),
			})
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