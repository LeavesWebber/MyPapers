FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=Linux \
    GOARCH=amd64

# 移动到工作目录: /build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY ./api ./api
COPY ./config ./config
COPY ./core ./core
COPY ./dao ./dao
COPY ./global ./global
COPY ./image ./image
COPY ./initialize ./initialize
COPY ./log ./log
COPY ./logic ./logic
COPY ./middleware ./middleware
COPY ./model ./model
COPY ./public ./public
COPY ./router ./router
COPY ./utils ./utils
COPY ./times.ttf ./times.ttf
COPY ./config.yaml ./config.yaml
COPY ./main.go ./main.go

# 将我们的代码编译成二进制可执行文件 my-papers
RUN go build -o my-papers .

##################
# 接下来创建一个小镜象
##################
FROM scratch

COPY ./config.yaml ./config.yaml
COPY ./times.ttf ./times.ttf
COPY ./image ./image

# 从my-papers镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/my-papers /

# 声明服务端口
EXPOSE 8887

# 需要运行的命令
ENTRYPOINT ["/my-papers","/config.yaml"]