package api

import (
	"gitee.com/zyw0605688_admin/go_mysqldump/config"
	"github.com/gin-gonic/gin"
)

func S3Update(c *gin.Context) {
	var req config.S3Config
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

func S3List(c *gin.Context) {
	var res []config.S3Config
	config.GlobalDB.Find(&res)
	c.JSON(200, gin.H{
		"code": 0,
		"data": res,
		"msg":  "",
	})
}

func S3Delete(c *gin.Context) {
	config.GlobalDB.Where("id = ?", c.Query("ID")).Delete(&config.S3Config{})
	c.JSON(200, gin.H{
		"code": 0,
		"data": "",
		"msg":  "",
	})
}
