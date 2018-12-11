package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type serverItem struct {
	ServerID int    `json:"server_id" form:"server_id"`
	Name     string `json:"name" form:"server_name" binding:"required"`
	IP       string `json:"ip" form:"server_ip" binding:"required"`
	Port     int    `json:"port" form:"server_port" binding:"required"`
	User     string `json:"user" form:"server_user" binding:"required"`
	Passwd   string `json:"passwd" form:"server_passwd" binding:"required"`
}

func serverList(c *gin.Context) {
	var servers []serverItem
	rows, err := db.Query("SELECT server_id, name, ip, port, user, passwd FROM server")

	if err != nil {
		log.Printf("fail to query server list:%s", err)
		c.JSON(http.StatusOK, formatReturn(false, err, ""))
		return
	}
	defer rows.Close()
	server := serverItem{}
	for rows.Next() {
		err := rows.Scan(&server.ServerID, &server.Name, &server.IP, &server.Port, &server.User, &server.Passwd)
		if err != nil {
			log.Printf("fail to scan server:%s", err)
			c.JSON(http.StatusOK, formatReturn(false, nil, ""))
			return
		}
		servers = append(servers, server)
	}
	if err := rows.Err(); err != nil {
		log.Printf("fail to get:%s", err)
	}

	c.JSON(http.StatusOK, formatReturn(true, servers, ""))
	return
}

func serverAdd(c *gin.Context) {
	var server serverItem
	if err := c.ShouldBind(&server); err != nil {
		c.JSON(http.StatusOK, formatReturn(false, err, "参数都必须填写"))
		return
	}

	res, err := db.Exec("INSERT INTO server (name, ip, port, user, passwd) VALUES (?,?,?,?,?)", server.Name, server.IP, server.Port, server.User, server.Passwd)
	if err != nil {
		log.Printf("add server fail, err", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Printf("get add server_id fail, err:", err)
	}
	c.JSON(http.StatusOK, formatReturn(true, map[string]int64{"insert_id": lastID}, ""))
	return
}

func serverDel(c *gin.Context) {
	var serverID = c.Query("server_id")
	if serverID == "" {
		c.JSON(http.StatusOK, formatReturn(false, nil, "id 必传"))
		return
	}

	_, err := db.Exec("DELETE FROM server WHERE server_id = ?", serverID)
	if err != nil {
		log.Printf("del server fail, err", err)
	}

	c.JSON(http.StatusOK, formatReturn(true, nil, "删除成功"))
	return
}

func serverInfo(c *gin.Context) {
	var serverID = c.Query("server_id")
	if serverID == "" {
		c.JSON(http.StatusOK, formatReturn(false, nil, "id 必传"))
		return
	}

	var server serverItem
	err := db.QueryRow("SELECT server_id, name, ip, port, user, passwd FROM server WHERE server_id =?", serverID).Scan(&server.ServerID, &server.Name, &server.IP, &server.Port, &server.User, &server.Passwd)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, formatReturn(false, nil, "记录不存在"))
			return
		} else {
			log.Printf("fail to query server e:%s id=%s", err, serverID)
			c.JSON(http.StatusOK, formatReturn(false, err, ""))
			return
		}
	}
	c.JSON(http.StatusOK, formatReturn(true, server, ""))
}

func serverEdit(c *gin.Context) {
	var server serverItem

	if err := c.ShouldBind(&server); err != nil {
		c.JSON(http.StatusOK, formatReturn(false, err, "参数都必须填写"))
		return
	}
	if server.ServerID == 0 {
		c.JSON(http.StatusOK, formatReturn(false, nil, "id不能为空"))
		return
	}

	_, err := db.Exec("UPDATE server SET name=?, ip=?, port=?, user=?, passwd=? WHERE server_id=?", server.Name, server.IP, server.Port, server.User, server.Passwd, server.ServerID)
	if err != nil {
		log.Printf("update server fail, err %s; id=%d", err, server.ServerID)
		c.JSON(http.StatusOK, formatReturn(false, nil, "更新失败"))
		return
	}

	c.JSON(http.StatusOK, formatReturn(true, nil, "更新成功"))
	return
}

func testLink(c *gin.Context) {
	var serverIDStr = c.Query("server_id")
	if serverIDStr == "" {
		c.JSON(http.StatusOK, formatReturn(false, nil, "server_id 必传"))
		return
	}
	serverID, _ := strconv.Atoi(serverIDStr)

	server := getServerInfoByServerID(int(serverID))
	if server.IP == "" {
		c.JSON(http.StatusOK, formatReturn(false, nil, "该服务器不存在"))
		return
	}

	res, err := runCmd(server, "whoami")
	log.Printf("res; %s", res)
	if err != nil {
		c.JSON(http.StatusOK, formatReturn(false, err, ""))
		return
	}

	if server.User == strings.TrimSpace(string(res)) {
		c.JSON(http.StatusOK, formatReturn(true, string(res), "登录正常"))
		return
	}
	c.JSON(http.StatusOK, formatReturn(false, string(res), "登录测试失败"))
}

func getServerInfoByServerID(serverID int) *serverItem {
	server := new(serverItem)
	if serverID <= 0 {
		return server
	}

	err := db.QueryRow("SELECT server_id, name, ip, port, user, passwd FROM server WHERE server_id =?", serverID).Scan(&server.ServerID, &server.Name, &server.IP, &server.Port, &server.User, &server.Passwd)
	if err != nil {
		log.Printf("server get one sql error: %s", err)
	}
	return server
}
