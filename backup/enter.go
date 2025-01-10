package backup

import (
	"fmt"
	"gitee.com/zyw0605688_admin/go_mysqldump/config"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var MyCron *cron.Cron
var MyExecFilePath string

func StartAndReload() {
	MyCron.Stop()
	// 获取所有配置
	var list []config.DBConfig
	config.GlobalDB.Find(&list)
	// 循环开始任务
	for _, item := range list {
		if item.IsBackup {
			// 定时备份数据
			_, err := MyCron.AddFunc(item.Cron, func() {
				Dump(MyExecFilePath, item)
			})
			if err != nil {
				fmt.Println("添加定时任务失败", err)
				return
			}
		}
	}
	MyCron.Start()
}

func CleanFile() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		// 清理临时目录，也就是mysqldump命令行工具临时文件
		tempDirPath := filepath.Dir(MyExecFilePath)
		err := os.RemoveAll(tempDirPath)
		if err != nil {
			fmt.Println("删除临时文件出错", err)
			fmt.Println("请自行删除", tempDirPath)
		}
		os.Exit(0)
	}()
}
