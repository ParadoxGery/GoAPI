package io

import (
	"github.com/stianeikeland/go-rpio"
	"github.com/gin-gonic/gin"
	"strconv"
	"log"
)

type ioHandler struct {
	isIoWorking bool
}

func (i ioHandler) IoHigh() gin.HandlerFunc {
	if !i.isIoWorking {
		return func(c *gin.Context) {
			c.JSON(500,  gin.H{
				"message": "gpio is not working",
			})
		}
	}

	return func(c *gin.Context) {
		pinNum, err := strconv.Atoi(c.Param("pin"))

		if err != nil {
			log.Fatal("invalid param")
		}

		pin := rpio.Pin(pinNum)
		pin.Output()
		pin.High()
	}
}

func IoHandler() ioHandler {
	err := rpio.Open()

	handler := ioHandler{
		err == nil,
	}

	return handler
}