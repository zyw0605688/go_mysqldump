package backup

import (
	"bytes"
	"fmt"
	"gitee.com/zyw0605688_admin/go_mysqldump/config"
	"github.com/duke-git/lancet/v2/fileutil"
	"os"
	"os/exec"
	"time"
)

func Dump(execFilePath *string, config *config.Config) {
	var backupZipFilePath = time.Now().Format("20060102150405") + ".zip"
	defer os.Remove(backupZipFilePath)
	for i, item := range config.Db {
		cmd := exec.Command(*execFilePath, "-h", item.Host, "-P", item.Port, "-u", item.Username, "-p"+item.Password, "--databases")
		cmd.Args = append(cmd.Args, item.Databases...)
		fmt.Println("即将执行命令：", cmd.String())
		// 创建一个缓冲区来捕获输出
		var out bytes.Buffer
		cmd.Stdout = &out

		// 执行命令
		err := cmd.Run()
		if err != nil {
			fmt.Println("执行命令失败", err)
			continue
		}

		// 将输出写入文件
		now := time.Now().Format("20060102150405")
		backupFilePath := item.Host + "_" + now + ".sql"
		err = fileutil.WriteBytesToFile(backupFilePath, out.Bytes())
		if err != nil {
			fmt.Println("写入文件失败", err)
			continue
		}
		fmt.Println("备份成功:", backupFilePath)
		if i == 0 {
			err = fileutil.Zip(backupFilePath, backupZipFilePath)
		} else {
			err = fileutil.ZipAppendEntry(backupFilePath, backupZipFilePath)
		}
		if err != nil {
			fmt.Println("备份文件写入压缩包失败", err)
			continue
		}
	}
	// 上传文件到s3
	for _, item := range config.S3 {
		err := uploadFileToS3(backupZipFilePath, item)
		if err != nil {
			fmt.Println("上传文件到s3失败", err)
			return
		}
	}
}
