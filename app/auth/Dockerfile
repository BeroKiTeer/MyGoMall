# 选择基础镜像
FROM golang:1.23 AS builder

ENV GO_ENV=online

# 设置工作目录为 /app
WORKDIR /app

# 将当前目录下的 go.mod 和 go.sum 文件复制到容器中
COPY go.mod go.sum ./
# 公共模块
COPY common ./common

# 下载依赖
RUN go mod tidy

# 将当前目录的所有文件复制到容器中
COPY . .

# 编译应用程序
RUN go build -o auth ./main.go

# 运行阶段使用较小的镜像
FROM alpine:latest

# 安装所需的库（如 ca-certificates）
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从 builder 镜像中复制编译好的应用程序
COPY --from=builder /app/auth .

# 如果有 .env 文件，则复制 .env 文件
COPY .env .env

# 暴露容器的端口
EXPOSE 8080

# 启动程序
CMD ["./auth"]
