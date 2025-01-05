package routes

import (
	"gitee.com/zyw0605688_admin/go_mysqldump/api"
	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) {
	db := router.Group("/db")
	{
		db.POST("/update", api.DBUpdate)
		db.GET("/list", api.DBList)
		db.DELETE("/delete", api.DBDelete)
	}
	s3 := router.Group("/s3")
	{
		s3.POST("/update", api.S3Update)
		s3.GET("/list", api.S3List)
		s3.DELETE("/delete", api.S3Delete)
	}
}
