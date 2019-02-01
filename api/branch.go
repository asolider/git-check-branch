package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"sync"

// 	"github.com/gin-gonic/gin"
// )

// // 项目下分支列表
// func branchList(c *gin.Context) {
// 	projectID, _ := strconv.Atoi(c.Param("projectId"))
// 	projectInfo := getProjectInfoByProjectID(projectID)
// 	serverInfo := getServerInfoByServerID(projectInfo.ServerID)

// 	if projectInfo.ProjectID == 0 || serverInfo.ServerID == 0 {
// 		c.JSON(http.StatusOK, formatReturn(false, nil, "信息异常"))
// 		return
// 	}
// 	cmd := fmt.Sprintf("cd %s && git branch -r ", projectInfo.Location)
// 	resBytes, err := runCmd(serverInfo, cmd)
// 	if err != nil {
// 		log.Printf("查看分支列表出错：%s", err)
// 		c.JSON(http.StatusOK, formatReturn(false, err, "执行出错"))
// 		return
// 	}
// 	// 第一行为 origin/HEAD
// 	formatRes := formatCmdReturn(resBytes)
// 	res := make([]string, 0)
// 	for _, v := range formatRes {
// 		if strings.HasPrefix(v, "origin/HEAD") {
// 			continue
// 		}
// 		res = append(res, strings.TrimPrefix(v, "origin/"))
// 	}
// 	c.JSON(http.StatusOK, formatReturn(true, gin.H{"cmd": cmd, "branch_list": res}, "success"))
// }

// // 切换项目的分支
// func branchCheck(c *gin.Context) {
// 	projectID, _ := strconv.Atoi(c.Param("projectId"))
// 	var branchName = c.Query("selectbranch")
// 	if branchName == "" {
// 		c.JSON(http.StatusOK, formatReturn(false, nil, "未选择分支"))
// 		return
// 	}

// 	projectInfo := getProjectInfoByProjectID(projectID)
// 	serverInfo := getServerInfoByServerID(projectInfo.ServerID)

// 	if projectInfo.ProjectID == 0 || serverInfo.ServerID == 0 {
// 		c.JSON(http.StatusOK, formatReturn(false, nil, "信息异常"))
// 		return
// 	}
// 	cmd := fmt.Sprintf("cd %s && git fetch && git checkout %s ", projectInfo.Location, branchName)
// 	resBytes, err := runCmd(serverInfo, cmd)
// 	if err != nil {
// 		log.Printf("切换分支出错：%s", err)
// 		c.JSON(http.StatusOK, formatReturn(false, gin.H{"cmd": cmd, "output": formatCmdReturn(resBytes)}, "执行出错"+err.Error()))
// 		return
// 	}
// 	c.JSON(http.StatusOK, formatReturn(true, gin.H{"cmd": cmd, "output": formatCmdReturn(resBytes)}, "success"))
// }

// // 刷新目前的分支情况，更新数据库
// func refreshBranch(c *gin.Context) {
// 	// get project list
// 	projectIds := make([]int, 0)
// 	rows, err := db.Query("SELECT project_id FROM project")

// 	if err != nil {
// 		log.Printf("fail to query projectid list:%s", err)
// 		c.JSON(http.StatusOK, formatReturn(false, err, ""))
// 		return
// 	}
// 	defer rows.Close()
// 	projectID := 0
// 	for rows.Next() {
// 		err := rows.Scan(&projectID)
// 		if err != nil {
// 			log.Printf("fail to scan project:%s", err)
// 			c.JSON(http.StatusOK, formatReturn(false, nil, ""))
// 			return
// 		}
// 		projectIds = append(projectIds, projectID)
// 	}
// 	if err := rows.Err(); err != nil {
// 		log.Printf("fail to get:%s", err)
// 		rows.Close()
// 	}

// 	if len(projectIds) == 0 {
// 		c.JSON(http.StatusOK, formatReturn(false, projectIds, "项目列表为空"))
// 		return
// 	}

// 	var wg sync.WaitGroup
// 	for _, v := range projectIds {
// 		wg.Add(1)
// 		go func(projectID int) {
// 			defer wg.Done()
// 			updateCurrentBranch(projectID)
// 		}(v)
// 	}
// 	wg.Wait()

// 	c.JSON(http.StatusOK, formatReturn(true, projectIds, ""))
// 	return

// }

// func updateCurrentBranch(projectID int) bool {
// 	projectInfo := getProjectInfoByProjectID(projectID)
// 	serverInfo := getServerInfoByServerID(projectInfo.ServerID)

// 	if projectInfo.ProjectID == 0 || serverInfo.ServerID == 0 {
// 		return false
// 	}
// 	cmd := fmt.Sprintf("cd %s && git branch", projectInfo.Location)
// 	resBytes, err := runCmd(serverInfo, cmd)
// 	if err != nil {
// 		log.Printf("服务器 %s 查看当前分支出错：%s", serverInfo.Name, err)
// 		return false
// 	}
// 	curBranch := strings.Trim(strings.TrimSpace(string(resBytes)), "* ")

// 	_, err = db.Exec("UPDATE project SET cur_branch=? WHERE project_id=?", curBranch, projectID)
// 	if err != nil {
// 		log.Printf("update project branch fail, err %s; id=%d", err, projectID)
// 		return false
// 	}
// 	return true
// }
