package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type testController struct {
	baseController
}

func (ctl *testController) Test(c *gin.Context) {
	req := c.Request
	req.ParseForm()
	f := req.Form
	log.Printf("%s", f)

	c.JSON(200, f)
	return
}
