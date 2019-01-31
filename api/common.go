package main

import (
	"git-check-branch/api/model"
	"os"
	"path/filepath"
)

// 项目配置
var appConf = make(map[string]string)

// 项目执行目录
var appRoot string

func init() {
	appRoot, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	appConf["dbfile_location"] = "./data/db.yaml"

	// 数据文件初始化位置
	model.InitModel(filepath.Join(appRoot, appConf["dbfile_location"]))

}
