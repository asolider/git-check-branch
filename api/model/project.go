package model

// 项目类型的声明，及字段
type ProjectItem struct {
	ProjectID    int        `json:"project_id" form:"project_id"`
	ServerID     int        `json:"server_id" form:"server_id" binding:"required"`
	Name         string     `json:"name" form:"name" binding:"required"`
	Location     string     `json:"location" form:"location" binding:"required"`
	CurBranch    string     `json:"cur_branch" form:"cur_branch"`
	BelongServer ServerItem `json:"belong_server"`
}

type projectModel struct{}

func (project *projectModel) info() {}
func (project *projectModel) add()  {}
func (project *projectModel) edit() {}
func (project *projectModel) del()  {}
func (project *projectModel) all()  {}
