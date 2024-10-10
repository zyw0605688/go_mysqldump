package main

import (
	_ "embed"
	"fmt"
	"gitee.com/zyw0605688_admin/go_mysqldump/backup"
	"gitee.com/zyw0605688_admin/go_mysqldump/config"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/robfig/cron/v3"
	"os"
	"path/filepath"
	"runtime"
)

//go:embed assets/mysqldump
var mysqldumpLinux []byte

//go:embed assets/mysqldump.exe
var mysqldumpWindows []byte

func main() {
	// 加载json配置
	conf, err := config.GetConfig()
	if err != nil {
		fmt.Println("读取配置报错", err)
		return
	}
	fmt.Println("配置信息：", conf)
	// 根据系统，获取mysqldump的绝对路径
	execFilePath, err := getExecFilePath()
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
	select {}
}

func getExecFilePath() (*string, error) {
	var execFilePath string
	var execFileBytes []byte
	root, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	tempDir, err := os.MkdirTemp(root, "mysqldump-")
	if err != nil {
		fmt.Println("创建临时目录失败", err)
		return nil, err
	}
	if runtime.GOOS == "windows" {
		execFilePath = filepath.Join(tempDir, "mysqldump.exe")
		execFileBytes = mysqldumpWindows
	} else {
		execFilePath = filepath.Join(tempDir, "mysqldump")
		execFileBytes = mysqldumpLinux
	}
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
