package model

// ServerItem 服务器类型的声明，及字段
type ServerItem struct {
	ServerID int    `json:"server_id" form:"server_id"`
	Name     string `json:"name" form:"server_name" binding:"required"`
	IP       string `json:"ip" form:"server_ip" binding:"required"`
	Port     int    `json:"port" form:"server_port" binding:"required"`
	User     string `json:"user" form:"server_user" binding:"required"`
	Passwd   string `json:"passwd" form:"server_passwd" binding:"required"`
}

// ServerModel 服务器配置操作model
type ServerModel struct{}

// Info 服务器详情
func (model *ServerModel) Info(serverID int) ServerItem {
	dbfile := getFileContent()
	return dbfile.Server[serverID]
}

// Add 添加
func (model *ServerModel) Add(item ServerItem) (insertID int, err error) {
	dbfile := getFileContent()
	item.ServerID = dbfile.MaxID + 1
	dbfile.MaxID = dbfile.MaxID + 1
	dbfile.Server[item.ServerID] = item
	err = saveFileContent(dbfile)
	if err != nil {
		return 0, err
	}
	return item.ServerID, nil
}

// Edit 编辑
func (model *ServerModel) Edit(serverID int, item ServerItem) (int, error) {
	dbfile := getFileContent()
	dbfile.Server[serverID] = item
	err := saveFileContent(dbfile)
	if err != nil {
		return serverID, err
	}
	return serverID, nil
}

// Del 删除一个服务器配置
func (model *ServerModel) Del(serverID int) (int, error) {
	dbfile := getFileContent()
	delete(dbfile.Server, serverID)
	err := saveFileContent(dbfile)
	if err != nil {
		return serverID, err
	}
	return serverID, nil
}

// All 列表
func (model *ServerModel) All() map[int]ServerItem {
	dbfile := getFileContent()
	return dbfile.Server
}
