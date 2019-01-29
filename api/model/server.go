package model

// 服务器类型的声明，及字段
type serverItem struct {
	ServerID int    `json:"server_id" form:"server_id"`
	Name     string `json:"name" form:"server_name" binding:"required"`
	IP       string `json:"ip" form:"server_ip" binding:"required"`
	Port     int    `json:"port" form:"server_port" binding:"required"`
	User     string `json:"user" form:"server_user" binding:"required"`
	Passwd   string `json:"passwd" form:"server_passwd" binding:"required"`
}
