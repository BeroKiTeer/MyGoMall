# 选择基础镜像
FROM golang:1.23 AS builder

ENV GO_ENV=online
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录为
WORKDIR /usr/src/mygomall

# 将工作目录下的 go.mod 和 go.sum 文件复制到容器中
COPY ../../../app/user/go.mod ../../../app/user/go.sum ./app/user/
COPY ../../../common common

# 下载依赖
WORKDIR /usr/src/mygomall/app/user
RUN go mod download
RUN go mod verify
# 将这个微服务目录的所有文件复制到容器中
COPY ../../../app/user /usr/src/mygomall/app/user/

# 构建应用程序
RUN go build -o /opt/mygomall/app/user/server

FROM busybox

COPY --from=builder /opt/mygomall/app/user/server /opt/mygomall/app/user/server

COPY ../../../app/user/conf /opt/mygomall/app/user/conf

WORKDIR /opt/mygomall/app/user

# 暴露的端口
# EXPOSE 8883

# 启动程序
CMD ["./server"]
