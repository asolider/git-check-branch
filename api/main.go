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
		projectGroup.GET("/list", projectList)
		projectGroup.GET("/info", projectInfo)
		projectGroup.GET("/del", projectDel)
		projectGroup.POST("/add", projectAdd)
		projectGroup.POST("/edit", projectEdit)
	}

	r.Run(":8080")
}