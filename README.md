<div align="center">
 <h1>🛍️ MyGoMall<br/><small>一个生产级教学模板</small></h1>
 <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white"/>
 <img src="https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white"/>
 <img src="https://img.shields.io/badge/cloudwego-%23008ECF.svg?style=for-the-badge&logo=bytedance&logoColor=white"/>
</div>
# 🌟 简介

MyGoMall 是一个基于分布式微服务架构的电商平台，提供用户认证、商品管理、购物车、订单、支付等功能。

## ✨ 核心特性

- 🔐 **认证系统** - 基于JWT的用户注册和登录
- 📦 **商品管理** - 完整的商品目录系统
- 🛒 **购物车** - 强大的购物车功能
- 📋 **订单处理** - 订单管理和追踪
- 💳 **支付集成** - 支付网关集成就绪
- 🏗️ **清晰架构** - 行业标准的项目结构
- 📝 **详细日志** - 全面的日志系统
- ⚙️ **易于配置** - 基于YAML的配置管理

> [!NOTE]  
>
> - 需要 Go >= 1.18
> - 需要 MySQL >= 8.0
> - 推荐 Redis >= 6.0 用于会话管理

## 📚 目录

- [功能概述](#-功能概述)
- [技术栈](#-技术栈)
- [项目结构](#-项目结构)
- [快速开始](#-快速开始)
  - [前置要求](#前置要求)
  - [安装说明](#安装说明)
  - [配置说明](#配置说明)
- [API文档](#-api文档)
- [开发指南](#-开发指南)
- [数据库架构](#-数据库架构)
- [贡献指南](#-贡献指南)
- [许可证](#-许可证)
- [作者](#-作者)

## 🛠️ 技术栈

<div align="center">
  <table>
    <tr>
      <td align="center" width="96">
        <img src="https://cdn.simpleicons.org/go" width="48" height="48" alt="Go" />
        <br>Go
      </td>
      <td align="center" width="96">
        <img src="https://cdn.simpleicons.org/mysql" width="48" height="48" alt="MySQL" />
        <br>MySQL
      </td>
      <td align="center" width="96">
        <img src="https://cdn.simpleicons.org/redis" width="48" height="48" alt="Redis" />
        <br>Redis
      </td>
      <td align="center" width="96">
        <img src="https://cdn.simpleicons.org/jsonwebtokens" width="48" height="48" alt="JWT" />
        <br>JWT
      </td>
      <td align="center" width="96">
        <img src="https://cdn.simpleicons.org/bytedance" width="48" height="48" alt="Kitex" />
        <br>Kitex
      </td>
      <td align="center" width="96">
        <img src="https://cdn.simpleicons.org/go/00ADD8" width="48" height="48" alt="GORM" />
        <br>GORM
      </td>
      <td align="center" width="96">
        <img src="https://cdn.simpleicons.org/bytedance" width="48" height="48" alt="Hertz" />
        <br>Hertz
      </td>
    </tr>
  </table>
</div>
> [!TIP]  
> 我们的技术栈中的每个组件都是基于其可靠性和在生产环境中的广泛采用而选择的。

### 开发中的挑战

1. **认证流程**

   - 安全存储密钥
   - 正确解码Token并认证
   - 处理令牌过期

2. **表单处理**
   - 版本1：使用HTML5验证
   - 版本2：实现受控组件

3. **容器化部署**

   - Consul的服务注册与发现
   - 检查网关与路由的配置
   - Redis的集群部署

4. **下单-扣除库存-支付**

   - 准确的服务功能
   - 操作的先后顺序
   - Saga分布式事务

5. 

   📂 项目结构

```tex
MyGoMall/
├── apis/                  # API层
│   └── ...   				# Hertz 代码
├── app/                  # 各个微服务
│   ├── cart/              # 购物车服务
│   ├── order/              # 订单服务
│   └── ...              # 其他服务
├── common/             # 可复用包
│   ├── clientsuite/     # 客户端组件
│   ├── constant/         # 统一常量
│   ├── kitex_gen/          # Kitex 生成的代码
│   ├── mtl/             # metrics tracer logger
│   ├── rpc/             # RPC 客户端
│   ├── script/         # 执行脚本
│   ├── serversuite/    # 服务端组件
│   └── utils/			# 工具类
├── deploy/             # 容器化部署相关
│   └── ...      
├── docs/				# 说明文件以及相关文档
│   └── ....
├── idl/             # IDL文件(protobuf)
├── .gitignore			# 防止小孩误食
└── go.work				# 工作区文件
```

```

```

## 🚀 快速开始

### 前置要求

> [!IMPORTANT]  
> 在开始之前，确保您已安装以下内容：
> - Go 1.16或更高版本
> - MySQL 8.0或更高版本
> - Git

### 安装说明

1. 克隆仓库：
```bash
git clone https://github.com/BeroKiTeer/MyGoMall.git
cd MyGoMall
```

2. 安装依赖：
```bash
go mod download
或
go mod tidy
```

3. 设置数据库：
```bash
cd docs/database
# 按照提示，在 `MyGoMall` database 中建立数据表与索引
```

4. 配置应用：
```bash
cd app/[服务名]/conf/dev
# 使用您的数据库凭证编辑 各个微服务的配置文件
```

5. 启动服务器：
```bash
go run cmd/server/main.go
```

## 📝 API文档

```bash
cd docs/api
```



## 🗄️ 数据库架构

我们的综合电商数据库包括：

- `user`: 用户账户和认证
- `product`: 商品基本信息
- `category`: 商品类别
- `category_product` : 中间表，在`product`表与`category`表中建立联系
- `order`: 订单处理
- `payment`: 支付记录

## 🤝 贡献指南

我们欢迎贡献！请按照以下步骤操作：

1. Fork 本仓库
2. 创建您的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交您的更改 (`git commit -m '添加一些很棒的特性'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启一个 Pull Request

## 📄 许可证

本项目采用 Apache-2.0 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙋‍♀ 作者


---

<div align="center">
为 Go 学习者用❤️制作
<br/>
⭐ 在 GitHub 上为我们加注星标 | 📖 阅读 Wiki | 🐛 报告问题
</div>