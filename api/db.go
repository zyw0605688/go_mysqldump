package api

import (
	"fmt"
	"gitee.com/zyw0605688_admin/go_mysqldump/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

func DBUpdate(c *gin.Context) {
	var req config.DBConfig
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	config.GlobalDB.Save(&req)
	c.JSON(200, gin.H{
		"code": 0,
		"data": req,
		"msg":  "",
	})
}

func DBList(c *gin.Context) {
	var res []config.DBConfig
	config.GlobalDB.Find(&res)
	c.JSON(200, gin.H{
		"code": 0,
		"data": res,
		"msg":  "",
	})
}

func DBDelete(c *gin.Context) {
	config.GlobalDB.Where("id = ?", c.Query("ID")).Delete(&config.DBConfig{})
	c.JSON(200, gin.H{
		"code": 0,
		"data": "",
		"msg":  "",
	})
}

func DbBackupList(c *gin.Context) {
	var item config.DBConfig
	config.GlobalDB.Where("id = ?", c.Query("ID")).First(&item)
	var list []map[string]string
	dir, err := ioutil.ReadDir("./mysql_backup")
	if err != nil {
		fmt.Printf("读取目录出错: %v\n", err)
		return
	}
	for _, entry := range dir {
		if !entry.IsDir() {
			fileName := entry.Name()
			if strings.Contains(fileName, item.Host) {
				obj := map[string]string{
					"file": fileName,
				}
				list = append(list, obj)
			}
		}
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": list,
		"msg":  "",
	})
}
