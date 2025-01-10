package backup

import (
	"context"
	"crypto/tls"
	"fmt"
	"gitee.com/zyw0605688_admin/go_mysqldump/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/duke-git/lancet/v2/fileutil"
	"gocloud.dev/blob/s3blob"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// 上传文件到S3
func uploadFileToS3(fileUrl string, s3Item config.S3Config) error {
	// 创建 AWS SDK 配置
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
		return err
	}

	// 打开Bucket
	bucket, err := s3blob.OpenBucket(context.Background(), sessions, s3Item.BucketName, nil)
	if err != nil {
		fmt.Println("打开Bucket失败", err)
		return err
	}
	defer bucket.Close()

	// 读取文件，上传
	tempZip := "./temp.zip"
	defer os.Remove(tempZip)
	err = fileutil.Zip(fileUrl, tempZip)
	if err != nil {
		fmt.Println(err)
	}
	fileContent, err := os.ReadFile(tempZip)
	if err != nil {
		fmt.Println("读取文件失败", err)
		return err
	}
	fileContentStr := string(fileContent)
	reader := strings.NewReader(fileContentStr)
	fileKey := "mysql_backup/" + time.Now().Format("20060102150405") + ".zip"
	w, err := bucket.NewWriter(context.Background(), fileKey, nil)
	if err != nil {
		fmt.Println("创建文件写入流失败", err)
		return err
	}
	_, err = io.Copy(w, reader)
	if err != nil {
		fmt.Println("写入文件失败", err)
		return err
	}
	err = w.Close()
	if err != nil {
		fmt.Println("关闭文件写入流失败", err)
		return err
	}
	fmt.Println("文件上传s3成功")
	return nil
}
