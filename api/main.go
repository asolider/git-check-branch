package main

import (
	"git-check-branch/api/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	// 服务器管理
	serverGroup := r.Group("/server")
	{
		serverController := new(controller.ServerController)
		serverGroup.GET("/list", serverController.List)
		serverGroup.POST("/add", serverController.Add)
		serverGroup.POST("/edit", serverController.Edit)
		serverGroup.GET("/del", serverController.Del)
		serverGroup.GET("/info", serverController.Info)
		// serverGroup.GET("/test", testLink)

	}

	// // 项目管理

	projectGroup := r.Group("/project")
	{
		projectController := new(controller.ProjectController)
		projectGroup.GET("/list", projectController.List)
		projectGroup.GET("/info", projectController.Info)
		projectGroup.GET("/del", projectController.Del)
		projectGroup.POST("/add", projectController.Add)
		projectGroup.POST("/edit", projectController.Edit)
	}

	// // 分支管理
	// branchGroup := r.Group("/branch")
	// {

	// 	branchGroup.GET("/list/:projectId", branchList)
	// 	branchGroup.GET("/check/:projectId", branchCheck)
	// 	branchGroup.GET("/refresh", refreshBranch)
	// }

	r.Run(":8080")
}
