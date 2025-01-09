FROM registry.cn-shanghai.aliyuncs.com/pingda/golang:1.23.1 AS gobuilder
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录：/go
WORKDIR /gobuilder
# 将代码复制到容器中
COPY . .
RUN go mod tidy

# 将我们的代码编译成二进制可执行文件 app
RUN go build -o app main.go

###################
# 接下来创建一个小镜像
###################
FROM registry.cn-shanghai.aliyuncs.com/pingda/ubuntu-cgo-cndate:24.04
# 从builder镜像中把/dist/app 拷贝到当前目录
WORKDIR /
# 对外映射目录，sqlite文件放置目录
RUN mkdir /mysql_backup
# 前端dist静态资源目录
COPY --from=gobuilder /gobuilder/assets /assets
COPY --from=gobuilder /gobuilder/app /

# 声明服务端口
EXPOSE 3028

# 需要运行的命令
ENTRYPOINT ["/app"]