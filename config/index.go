package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Cron string `json:"cron"`
	Db   []struct {
		Host      string   `json:"host"`
		Port      string   `json:"port"`
		Username  string   `json:"username"`
		Password  string   `json:"password"`
		Databases []string `json:"databases"`
	} `json:"db"`
	S3 struct {
		SecretID   string `json:"secretID"`
		SecretKey  string `json:"secretKey"`
		Endpoint   string `json:"endpoint"`
		BucketName string `json:"bucketName"`
		Region     string `json:"region"`
	} `json:"s3"`
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
