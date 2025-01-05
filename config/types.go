package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type S3Item struct {
	SecretId   string `json:"secretId"`
	SecretKey  string `json:"secretKey"`
	Endpoint   string `json:"endpoint"`
	BucketName string `json:"bucketName"`
	Region     string `json:"region"`
}

type DbItem struct {
	Host      string   `json:"host"`
	Port      string   `json:"port"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Databases []string `json:"databases"`
}

type Config struct {
	Cron string   `json:"cron"`
	Db   []DbItem `json:"db"`
	S3   []S3Item `json:"s3"`
}

type DBConfig struct {
	Host       string `gorm:"column:host" json:"host"`
	Port       string `gorm:"column:port" json:"port"`
	Username   string `gorm:"column:username" json:"username"`
	Password   string `gorm:"column:password" json:"password"`
	Cron       string `gorm:"column:cron" json:"cron"`
	HostPath   string `gorm:"column:hostPath" json:"hostPath"`
	SecretID   string `gorm:"column:secretID" json:"secretID"`
	SecretKey  string `gorm:"column:secretKey" json:"secretKey"`
	Endpoint   string `gorm:"column:endpoint" json:"endpoint"`
	BucketName string `gorm:"column:bucketName" json:"bucketName"`
	Region     string `gorm:"column:region" json:"region"`
}

func GetConfig() (conf *Config, err error) {
	root, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	configFilePath := filepath.Join(root, "./config.json")
	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("读取配置失败，请检查配置文件是否存在", err)
		return nil, err
	}
	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("反序列化配置失败，请检查配置文件，参考README配置说明", err)
		return nil, err
	}
	return &config, nil
}