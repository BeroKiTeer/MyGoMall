<div align="center">
 <h1>🛍️ MyGoMall<br/><small>一个生产级教学模板</small></h1>
 <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white"/>
 <img src="https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white"/>
 <img src="https://img.shields.io/badge/cloudwego-%23008ECF.svg?style=for-the-badge&logo=bytedance&logoColor=white"/>
</div>


> [!IMPORTANT]  
> 这是一个用于教学目的的模板项目。虽然它展示了生产级的实践，但在部署到生产环境之前，请务必全面审查并加强安全措施。

# 🌟 简介

这是一个使用Go构建的全面的生产级电商后端模板，专门为Go初学者设计。该项目展示了使用现代工具和框架进行Go Web开发的行业标准实践。

## ✨ 核心特性

- 🔐 **认证系统** - 基于JWT的用户注册和登录
- 📦 **商品管理** - 完整的商品目录系统
- 🛒 **购物车** - 强大的购物车功能
- 📋 **订单处理** - 订单管理和追踪
- 💳 **支付集成** - 支付网关集成就绪
- 🏗️ **清晰架构** - 行业标准的项目结构
- 📝 **详细日志** - 全面的日志系统
- ⚙️ **易于配置** - 基于YAML的配置管理
- 🔄 **数据库迁移** - 结构化的数据库架构管理

> [!NOTE]  
> - 需要 Go >= 1.16
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

### 开发提示

#### 常见挑战
1. **CORS问题**
   - 确保正确的CORS中间件配置
   - 检查浏览器开发工具中的请求头
   - 验证API端点

2. **认证流程**
   - 安全存储JWT令牌
   - 处理令牌过期
   - 实现正确的登出

3. **表单处理**
   - 版本1：使用HTML5验证
   - 版本2：实现受控组件

#### 最佳实践
1. **错误处理**
```javascript
// 版本1
fetch('/api/v1/login', {
  // ... fetch配置
}).catch(error => {
  document.getElementById('error').textContent = error.message;
});

// 版本2
try {
  await loginService.login(credentials);
} catch (error) {
  setError(error.response?.data?.message || '登录失败');
}
```

2. **API集成**
```javascript
// 版本1
const response = await fetch('/api/v1/register', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify(formData)
});

// 版本2
const authService = {
  register: async (userData) => {
    const response = await http.post('/api/v1/register', userData);
    return response.data;
  }
};
```

### 学习路径建议

1. 从版本1开始学习：
   - 基础HTML结构
   - 表单处理
   - API集成
   - 简单状态管理

2. 进阶到版本2学习：
   - React组件
   - Hooks和状态管理
   - 现代构建工具
   - 高级路由

3. 对比两种实现以理解：
   - 代码组织
   - 状态管理方法
   - API集成模式
   - 构建和部署流程

## 📂 项目结构

```
douyin-mall-go-template/
├── api/                  # API层
│   └── v1/              # API版本1处理器
├── cmd/                  # 应用程序入口点
│   └── server/          # 主服务器应用
├── configs/             # 配置文件
├── internal/            # 内部包
│   ├── dao/            # 数据访问对象
│   ├── middleware/     # HTTP中间件
│   ├── model/          # 数据模型和DTO
│   ├── routes/         # 路由定义
│   └── service/        # 业务逻辑层
├── pkg/                 # 可复用包
│   ├── db/             # 数据库工具
│   ├── logger/         # 日志工具
│   └── utils/          # 通用工具
├── frontend/
│   ├── src/
│   │   ├── components/    # 可复用React组件
│   │   ├── pages/         # 页面组件
│   │   ├── services/      # API服务
│   │   └── utils/         # 工具函数
│   ├── package.json
│   └── vite.config.js
└── public/             # 静态资源
```

## 🚀 快速开始

### 前置要求

> [!IMPORTANT]  
> 在开始之前，确保您已安装以下内容：
> - Go 1.16或更高版本
> - MySQL 8.0或更高版本
> - Git
> - Make（可选，用于使用Makefile命令）

### 安装说明

1. 克隆仓库：
```bash
git clone https://github.com/ChanMeng666/douyin-mall-go-template.git
cd douyin-mall-go-template
```

2. 安装依赖：
```bash
go mod download
或
go mod tidy
```

3. 设置数据库：
```bash
mysql -u root -p < docs/database/douyin_mall_go_template_structure_only.sql
```

4. 配置应用：
```bash
cp configs/config.yaml.example configs/config.yaml
# 使用您的数据库凭证编辑configs/config.yaml
```

5. 启动服务器：
```bash
go run cmd/server/main.go
```

## 📝 API文档



## 📖 开发指南

### 项目组件

> [!NOTE]  
> 每个组件都设计为模块化，并遵循SOLID原则：

- **api/v1/**: HTTP请求处理器
  - `health.go`: 健康检查端点
  - `user.go`: 用户相关端点

- **internal/middleware/**: 自定义中间件
  - `auth.go`: JWT认证
  - `cors.go`: CORS处理
  - `logger.go`: 请求日志

- **internal/model/**: 数据模型
  - `user.go`: 用户实体
  - `dto/`: 数据传输对象

- **internal/service/**: 业务逻辑
  - `user_service.go`: 用户相关操作
  - `product_service.go`: 商品相关操作
  - `order_service.go`: 订单处理逻辑

### 添加新功能

> [!TIP]  
> 按照以下步骤向项目添加新功能：

1. 在 `internal/routes/routes.go` 中定义路由
2. 在 `api/v1/` 中创建处理器
3. 在 `internal/service/` 中实现服务逻辑
4. 在 `internal/model/` 中定义模型
5. 在 `internal/dao/` 中添加数据访问层

## 🗄️ 数据库架构

我们的综合电商数据库包括：

- `users`: 用户账户和认证
- `products`: 商品目录管理
- `categories`: 商品分类
- `orders`: 订单处理
- `order_items`: 订单详情
- `shopping_cart_items`: 购物车管理
- `payment_records`: 支付记录
- `product_reviews`: 用户评价和评分

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