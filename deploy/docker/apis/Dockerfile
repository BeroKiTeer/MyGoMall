# 选择基础镜像
FROM golang:1.23 AS builder

ENV GO_ENV=online
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录为
WORKDIR /usr/src/mygomall

# 将工作目录下的 go.mod 和 go.sum 文件复制到容器中
COPY ../../../apis/go.mod apis/go.sum ./apis/
COPY ../../../common common

# 下载依赖
WORKDIR /usr/src/mygomall/apis
RUN go mod download
RUN go mod verify
# 将这个微服务目录的所有文件复制到容器中
COPY ../../../apis /usr/src/mygomall/apis/

# 构建应用程序
RUN go build -o /opt/mygomall/apis/server

FROM busybox

COPY --from=builder /opt/mygomall/apis/server /opt/mygomall/apis/server

COPY ../../../apis/conf /opt/mygomall/apis/conf

WORKDIR /opt/mygomall/apis

# 暴露的端口
EXPOSE 8080

# 启动程序
CMD ["./server"]
