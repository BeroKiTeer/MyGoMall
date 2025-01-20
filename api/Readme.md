# -- TODO API 网关 + 服务注册与发现（Consul）

### **1. 设计架构**

- **前端请求**：通过 Hertz API 网关发送 HTTP 请求。
- Hertz 做网关层：
  - 使用 Consul 进行服务发现。
  - 将 HTTP 请求转发到后端 Kitex 服务。
- **后端服务**：Kitex 服务通过 Consul 注册自身。

### 2.-- TODO 依赖注入

在项目中我们将使用类似于SpringIoC**依赖注入**的方式管理 Kitex 客户端
客户端路径：`MyGoMall\common\rpc_client`

### 3.期望的目录结构

```txt
/MyGoMall
│
├── /api		                     # Hertz API 网关代码
│   ├── /conf                        # API 网关配置文件（如 Consul 地址、日志等）
│   │   └── config.yaml
│   ├── /di                      	 # DI 容器
│   │   └── xxx_container.go
│   ├── /routes                      # Hertz API 路由定义
│   │   └── xxx.go                   # 相关路由
│   ├── /services                    # Consul 服务发现相关逻辑
│   │   └── consul_resolver.go       # Consul 服务发现封装
│   ├── main.go                      # API 网关启动文件
│   ├── go.mod                       # Go 模块定义
│   └── go.sum                       # Go 依赖锁定文件
```



### 4. 扩展内容

1. 负载均衡
2. 集成日志(Prometheus)
3. 容器化部署