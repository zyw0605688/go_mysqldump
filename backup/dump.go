package backup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitee.com/zyw0605688_admin/go_mysqldump/config"
	"github.com/duke-git/lancet/v2/fileutil"
	"os"
	"os/exec"
	"time"
)

func Dump(execFilePath string, item config.DBConfig, i int) {
	var backupZipFilePath = "/mysql_backup/" + time.Now().Format("20060102150405") + ".zip"
	defer os.Remove(backupZipFilePath)
	var dbList []string
	json.Unmarshal([]byte(item.Dbs.String()), &dbList)
	cmd := exec.Command(execFilePath, "-h", item.Host, "-P", item.Port, "-u", item.Username, "-p"+item.Password, "--databases")
	for _, a := range dbList {
		cmd.Args = append(cmd.Args, a)
	}
	fmt.Println("即将执行命令：", cmd.String())
	// 创建一个缓冲区来捕获输出
	var out bytes.Buffer
	cmd.Stdout = &out

	// 执行命令
	err := cmd.Run()
	if err != nil {
		fmt.Println("执行命令失败", err)
	}

	// 将输出写入文件
	now := time.Now().Format("20060102150405")
	backupFilePath := "/mysql_backup/" + item.Host + "_" + now + ".sql"
	err = fileutil.WriteBytesToFile(backupFilePath, out.Bytes())
	if err != nil {
		fmt.Println("写入文件失败", err)
	}
	fmt.Println("备份成功:", backupFilePath)
	if i == 0 {
		err = fileutil.Zip(backupFilePath, backupZipFilePath)
	} else {
		err = fileutil.ZipAppendEntry(backupFilePath, backupZipFilePath)
	}
	if err != nil {
		fmt.Println("备份文件写入压缩包失败", err)
	}
	// 上传文件到s3
	var s3IdList []uint
	json.Unmarshal([]byte(item.S3s.String()), &s3IdList)
	if len(s3IdList) > 0 {
		for _, s := range s3IdList {
			var s3item config.S3Config
			config.GlobalDB.Where("ID = ?", s).First(&s3item)
			err := uploadFileToS3(backupZipFilePath, s3item)
			if err != nil {
				fmt.Println("上传文件到s3失败", err)
				return
			}
		}
	}

}
