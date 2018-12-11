package main

import "github.com/gin-gonic/gin"

type testController struct {
	baseController
}

type a struct {
	Age []int
}

type b struct {
	Name string
	a
}

func (ctl *testController) Test(c *gin.Context) {
	//abc := gin.H{}
	//ctl.Error(c, abc, "abc")
	abc := b{Name: "hello", a: a{Age: []int{1, 2}}}
	ctl.Error(c, abc, "abc")
	return

	c.JSON(200, abc)
	return
}
