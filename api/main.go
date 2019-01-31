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
	// projectController := new(projectController)
	// projectGroup := r.Group("/project")
	// {
	// 	projectGroup.GET("/list", projectController.projectList)
	// 	projectGroup.GET("/info", projectInfo)
	// 	projectGroup.GET("/del", projectDel)
	// 	projectGroup.POST("/add", projectAdd)
	// 	projectGroup.POST("/edit", projectEdit)
	// }

	// // 分支管理
	// branchGroup := r.Group("/branch")
	// {

	// 	branchGroup.GET("/list/:projectId", branchList)
	// 	branchGroup.GET("/check/:projectId", branchCheck)
	// 	branchGroup.GET("/refresh", refreshBranch)
	// }

	// testController := new(testController)
	// r.GET("/test", testController.Test)
	// r.POST("/test", testController.Test)

	// // socket
	// r.GET("/ping", ping)

	r.Run(":8080")
}
