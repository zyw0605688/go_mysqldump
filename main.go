package main

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/robfig/cron/v3"
	"gocloud.dev/blob/s3blob"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
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

//go:embed mysqldump
var mysqldumpLinux []byte

//go:embed mysqldump.exe
var mysqldumpWindows []byte

func main() {
	// 加载读取json配置
	config := getConfig()
	fmt.Println("配置列表：", config)
	// 根据系统，获取mysqldump的绝对路径
	execFilePath := getExecFilePath()
	// 定时备份数据
	c := cron.New()
	_, err := c.AddFunc(config.Cron, func() {
		backup(execFilePath, config)
	})
	if err != nil {
		fmt.Println("添加定时任务失败", err)
		return
	}
	c.Start()
	select {}
}

func getConfig() *Config {
	configFile, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println("读取配置失败，请检查配置文件是否存在", err)
		return nil
	}
	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("反序列化配置失败，请检查配置文件，参考README配置说明", err)
		return nil
	}
	return &config
}

func getExecFilePath() *string {
	var execFilePath string
	root, _ := os.Getwd()
	tempDir, _ := os.MkdirTemp(root, "mysqldump-")
	if runtime.GOOS == "windows" {
		execFilePath = filepath.Join(tempDir, "mysqldump.exe")
		file, _ := os.Create(execFilePath)
		defer file.Close()
		io.Copy(file, bytes.NewReader(mysqldumpWindows))
	} else {
		execFilePath = filepath.Join(tempDir, "mysqldump")
		file, _ := os.Create(execFilePath)
		defer file.Close()
		io.Copy(file, bytes.NewReader(mysqldumpLinux))
	}
	// 设置可执行权限
	_ = os.Chmod(execFilePath, 0755)
	return &execFilePath
}

func backup(execFilePath *string, config *Config) {
	for _, item := range config.Db {
		cmd := exec.Command(*execFilePath, "-h", item.Host, "-P", item.Port, "-u", item.Username, "-p"+item.Password, "--databases")
		cmd.Args = append(cmd.Args, item.Databases...)
		fmt.Println("执行命令：", cmd.String())
		// 创建一个缓冲区来捕获输出
		var out bytes.Buffer
		cmd.Stdout = &out

		// 执行命令
		err := cmd.Run()
		if err != nil {
			fmt.Println("执行命令失败", err)
			return
		}

		// 将输出写入文件
		now := time.Now().Format("20060102150405")
		backupFile := item.Host + "_" + now + ".sql"
		err = os.WriteFile(backupFile, out.Bytes(), 0644)
		if err != nil {
			fmt.Println("写入文件失败", err)
			return
		}
		fmt.Println("备份成功:", backupFile)
		uploadFileToS3(backupFile, config)
	}
}

// 删除文件到S3
func uploadFileToS3(fileUrl string, config *Config) {
	// 创建 AWS SDK 配置
	s3Config := aws.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(config.S3.SecretID, config.S3.SecretKey, "")).
		WithEndpoint(config.S3.Endpoint).
		WithRegion(config.S3.Region)

	// 创建 AWS 会话
	sessions, err := session.NewSession(s3Config)
	if err != nil {
		fmt.Println("创建s3会话失败", err)
	}

	// 打开Bucket
	bucket, err := s3blob.OpenBucket(context.Background(), sessions, config.S3.BucketName, nil)
	if err != nil {
		fmt.Println("打开Bucket失败", err)
	}
	defer bucket.Close()

	// 读取文件，上传
	fileContent, err := os.ReadFile(fileUrl)
	if err != nil {
		fmt.Println("读取文件失败", err)
	}
	fileContentStr := string(fileContent)
	reader := strings.NewReader(fileContentStr)
	w, err := bucket.NewWriter(context.Background(), "mysql_backup/"+fileUrl, nil)
	if err != nil {
		fmt.Println("创建文件写入流失败", err)
	}
	_, err = io.Copy(w, reader)
	if err != nil {
		fmt.Println("写入文件失败", err)
	}
	err = w.Close()
	if err != nil {
		fmt.Println("关闭文件写入流失败", err)
	}
	fmt.Println("文件上传s3成功")
}
