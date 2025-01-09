# web前端部分，进入www目录，安装依赖，打包到dist
FROM registry.cn-shanghai.aliyuncs.com/pingda/node:20.18.0 AS webbuilder
WORKDIR /webbuilder
# 将所有文件复制到当前目录
COPY . .
RUN cd /webbuilder/www \
    npm install --registry=http://registry.npmmirror.com \
    npm run build


FROM registry.cn-shanghai.aliyuncs.com/pingda/golang:1.23.1 AS gobuilder
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /gobuilder
# 将上一阶段的所有文件复制到当前目录
COPY --from=webbuilder /webbuilder .
RUN go mod tidy

# 编译成二进制可执行文件 app
RUN go build -o app main.go

###################
# 接下来创建一个小镜像，用来做最终的运行
###################
FROM registry.cn-shanghai.aliyuncs.com/pingda/ubuntu-cgo-cndate:24.04
WORKDIR /
# 对外映射目录，sqlite文件放置目录
RUN mkdir /mysql_backup
# mysqldump文件要拷进来
COPY --from=gobuilder /gobuilder/assets /assets
# 前端静态资源目录
RUN mkdir /assets/WebUI
COPY --from=gobuilder /gobuilder/www/dist /assets/WebUI
# 后端
COPY --from=gobuilder /gobuilder/app /

# 声明服务端口
EXPOSE 3028

# 需要运行的命令
ENTRYPOINT ["/app"]