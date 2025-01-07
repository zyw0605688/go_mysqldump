package config

import (
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type DBConfig struct {
	gorm.Model
	Host         string         `gorm:"column:host" json:"host"`
	Port         string         `gorm:"column:port" json:"port"`
	Username     string         `gorm:"column:username" json:"username"`
	Password     string         `gorm:"column:password" json:"password"`
	Cron         string         `gorm:"column:cron" json:"cron"`
	Dbs          datatypes.JSON `gorm:"column:dbs;type:json" json:"dbs"`
	IsLocalStore bool           `gorm:"column:is_local_store;type:is_local_store" json:"is_local_store"`
	S3s          datatypes.JSON `gorm:"column:s3s;type:json" json:"s3s"`
}

type S3Config struct {
	gorm.Model
	Name       string `gorm:"column:name" json:"name"`
	AccessKey  string `gorm:"column:accessKey" json:"accessKey"`
	SecretKey  string `gorm:"column:secretKey" json:"secretKey"`
	Endpoint   string `gorm:"column:endpoint" json:"endpoint"`
	BucketName string `gorm:"column:bucketName" json:"bucketName"`
	Region     string `gorm:"column:region" json:"region"`
}

var GlobalDB *gorm.DB

func InitDb() {
	db, err := gorm.Open(sqlite.Open("./mysql_backup/sqlite.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("数据库连接失败：", err)
		return
	}
	fmt.Println("数据库连接成功!")

	err = db.AutoMigrate(
		&DBConfig{},
		&S3Config{},
	)
	if err != nil {
		fmt.Println("迁移数据表失败：", err)
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	GlobalDB = db

	return
}
