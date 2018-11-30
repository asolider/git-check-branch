package main

import (
	"github.com/gin-gonic/gin"
)

func init() {
	db = getDB()
}

func main() {
	r := gin.Default()

	serverGroup := r.Group("/server")
	{
		serverGroup.GET("/list", serverList)
		serverGroup.POST("/add", serverAdd)
		serverGroup.GET("/del", serverDel)
		serverGroup.GET("/info", serverInfo)
		serverGroup.POST("/edit", serverEdit)
	}

	projectGroup := r.Group("/project")
	{
		projectGroup.GET("/list", serverList)
		// projectGroup.POST("/add", posting)
		// projectGroup.POST("/edit", posting)
		// projectGroup.GET("/del", posting)
	}

	r.Run(":8080")
}
