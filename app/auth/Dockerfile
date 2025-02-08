# 基础镜像
FROM golang:1.23
# 工作目录
WORKDIR /app

# 复制模块依赖
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制服务端代码
COPY . .

# 构建服务端
RUN go build -o kitex-server .

# 启动服务端
CMD ["./auth"]