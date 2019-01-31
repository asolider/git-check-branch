package model

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// 存储文件结构
type db struct {
	MaxID   int                 `yaml:"maxid"`
	Server  map[int]ServerItem  `yaml:"server"`
	Project map[int]ProjectItem `yaml:"project"`
}

// 数据文件保存位置
var dataFileLocation string
var dataFileBakLocation string

// InitModel 初始化数据文件的配置
func InitModel(datafile string) {
	dataFileLocation = datafile
	dataFileBakLocation = datafile + ".bak"
}

// 获取保存在文件上的数据
func getFileContent() db {
	var fileData = db{Server: make(map[int]ServerItem), Project: make(map[int]ProjectItem)}

	file, _ := os.OpenFile(dataFileLocation, os.O_RDONLY|os.O_CREATE, 0666)
	defer file.Close()

	bts, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("read dbfile err,", err)
		return fileData
	}

	err = yaml.Unmarshal(bts, &fileData)
	if err != nil {
		log.Printf("Unmarshal dbfile err,", err)
	}
	return fileData
}

// 重写数据到文件上
func saveFileContent(data db) error {
	bakFile()

	file, _ := os.OpenFile(dataFileLocation, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		recoverBak()
		log.Printf("marshal dbfile err,", err)
		return err
	}

	_, err = file.Write(yamlData)
	if err != nil {
		recoverBak()
		log.Printf("write dbfile err,", err)
		return err
	}
	return nil
}

// 写文件之前先保存个备份，以防意外
func bakFile() {
	// 把之前备份的给删了
	os.Remove(dataFileBakLocation)
	os.Rename(dataFileLocation, dataFileBakLocation)
}

// 恢复备份文件
func recoverBak() {
	os.Rename(dataFileBakLocation, dataFileLocation)
}
