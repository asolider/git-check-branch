package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type baseController struct{}

func (base *baseController) Success(c *gin.Context, data interface{}, message string) {
	res := gin.H{
		"status": true,
		"data":   data,
		"msg":    message,
	}

	c.JSON(http.StatusOK, res)
}

func (base *baseController) Error(c *gin.Context, data interface{}, message string) {
	if message == "" {
		message = "api出错"
	}
	res := gin.H{
		"status": false,
		"data":   data,
		"msg":    message,
	}

	c.JSON(http.StatusOK, res)
}

type formatData struct {
	Status bool        `json:"status" xml:"status"`
	Data   interface{} `json:"data" xml:"data"`
	Msg    string      `json:"msg" xml:"msg"`
}

func formatReturn(status bool, data interface{}, msg string) formatData {
	if status == false && msg == "" {
		msg = "系统错误"
	}
	if status == true && msg == "" {
		msg = "success"
	}
	return formatData{status, data, msg}
}

var db *sql.DB

func getDB() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/git_branch")
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(10 * time.Second)
	return db
}
