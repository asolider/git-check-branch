package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type baseController struct{}

func (base *baseController) Success(c *gin.Context, data interface{}, message string) {
	res := gin.H{
		"status": true,
		"data":   data,
		"msg":    message,
	}

	c.JSON(http.StatusOK, res)
}

func (base *baseController) Error(c *gin.Context, data interface{}, message string) {
	if message == "" {
		message = "api出错"
	}
	res := gin.H{
		"status": false,
		"data":   data,
		"msg":    message,
	}

	c.JSON(http.StatusOK, res)
}

type formatData struct {
	Status bool        `json:"status" xml:"status"`
	Data   interface{} `json:"data" xml:"data"`
	Msg    string      `json:"msg" xml:"msg"`
}
