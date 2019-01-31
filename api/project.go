package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type projectItem struct {
	ProjectID    int        `json:"project_id" form:"project_id"`
	ServerID     int        `json:"server_id" form:"server_id" binding:"required"`
	Name         string     `json:"name" form:"name" binding:"required"`
	Location     string     `json:"location" form:"location" binding:"required"`
	CurBranch    string     `json:"cur_branch" form:"cur_branch"`
	BelongServer serverItem `json:"belong_server"`
}

type projectController struct {
	baseController
}

func (ctl *projectController) projectList(c *gin.Context) {
	var projects []projectItem
	rows, err := db.Query("SELECT project_id, project.server_id, project.name, location, cur_branch, server.name as server_name FROM project LEFT JOIN server ON project.server_id=server.server_id")

	if err != nil {
		log.Printf("fail to query project list:%s", err)
		c.JSON(http.StatusOK, formatReturn(false, err, ""))
		return
	}
	defer rows.Close()
	project := projectItem{BelongServer: serverItem{}}
	for rows.Next() {
		err := rows.Scan(&project.ProjectID, &project.ServerID, &project.Name, &project.Location, &project.CurBranch, &project.BelongServer.Name)
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
	log.Print(projects)
	ctl.Success(c, projects, "")
	//c.JSON(http.StatusOK, formatReturn(true, projects, ""))
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

func getProjectInfoByProjectID(projectID int) *projectItem {
	project := new(projectItem)
	if projectID <= 0 {
		return project
	}
	log.Printf("projectId: %d", projectID)

	err := db.QueryRow("SELECT project_id, server_id, name, location FROM project WHERE project_id =?", projectID).Scan(&project.ProjectID, &project.ServerID, &project.Name, &project.Location)
	if err != nil {
		log.Printf("project get one sql error: %s", err)
	}
	return project
}
