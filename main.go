package main

import (
	_ "embed"
	"fmt"
	"gitee.com/zyw0605688_admin/go_mysqldump/backup"
	"gitee.com/zyw0605688_admin/go_mysqldump/config"
	"gitee.com/zyw0605688_admin/go_mysqldump/routes"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

//go:embed assets/mysqldump
var mysqldumpLinux []byte

func main() {
	// 加载数据库
	config.InitDb()
	// 加载json配置
	conf, err := config.GetConfig()
	if err != nil {
		fmt.Println("读取配置报错", err)
		return
	}
	fmt.Println("配置信息：", conf)
	// 获取mysqldump的绝对路径
	execFilePath, err := getExecFilePath()
	fmt.Println("mysqldump命令工具临时位置：", *execFilePath)
	if err != nil {
		fmt.Println("获取mysqldump命令地址报错", err)
		return
	}
	// 定时备份数据
	c := cron.New()
	_, err = c.AddFunc(conf.Cron, func() {
		backup.Dump(execFilePath, conf)
	})
	if err != nil {
		fmt.Println("添加定时任务失败", err)
		return
	}
	c.Start()

	// 优雅退出程序，清理临时文件
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		// 清理临时目录，也就是mysqldump命令行工具临时文件
		tempDirPath := filepath.Dir(*execFilePath)
		err := os.RemoveAll(tempDirPath)
		if err != nil {
			fmt.Println("删除临时文件出错", err)
			fmt.Println("请自行删除", tempDirPath)
		}
		os.Exit(0)
	}()

	r := gin.Default()
	r.MaxMultipartMemory = 10240000 << 20
	// 允许跨域
	r.Use(cors.Default())
	// 静态资源
	r.Static("/www", "./assets/WebUI")
	routes.InitRouters(r)
	err = r.Run(":3028")
	if err != nil {
		fmt.Println("服务启动失败：", err)
		return
	}
	fmt.Println("服务启动成功,测试地址http://0.0.0.0:3028")
}

func getExecFilePath() (*string, error) {
	var execFilePath string
	var execFileBytes []byte
	tempDir, err := os.MkdirTemp("", "mysqldump-")
	if err != nil {
		fmt.Println("创建临时目录失败", err)
		return nil, err
	}
	execFilePath = filepath.Join(tempDir, "mysqldump")
	execFileBytes = mysqldumpLinux
	err = fileutil.WriteBytesToFile(execFilePath, execFileBytes)
	if err != nil {
		fmt.Println("创建mysqldump文件失败", err)
		return nil, err
	}
	// 设置可执行权限
	err = os.Chmod(execFilePath, 0755)
	if err != nil {
		fmt.Println("设置mysqldump可执行权限失败", err)
		return nil, err
	}
	return &execFilePath, err
}
