package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DsnRequest struct {
	Dsn string `json:"dsn"`
}

func GetDbsByDsn(c *gin.Context) {
	var req DsnRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}

	db, err := gorm.Open(mysql.Open(req.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("数据库连接失败：", err)
		return
	}
	fmt.Println("数据库连接成功!___________")
	var tableList []string
	db.Raw("show DATABASES;").Scan(&tableList)
	c.JSON(200, gin.H{
		"code": 0,
		"data": tableList,
		"msg":  "",
	})
}
