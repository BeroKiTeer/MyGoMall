# 选择基础镜像
FROM golang:1.23 AS builder

ENV GO_ENV=online
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录为
WORKDIR /usr/src/mygomall

# 将工作目录下的 go.mod 和 go.sum 文件复制到容器中
COPY ../../../app/auth/go.mod ../../../app/auth/go.sum ./app/auth/
COPY ../../../common common

# 下载依赖
WORKDIR /usr/src/mygomall/app/auth
RUN go mod download
RUN go mod verify
# 将这个微服务目录的所有文件复制到容器中
COPY ../../../app/auth /usr/src/mygomall/app/auth/

# 构建应用程序
RUN go build -o /opt/mygomall/app/auth/server

FROM busybox

COPY --from=builder /opt/mygomall/app/auth/server /opt/mygomall/app/auth/server

COPY ../../../app/auth/conf /opt/mygomall/app/auth/conf

WORKDIR /opt/mygomall/app/auth
# 启动程序
CMD ["./server"]
