package api

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gitee.com/zyw0605688_admin/go_mysqldump/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"gocloud.dev/blob"
	"gocloud.dev/blob/s3blob"
	"io/ioutil"
	"net/http"
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
	// 获取本地文件列表
	var localFileList []string
	dir, err := ioutil.ReadDir("./mysql_backup")
	if err != nil {
		fmt.Printf("读取目录出错: %v\n", err)
		return
	}
	for _, entry := range dir {
		if !entry.IsDir() {
			fileName := entry.Name()
			if strings.Contains(fileName, item.Host) {
				localFileList = append(localFileList, fileName)
			}
		}
	}
	// 获取s3文件列表
	var s3FileList []string
	var s3IdList []uint
	json.Unmarshal([]byte(item.S3s.String()), &s3IdList)
	if len(s3IdList) > 0 {
		for _, s := range s3IdList {
			var s3Item config.S3Config
			config.GlobalDB.Where("ID = ?", s).First(&s3Item)
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			s3Config := aws.NewConfig().
				WithCredentials(credentials.NewStaticCredentials(s3Item.AccessKey, s3Item.SecretKey, "")).
				WithEndpoint(s3Item.Endpoint).
				WithRegion(s3Item.Region).WithHTTPClient(&http.Client{Transport: tr})

			// 创建 AWS 会话
			sessions, err := session.NewSession(s3Config)
			if err != nil {
				fmt.Println("创建s3会话失败", err)
				return
			}

			// 打开Bucket
			bucket, err := s3blob.OpenBucket(context.Background(), sessions, s3Item.BucketName, nil)
			if err != nil {
				fmt.Println("打开Bucket失败", err)
				return
			}
			defer bucket.Close()
			iter := bucket.List(&blob.ListOptions{
				Prefix: "mysql_backup/",
			})
			for {
				bucketFileItem, _ := iter.Next(context.Background())
				if bucketFileItem == nil {
					break
				}
				s3FileList = append(s3FileList, bucketFileItem.Key)
			}
		}
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": map[string]interface{}{
			"localFileList": localFileList,
			"s3FileList":    s3FileList,
		},
		"msg": "",
	})
}
