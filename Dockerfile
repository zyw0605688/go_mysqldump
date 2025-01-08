# web前端部分，进入www目录，安装依赖，打包到dist
FROM 172.16.66.13:88/library/node:18 AS webbuilder
# 移动到工作目录：/home/www
WORKDIR /web
# 将前端代码复制到容器/web
COPY www /web
RUN cd www
# 安装依赖,打包
RUN npm install --registry=https://mirrors.huaweicloud.com/repository/npm/
RUN npm run build


# golang后端部分，编译
FROM hub.atomgit.com/amd64/golang:1.21.1-alpine AS gobuilder
RUN apk add --no-cache gcc musl-dev
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录：/go
WORKDIR /go
# 将代码复制到容器中
COPY . .
RUN go mod tidy

# 将我们的代码编译成二进制可执行文件 app
RUN go build -o app main.go

###################
# 接下来创建一个小镜像
###################
FROM hub.atomgit.com/amd64/alpine:latest

# 设置时区
RUN apk add --no-cache tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 从builder镜像中把/dist/app 拷贝到当前目录
WORKDIR /
# 对外映射目录，sqlite文件放置目录
RUN mkdir /mysql_backup
# 前端dist静态资源目录
RUN mkdir -p /assets/WebUI
COPY --from=webbuilder /web/www/dist /assets/WebUI
# 后端程序文件目录
COPY --from=gobuilder /go/assets/mysqldump /assets
COPY --from=gobuilder /go/app /

# 声明服务端口
EXPOSE 3028

# 需要运行的命令
ENTRYPOINT ["/app"]