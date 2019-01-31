package controller

import (
	"fmt"
	"git-check-branch/api/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ServerController 服务器 操作
type ServerController struct {
	baseController
	serverModel model.ServerModel
}

// List 服务器列表
func (ctl *ServerController) List(c *gin.Context) {
	list := ctl.serverModel.All()
	var serverList = make([]model.ServerItem, 0)

	for _, server := range list {
		serverList = append(serverList, server)
	}
	ctl.Success(c, serverList, "list")
	return
}

// Add 添加服务器
func (ctl *ServerController) Add(c *gin.Context) {
	var server model.ServerItem
	if err := c.ShouldBind(&server); err != nil {
		fmt.Printf("%#v", server)
		ctl.Error(c, gin.H{}, "参数必填")
		return
	}
	lastID, err := ctl.serverModel.Add(server)
	if err != nil {
		ctl.Error(c, gin.H{"inset_id": lastID}, err.Error())
		return
	}
	ctl.Success(c, gin.H{"insert_id": lastID}, "添加成功")
	return
}

// Del 删除服务器
func (ctl *ServerController) Del(c *gin.Context) {
	var serverID, err = strconv.Atoi(c.Query("server_id"))
	if serverID <= 0 || err != nil {
		ctl.Error(c, gin.H{}, "server_id 必传")
		return
	}

	ctl.serverModel.Del(serverID)
	ctl.Success(c, gin.H{"server_id": serverID}, "删除成功")
	return
}

// Info 服务器信息
func (ctl *ServerController) Info(c *gin.Context) {
	var serverID, err = strconv.Atoi(c.Query("server_id"))
	if serverID <= 0 || err != nil {
		ctl.Error(c, gin.H{}, "server_id 必传")
		return
	}

	if info := ctl.serverModel.Info(serverID); info.ServerID != 0 {
		ctl.Success(c, info, "info")
		return
	}
	ctl.Error(c, gin.H{}, "配置不存在")
}

// Edit 服务器编辑
func (ctl *ServerController) Edit(c *gin.Context) {
	var server model.ServerItem
	fmt.Printf("%v", ctl.serverModel.Info(server.ServerID))
	if err := c.ShouldBind(&server); err != nil {
		ctl.Error(c, gin.H{}, "参数必填")
		return
	}
	if old := ctl.serverModel.Info(server.ServerID); old.ServerID == 0 {
		ctl.Error(c, gin.H{}, "该服务器配置不存在")
		return
	}
	_, err := ctl.serverModel.Edit(server.ServerID, server)
	if err != nil {
		ctl.Error(c, gin.H{}, err.Error())
		return
	}
	ctl.Success(c, gin.H{}, "编辑成功")
	return
}

// 测试连接状态
// func (ctl *serverController) testLink(c *gin.Context) {
// 	var serverIDStr = c.Query("server_id")
// 	if serverIDStr == "" {
// 		c.JSON(http.StatusOK, formatReturn(false, nil, "server_id 必传"))
// 		return
// 	}
// 	serverID, _ := strconv.Atoi(serverIDStr)

// 	server := getServerInfoByServerID(int(serverID))
// 	if server.IP == "" {
// 		c.JSON(http.StatusOK, formatReturn(false, nil, "该服务器不存在"))
// 		return
// 	}

// 	res, err := runCmd(server, "whoami")
// 	log.Printf("res; %s", res)
// 	if err != nil {
// 		c.JSON(http.StatusOK, formatReturn(false, err, ""))
// 		return
// 	}

// 	if server.User == strings.TrimSpace(string(res)) {
// 		c.JSON(http.StatusOK, formatReturn(true, string(res), "登录正常"))
// 		return
// 	}
// 	c.JSON(http.StatusOK, formatReturn(false, string(res), "登录测试失败"))
// }

// func getServerInfoByServerID(serverID int) *serverItem {
// 	server := new(serverItem)
// 	if serverID <= 0 {
// 		return server
// 	}

// 	err := db.QueryRow("SELECT server_id, name, ip, port, user, passwd FROM server WHERE server_id =?", serverID).Scan(&server.ServerID, &server.Name, &server.IP, &server.Port, &server.User, &server.Passwd)
// 	if err != nil {
// 		log.Printf("server get one sql error: %s", err)
// 	}
// 	return server
// }
