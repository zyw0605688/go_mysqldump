
## 一个定时备份mysql的工具，上传s3

### 使用到的模块
1. mysqldump
2. cron
3. go cdk
4. go embed

### 代码执行流程
1. 读取配置文件config.json（同目录下得有这个文件）
2. 根据配置文件，生成定时任务
3. 任务里调用mysqldump，远程备份数据库。基本命令如下
   ./mysqldump -h 172.16.66.99 -P 3306 -u root -proot --databases ZmosPublicDb2 ZmosPublicDb > backup_file.sql
4. 将备份文件上传到指定s3云存储

### 配置文件
```
{
  "cron": "30 * * * *",
  "db": [
    {
      "Host": "172.16.66.xx",
      "Port": "3306",
      "Username": "root",
      "Password": "root",
      "Databases": ["ZmosPublicDb", "ZmosPublicDb2"]
    },
    {
      "host": "172.16.66.xx",
      "port": "33060",
      "username": "root",
      "password": "root",
      "databases": ["testaaa"]
    }
  ],
  "s3": {
    "secretID": "xxx",
    "secretKey": "xxx",
    "endpoint": "obs.cn-east-3.myhuaweicloud.com",
    "bucketName": "v3-dev",
    "region": "cn-east-3"
  }
}
```
1. cron: 定时任务的cron表达式，五个字段。在线生成[https://cron.ciding.cc/](https://cron.ciding.cc/)
2. db: 数据库配置，支持多个数据库。每个配置域下，host、port、username、password、databases是必须的。大小写，格式需严格符合。
3. s3: 支持华为云obs, 阿里云oss,腾讯云cos,七牛云, 又拍云，百度云，minio,亚马逊s3等任意支持s3协议的云存储。每个配置域下secretID、secretKey、endpoint、bucketName、region是必须的。大小写，格式需严格符合。

### 使用流程
1. 下载release下可执行程序go_mysqldump_linux，config.json文件，放到同一目录下
2. 修改config.json配置文件
3. 运行程序go_mysqldump_linux(会占用终端)，也可以使用nohup go_mysqldump_linux > go_mysqldump.log 2>&1 &
4. 结束进程，使用ps aux | grep 'go_mysqldump_' 查看进程，再使用   kill -9 pid    结束进程

### 二次开发，打包
1. 主程序在main.go中，很简单的几个方法。可参考上面的代码执行流程
2. 打包命令
```
go env -w CGO_ENABLED=0  GOOS=linux/windows  GOARCH=amd64
go build -o go_mysqldump_linux main.go
go build -o "go_mysqldump_windows.exe" main.go
```