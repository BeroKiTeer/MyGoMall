# **API 接口文档**

## **1. 文档信息**

| **属性**     | **内容**               |
| ------------ | ---------------------- |
| **项目名称** | 请输入项目名称         |
| **文档版本** | v1.0                   |
| **编写人**   | 请输入姓名             |
| **编写日期** | 请输入日期             |
| **审批人**   | 请输入审批人           |
| **状态**     | 草稿 / 审核中 / 已批准 |

------

## **2. 架构概述**

> 本文档描述了基于 **Protobuf + Kitex + Hertz** 技术栈的 API 设计，适用于 **微服务架构** 下的 **HTTP + RPC** 业务场景。

1. **Nginx 作为主网关**

   - 负责**流量管理、负载均衡、API 认证**

   - 统一**入口网关**，可进行**API 限流**、**日志收集**等操作

   - 可能会结合 Kong / APISIX 进行 API 管理

2. **Hertz 作为内部 HTTP 适配层**

   - 仅负责 **HTTP 到 RPC 的适配**

   - 解析 HTTP 请求，并调用 Kitex 后端 RPC 服务

   - 对于前端来说，提供一致的 **RESTful API**

3. **Kitex 作为微服务 RPC 框架**

   - 处理**核心业务逻辑**

   - 高效**服务间通信**，支持**负载均衡、熔断、超时控制**

   - 内部使用 **Protobuf** 作为数据传输格式

4. **Protobuf 作为序列化协议**

   - 高效、轻量级，适用于 RPC 远程调用

   - 规范 API 定义，支持**跨语言通信**（Go、Java、Python）

### **优势分析**

✅ **高性能 & 低延迟**

- **Nginx 提前过滤流量**，减少不必要的请求进入微服务系统，降低 Hertz 负载。
- **Hertz 仅做协议转换**，避免业务逻辑开销，保持 **轻量级 HTTP 适配层**。
- **Kitex 作为高性能 RPC 框架**，基于 **netpoll** 进行高并发优化，减少 goroutine 切换损耗。
- **Protobuf 提供高效数据传输**，比 JSON 更小更快，降低网络和序列化开销。

✅ **架构清晰，职责分离**

| **组件**     | **职责**                               |
| ------------ | -------------------------------------- |
| **Nginx**    | 入口网关（认证、限流、API 管理）       |
| **Hertz**    | 内部 HTTP 适配层（RESTful API 转 RPC） |
| **Kitex**    | 业务逻辑处理（微服务核心）             |
| **Protobuf** | 轻量级数据序列化                       |

- **Nginx 负责统一流量管理**，可以**屏蔽内部微服务细节**，避免暴露过多内部 API。
- **Hertz 仅作为适配层**，不做业务逻辑，确保**高效转发**，避免 HTTP 层带来的额外开销。
- **Kitex 负责业务逻辑**，采用高性能 RPC 通信，减少 HTTP 请求的**序列化/反序列化损耗**。

✅ **易扩展**

- **前端不需要适应 Kitex**，仍然可以用 **RESTful API** 方式访问（Hertz 提供 RESTful 适配）。
- **内部微服务直接基于 Kitex 进行高效通信**，无需经过 Hertz，提高系统吞吐量。

✅ **更好的安全性**

- 通过 **Nginx 进行 API 认证**，可以限制**访问某些敏感 API**（如 admin 后台）。
- **Hertz 只对 Nginx 开放 API**，不会直接暴露微服务，提高安全性。
- **Kitex 仅用于内部 RPC 通信**，所有RPC通信不会对外暴露，减少直接暴露的攻击面。

------

## **3. 全局约定**

### **3.1 请求格式**

- **协议**：`HTTPS`
- **编码**：`UTF-8`
- **数据格式**：`JSON + Protobuf`
- 请求方法
  - `GET`：查询数据
  - `POST`：提交数据
  - `PUT`：更新数据
  - `DELETE`：删除数据

### **3.2 认证方式**

| **认证方式** | **描述**                        |
| ------------ | ------------------------------- |
| JWT          | `Authorization: Bearer <token>` |
| API Key      | 在 Header 中携带 `X-API-Key`    |

### **3.3 公共响应格式**

所有 API 都遵循以下响应格式：

```json
{
  "code": 200,
  "message": "请求成功",
  "data": { }
}
```

| **字段**  | **类型**       | **描述**                              |
| --------- | -------------- | ------------------------------------- |
| `code`    | `int`          | 状态码，200 表示成功，非 200 表示失败 |
| `message` | `string`       | 返回信息                              |
| `data`    | `object/array` | 返回的数据内容                        |

+ 状态码说明：

  | **状态码** | **描述**            |
  | ---------- | ------------------- |
  | `200`      | 请求成功            |
  | `400`      | 请求参数错误        |
  | `401`      | 未认证或 Token 过期 |
  | `403`      | 没有权限            |
  | `404`      | 资源未找到          |
  | `500`      | 服务器内部错误      |

------

## **4. Protobuf 定义**

```protobuf
syntax = "proto3";

package user;

option go_package = "/user";

// 用户服务
service UserService {
    // 注册用户
    rpc Register(RegisterReq) returns (RegisterResp) {}

    // 用户登录
    rpc Login(LoginReq) returns (LoginResp) {}

    // 获取用户信息
    rpc GetUser(UserReq) returns (UserResp) {}
}

// 注册请求
message RegisterReq {
    string email = 1;
    string password = 2;
    string confirm_password = 3;
}

// 注册响应
message RegisterResp {
    int32 user_id = 1;
}

// 登录请求
message LoginReq {
    string email = 1;
    string password = 2;
}

// 登录响应
message LoginResp {
    int32 user_id = 1;
    string token = 2;
}

// 获取用户信息请求
message UserReq {
    int32 user_id = 1;
}

// 获取用户信息响应
message UserResp {
    int32 user_id = 1;
    string email = 2;
    string username = 3;
    optional string phone = 4;
}
```

------

## **5. API 详细说明**

### **5.1 用户注册**

- **接口描述**：新用户注册接口
- **HTTP 请求方法**：`POST`
- **请求 URL**：`/api/v1/users/register`
- **是否需要认证**：❌

#### **请求参数**

| **字段名**         | **类型** | **必填** | **描述** |
| ------------------ | -------- | -------- | -------- |
| `email`            | `string` | ✅        | 用户邮箱 |
| `password`         | `string` | ✅        | 用户密码 |
| `confirm_password` | `string` | ✅        | 确认密码 |

#### **请求示例**

```json
{
  "email": "example@domain.com",
  "password": "12345678",
  "confirm_password": "12345678"
}
```

#### **响应参数**

| **字段名** | **类型** | **描述** |
| ---------- | -------- | -------- |
| `user_id`  | `int32`  | 用户 ID  |

#### **响应示例**

```json
{
  "user_id": 1001
}
```

------

### **5.2 用户登录**

- **接口描述**：用户登录
- **HTTP 请求方法**：`POST`
- **请求 URL**：`/api/v1/users/login`
- **是否需要认证**：❌

#### **请求参数**

| **字段名** | **类型** | **必填** | **描述** |
| ---------- | -------- | -------- | -------- |
| `email`    | `string` | ✅        | 用户邮箱 |
| `password` | `string` | ✅        | 用户密码 |

#### **请求示例**

```json
{
  "email": "example@domain.com",
  "password": "12345678"
}
```

#### **响应参数**

| **字段名** | **类型** | **描述**       |
| ---------- | -------- | -------------- |
| `user_id`  | `int32`  | 用户 ID        |
| `token`    | `string` | JWT 认证 Token |

#### **响应示例**

```json
{
  "user_id": 1001,
  "token": "eyJhbGciOiJIUzI1NiIsInR5..."
}
```

------

### **5.3 获取用户信息**

- **接口描述**：获取用户信息
- **HTTP 请求方法**：`GET`
- **请求 URL**：`/api/v1/users/profile`
- **是否需要认证**：✅

#### **请求 Header**

```http
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5...
```

#### **响应参数**

| **字段名** | **类型**            | **描述** |
| ---------- | ------------------- | -------- |
| `user_id`  | `int32`             | 用户 ID  |
| `email`    | `string`            | 用户邮箱 |
| `username` | `string`            | 用户名   |
| `phone`    | `string (optional)` | 手机号   |

#### **响应示例**

```json
{
  "user_id": 1001,
  "email": "example@domain.com",
  "username": "example",
  "phone": "13800138000"
}
```

------

## **6. UML 时序图**

```uml
@startuml
participant Client
participant Hertz as "Hertz (HTTP Gateway)"
participant Kitex as "Kitex (User Service)"
participant MySQL as "Database"

Client -> Hertz: POST /api/v1/users/register
Hertz -> Kitex: Register(RegisterReq)
Kitex -> MySQL: INSERT INTO users (email, password)
MySQL -> Kitex: Success
Kitex -> Hertz: RegisterResp(user_id)
Hertz -> Client: HTTP 200 OK
@enduml
```

## **7. 变更记录**

| **版本** | **修改内容** | **修改人** | **修改日期** |
| -------- | ------------ | ---------- | ------------ |
| v1.0     | 初版         | 张三       | 2025-02-04   |
| v1.1     | 增加认证说明 | 李四       | 2025-02-10   |

------

## **8. 参考资料**

> - Kitex 官方文档
> - Hertz 官方文档
> - Protobuf 语法