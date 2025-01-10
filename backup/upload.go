package backup

import (
	"context"
	"fmt"
	"gitee.com/zyw0605688_admin/go_mysqldump/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"gocloud.dev/blob/s3blob"
	"io"
	"os"
	"strings"
)

// 上传文件到S3
func uploadFileToS3(fileUrl string, fileKey string, s3Item config.S3Config) error {
	// 创建 AWS SDK 配置
	s3Config := aws.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(s3Item.AccessKey, s3Item.SecretKey, "")).
		WithEndpoint(s3Item.Endpoint).
		WithRegion(s3Item.Region)

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
	fileContent, err := os.ReadFile(fileUrl)
	if err != nil {
		fmt.Println("读取文件失败", err)
		return err
	}
	fileContentStr := string(fileContent)
	reader := strings.NewReader(fileContentStr)
	w, err := bucket.NewWriter(context.Background(), fileUrl, nil)
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
