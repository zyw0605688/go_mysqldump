package api

import (
	"gitee.com/zyw0605688_admin/go_mysqldump/config"
	"github.com/gin-gonic/gin"
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
