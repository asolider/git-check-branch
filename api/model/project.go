package model

// ProjectItem 项目类型的声明，及字段
type ProjectItem struct {
	ProjectID int    `json:"project_id" form:"project_id"`
	ServerID  int    `json:"server_id" form:"server_id" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required"`
	Location  string `json:"location" form:"location" binding:"required"`
	CurBranch string `json:"cur_branch" form:"cur_branch"`
}

// ProjectModel 项目操作model
type ProjectModel struct{}

// Info 详情
func (model *ProjectModel) Info(projectID int) ProjectItem {
	dbfile := getFileContent()
	return dbfile.Project[projectID]
}

// Add 增加
func (model *ProjectModel) Add(item ProjectItem) (insertID int, err error) {
	dbfile := getFileContent()
	item.ProjectID = dbfile.MaxID + 1
	dbfile.MaxID = dbfile.MaxID + 1
	dbfile.Project[item.ProjectID] = item
	err = saveFileContent(dbfile)
	if err != nil {
		return 0, err
	}
	return item.ProjectID, nil
}

// Edit 编辑
func (model *ProjectModel) Edit(projectID int, item ProjectItem) (int, error) {
	dbfile := getFileContent()

	dbfile.Project[projectID] = item
	err := saveFileContent(dbfile)
	if err != nil {
		return projectID, err
	}
	return projectID, nil
}

// Del 删除
func (model *ProjectModel) Del(projectID int) (int, error) {
	dbfile := getFileContent()
	delete(dbfile.Project, projectID)
	err := saveFileContent(dbfile)
	if err != nil {
		return projectID, err
	}
	return projectID, nil
}

// All 列表
func (model *ProjectModel) All() map[int]ProjectItem {
	dbfile := getFileContent()
	return dbfile.Project
}
