### 一个定时备份mysql的工具，上传s3
1. 内嵌官方mysqldump命令行工具
2. 调用命令行工具，定时执行，上传到s3对象存储

### 项目地址
1. 代码仓库 https://gitee.com/zyw0605688_admin/go_mysqldump
2. dockerhub地址 https://hub.docker.com/r/zyw0605688/go-mysqldump

### 使用方式
1. 启动容器 docker run -itd --name go-mysqldump -v /data/mysql_backup:/mysql_backup -p 8080:3028 --restart always zyw0605688/go-mysqldump:latest
2. 打开页面，访问 http://x.x.x.x:8080/www/  8080是你暴露的端口号（可自行修改），3028是程序监听端口号

### 可视化配置以下信息
1. cron: 定时任务的cron表达式，五个字段。在线生成(https://crontab.run/zh)
2. db: 数据库配置，支持多个数据库。
3. s3: 支持华为云obs, 阿里云oss,腾讯云cos,七牛云, 又拍云，百度云，minio,亚马逊s3等任意支持s3协议的云存储。

### 使用到的模块
1. mysqldump
2. cron
3. go cdk
4. go embed

### 代码执行流程
1. 读取数据库存储的配置信息
2. 根据配置，生成定时任务
3. 任务里调用mysqldump，远程备份数据库。基本命令如下
   ./mysqldump -h 172.16.66.99 -P 3306 -u root -proot --databases ZmosPublicDb2 ZmosPublicDb > 20241009160701.sql
4. 将备份文件压缩上传到指定s3云存储,会存到bucketName/mysql_backup/20241009160701.zip 文件名是时间

### 二次开发
1. 主程序在main.go中，很简单的几个方法。可参考上面的代码执行流程
2. 如有需要，可以下载mysql包，替换mysqldump工具，以更换版本。目前使用的8.4.2 LTS版
3. 打包docker build -t zyw0605688/go-mysqldump:latest .