package controller

import (
	"git-check-branch/api/model"

	"github.com/gin-gonic/gin"
)

// BranchController 服务器 操作
type BranchController struct {
	baseController
	projectModel model.ProjectModel
	serverModel  model.ServerModel
}

// List 项目列表
func (ctl *BranchController) List(c *gin.Context) {}

// Check 切换分支
func (ctl *BranchController) Check(c *gin.Context) {}

// Refresh 刷新分支
func (ctl *BranchController) Refresh(c *gin.Context) {}
