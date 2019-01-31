package controller

import (
	"fmt"
	"git-check-branch/api/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProjectController 服务器 操作
type ProjectController struct {
	baseController
	projectModel model.ProjectModel
	serverModel  model.ServerModel
}

// List 项目列表
func (ctl *ProjectController) List(c *gin.Context) {
	list := ctl.projectModel.All()

	type res struct {
		model.ProjectItem
		BelongServer model.ServerItem `json:"belong_server"`
	}

	var projectLst = make([]res, 0)

	for _, project := range list {
		serverInfo := ctl.serverModel.Info(project.ServerID)
		projectLst = append(projectLst, res{project, serverInfo})
	}
	ctl.Success(c, projectLst, "list")
	return
}

// Info 项目信息
func (ctl *ProjectController) Info(c *gin.Context) {
	var projectID, err = strconv.Atoi(c.Query("project_id"))
	if projectID <= 0 || err != nil {
		ctl.Error(c, gin.H{}, "server_id 必传")
		return
	}

	if info := ctl.projectModel.Info(projectID); info.ProjectID != 0 {
		ctl.Success(c, info, "info")
		return
	}
	ctl.Error(c, gin.H{}, "配置不存在")
}

// Add 添加项目
func (ctl *ProjectController) Add(c *gin.Context) {
	var project model.ProjectItem
	if err := c.ShouldBind(&project); err != nil {
		fmt.Printf("%#v", project)
		ctl.Error(c, gin.H{}, "参数填写错误"+err.Error())
		return
	}
	lastID, err := ctl.projectModel.Add(project)
	if err != nil {
		ctl.Error(c, gin.H{"insert_id": lastID}, err.Error())
		return
	}
	ctl.Success(c, gin.H{"insert_id": lastID}, "添加成功")
	return
}

// Edit 添加项目
func (ctl *ProjectController) Edit(c *gin.Context) {
	var project model.ProjectItem
	if err := c.ShouldBind(&project); err != nil {
		fmt.Printf("%#v", project)
		ctl.Error(c, gin.H{}, "参数填写错误"+err.Error())
		return
	}

	_, err := ctl.projectModel.Edit(project.ProjectID, project)
	if err != nil {
		ctl.Error(c, gin.H{}, err.Error())
		return
	}
	ctl.Success(c, gin.H{}, "编辑成功")
	return
}

// Del 删除项目配置
func (ctl *ProjectController) Del(c *gin.Context) {
	var projectID, err = strconv.Atoi(c.Query("project_id"))
	if projectID <= 0 || err != nil {
		ctl.Error(c, gin.H{}, "server_id 必传")
		return
	}

	ctl.projectModel.Del(projectID)
	ctl.Success(c, gin.H{}, "删除成功")
	return
}
