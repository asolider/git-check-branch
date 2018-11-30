package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type projectItem struct {
	ProjectID int    `json:"project_id" form:"project_id"`
	ServerID  int    `json:"server_id" form:"server_id" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required"`
	Location  string `json:"location" form:"location" binding:"required"`
	CurBranch string `json:"cur_branch" form:"cur_branch"`
}

func projectList(c *gin.Context) {
	var projects []projectItem
	rows, err := db.Query("SELECT project_id, server_id, name, location FROM project")

	if err != nil {
		log.Printf("fail to query project list:%s", err)
		c.JSON(http.StatusOK, formatReturn(false, err, ""))
		return
	}
	defer rows.Close()
	project := projectItem{}
	for rows.Next() {
		err := rows.Scan(&project.ProjectID, &project.ServerID, &project.Name, &project.Location)
		if err != nil {
			log.Printf("fail to scan project:%s", err)
			c.JSON(http.StatusOK, formatReturn(false, nil, ""))
			return
		}
		projects = append(projects, project)
	}
	if err := rows.Err(); err != nil {
		log.Printf("fail to get:%s", err)
	}

	c.JSON(http.StatusOK, formatReturn(true, projects, ""))
	return
}

func projectAdd(c *gin.Context) {
	var project projectItem
	if err := c.ShouldBind(&project); err != nil {
		c.JSON(http.StatusOK, formatReturn(false, err, "参数都必须填写"))
		return
	}

	res, err := db.Exec("INSERT INTO project (server_id, name, location) VALUES (?,?,?)", project.ServerID, project.Name, project.Location)
	if err != nil {
		log.Printf("add project fail, err", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Printf("get add server_id fail, err:", err)
	}
	c.JSON(http.StatusOK, formatReturn(true, map[string]int64{"insert_id": lastID}, ""))
	return
}

func projectDel(c *gin.Context) {
	var projectID = c.Query("project_id")
	if projectID == "" {
		c.JSON(http.StatusOK, formatReturn(false, nil, "id 必传"))
		return
	}

	_, err := db.Exec("DELETE FROM project WHERE project_id = ?", projectID)
	if err != nil {
		log.Printf("del project fail, err", err)
	}

	c.JSON(http.StatusOK, formatReturn(true, nil, "删除成功"))
	return
}

func projectInfo(c *gin.Context) {
	var projectID = c.Query("project_id")
	if projectID == "" {
		c.JSON(http.StatusOK, formatReturn(false, nil, "id 必传"))
		return
	}

	project := projectItem{}
	err := db.QueryRow("SELECT project_id, server_id, name, location FROM project WHERE project_id =?", projectID).Scan(&project.ProjectID, &project.ServerID, &project.Name, &project.Location)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, formatReturn(false, nil, "记录不存在"))
			return
		} else {
			log.Printf("fail to query project e:%s id=%s", err, projectID)
			c.JSON(http.StatusOK, formatReturn(false, err, ""))
			return
		}
	}
	c.JSON(http.StatusOK, formatReturn(true, project, ""))
}

func projectEdit(c *gin.Context) {
	project := projectItem{}

	if err := c.ShouldBind(&project); err != nil {
		c.JSON(http.StatusOK, formatReturn(false, err, "参数都必须填写"))
		return
	}
	if project.ProjectID == 0 {
		c.JSON(http.StatusOK, formatReturn(false, nil, "id不能为空"))
		return
	}

	_, err := db.Exec("UPDATE project SET server_id=?, name=?, location=? WHERE project_id=?", project.ServerID, project.Name, project.Location, project.ProjectID)
	if err != nil {
		log.Printf("update project fail, err %s; id=%d", err, project.ProjectID)
		c.JSON(http.StatusOK, formatReturn(false, nil, "更新失败"))
		return
	}

	c.JSON(http.StatusOK, formatReturn(true, nil, "更新成功"))
	return
}
