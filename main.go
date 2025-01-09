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
	"path/filepath"
)

//go:embed assets/mysqldump
var mysqldumpLinux []byte

func main() {
	// 加载数据库
	config.InitDb()

	// 获取mysqldump执行路径
	myExecFilePath, _ := getExecFilePath()
	backup.MyExecFilePath = myExecFilePath
	// 初始化MyCron实例
	c := cron.New()
	backup.MyCron = c

	// 开始备份任务
	backup.StartAndReload()

	// 优雅退出程序，清理临时文件
	backup.CleanFile()

	// 开启web服务
	r := gin.Default()
	r.MaxMultipartMemory = 10240000 << 20
	// 允许跨域
	r.Use(cors.Default())
	// 静态资源
	r.Static("/www", "./assets/WebUI")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
			"data": "服务已启动，web页面请访问/www",
			"msg":  "",
		})
	})
	routes.InitRouters(r)
	err := r.Run(":3028")
	if err != nil {
		fmt.Println("服务启动失败：", err)
		return
	}
	fmt.Println("服务启动成功,测试地址http://0.0.0.0:3028")
}

func getExecFilePath() (resp string, err error) {
	var execFilePath string
	tempDir, err := os.MkdirTemp("", "mysqldump-")
	if err != nil {
		fmt.Println("创建临时目录失败", err)
		return
	}
	execFilePath = filepath.Join(tempDir, "mysqldump")
	err = fileutil.WriteBytesToFile(execFilePath, mysqldumpLinux)
	if err != nil {
		fmt.Println("创建mysqldump文件失败", err)
		return
	}
	// 设置可执行权限
	err = os.Chmod(execFilePath, 0755)
	if err != nil {
		fmt.Println("设置mysqldump可执行权限失败", err)
		return
	}
	resp = execFilePath
	fmt.Println("mysqldump命令工具临时位置：", resp)
	return
}
