# 选择基础镜像
FROM golang:1.23 AS builder

ENV GO_ENV=online
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录为
WORKDIR /usr/src/mygomall

# 将工作目录下的 go.mod 和 go.sum 文件复制到容器中
COPY ../../../app/stock/go.mod ../../../app/stock/go.sum ./app/stock/
COPY ../../../common common

# 下载依赖
WORKDIR /usr/src/mygomall/app/stock
RUN go mod download
RUN go mod verify
RUN go mod tidy

# 将这个微服务目录的所有文件复制到容器中
COPY ../../../app/stock /usr/src/mygomall/app/stock/

# 构建应用程序
RUN go build -o /opt/mygomall/app/stock/server

FROM busybox

COPY --from=builder /opt/mygomall/app/stock/server /opt/mygomall/app/stock/server

COPY ../../../app/stock/conf /opt/mygomall/app/stock/conf

WORKDIR /opt/mygomall/app/stock

# 暴露的端口
# EXPOSE 8883

# 启动程序
CMD ["./server"]
