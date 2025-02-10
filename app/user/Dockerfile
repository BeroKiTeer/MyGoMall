# ============================
# 🚀 构建阶段 (编译 Go 代码)
# ============================
FROM golang:1.23 AS builder

# 设置 Go 代理（防止国内访问不了）
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

# 复制 go.mod 和 go.sum（利用 Docker 缓存）
COPY go.mod go.sum ./
RUN go mod tidy

# 复制当前微服务代码
COPY . .

# 进行静态编译，防止运行时报 `no such file or directory`
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/user .

# ============================
# 🏃 运行阶段 (基于 CentOS)
# ============================
FROM centos:latest

# 设置工作目录
WORKDIR /root/

# 复制二进制文件（修正路径）
COPY --from=builder /app/user /user

# 赋予执行权限（确保文件可执行）
RUN chmod +x /user

# 运行 Go 服务
CMD ["/user"]
