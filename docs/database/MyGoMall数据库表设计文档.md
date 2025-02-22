## **1. 项目背景**🛍️

MyGoMall 是一个电商平台系统，旨在提供完整的在线购物体验。为了支持平台的功能需求，系统需要处理用户账户、商品、订单、支付、购物车、商品分类、地址和支付记录等各类信息。这个数据库设计用于存储与管理这些信息，确保各个模块之间的顺畅交互，支持高并发、大数据量的访问需求。

该数据库包含了多个表，每个表用于存储特定的业务数据，例如用户信息、商品信息、购物车信息、订单详情等。通过规范化的设计，系统能够支持商品浏览、购物、支付以及订单管理等核心电商功能。同时，通过外键约束保证数据的完整性和一致性，优化了平台的性能，并为后续的业务扩展提供了良好的基础。

## **2. 设计原则🛠️**

+ **数据一致性🔄**：整个数据库系统采用外键约束来保证不同表之间的数据一致性和完整性。每个表中的外键字段都与其它表的主键字段相关联，避免了数据孤立和冗余问题。
+ **高效性⚡**：为了提高查询和插入操作的效率，数据库设计中加入了合适的索引策略，尤其是在频繁查询的字段上（如 `user_id`、`product_id`、`order_id` 等），以加速数据访问，支持高并发。
+ **扩展性🚀**：数据库设计遵循了良好的扩展性原则，能够在后续业务扩展时灵活地增加新的表和字段，以适应不同的需求变化。
+ **规范化📊**：所有表结构遵循数据库规范化设计原则，避免了数据冗余、重复存储，减少了存储空间的浪费，并提高了数据的可维护性。
+ **安全性🔐**：对于用户敏感信息（如密码、支付信息等），设计了加密存储方案，确保用户隐私数据的安全。同时，外部接口和数据访问层需要进行严格的安全控制，以防止数据泄露。
+ **事务管理📝**：在涉及订单处理、支付等关键操作时，保证数据的一致性和完整性，确保事务能够成功执行或回滚，防止因系统故障或网络问题造成数据不一致。

## **3. 数据库表结构**🗃️

### 3.1 `user`👤

- **用途**：存储用户的基本信息

- **引擎**：`InnoDB` 

- **字符集**：`utf8mb4`

  | 字段名          | 数据类型       | 是否主键 | 是否允许为空 | 默认值                                                  | 备注                                 |
  | --------------- | -------------- | -------- | ------------ | ------------------------------------------------------- | ------------------------------------ |
  | `id`            | `BIGINT`       | ✅        | ❌            | `AUTO_INCREMENT`                                        | 主键，自增                           |
  | `email`         | `VARCHAR(255)` | ❌        | ✅            | `NULL`                                                  | 用户邮箱（可选）                     |
  | `username`      | `VARCHAR(100)` | ❌        | ❌            | `NULL`                                                  | 用户名                               |
  | `password_hash` | `VARCHAR(255)` | ❌        | ❌            | `NULL`                                                  | 加密存储的用户密码                   |
  | `phone_number`  | `VARCHAR(20)`  | ❌        | ❌            | `NULL`                                                  | 手机号码，唯一索引                   |
  | `address_id`    | `BIGINT`       | ❌        | ✅            | `NULL`                                                  | 用户地址id（可选）                   |
  | `role`          | `TINYINT`      | ❌        | ❌            | `0`                                                     | 用户角色（0-普通用户, 1-管理员）     |
  | `status`        | `TINYINT`      | ❌        | ❌            | `0`                                                     | 账户状态（0-正常, 1-禁用, 2-待审核） |
  | `created_at`    | `DATETIME`     | ❌        | ✅            | `DEFAULT CURRENT_TIMESTAMP`                             | 账户创建时间                         |
  | `updated_at`    | `DATETIME`     | ❌        | ✅            | `DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP` | 账户更新时间                         |

+ 索引策略

  | **索引名称** | **索引字段**      | **索引类型** | **是否唯一** | **说明**             |
  | ------------ | ----------------- | ------------ | :----------: | -------------------- |
  | `pk_id`      | `id`              | `PRIMARY`    |      ✅       | 主键索引             |
  | `idx_email`  | `order_no`        | `Regular`    |      ❌       | 订单编号唯一索引     |
  | `idx_email`  | `user_id, status` | `UNIQUE`     |      ✅       | 用户订单状态查询加速 |

### 3.2  `cart`🛒

- **用途**：存储购物车的基本信息

- **引擎**：`InnoDB` 

- **字符集**：`utf8mb4`

- **主键策略**：AUTO


| 字段名       | 数据类型   | 是否主键 | 是否允许为空 | 默认值                                                  | 备注                   |
| ------------ | ---------- | -------- | ------------ | ------------------------------------------------------- | ---------------------- |
| `id`         | `BIGINT`   | ✅        | ❌            | `AUTO_INCREMENT`                                        | 主键，用于区分每条数据 |
| `user_id`    | `BIGINT`   | ❌        | ❌            | `NULL`                                                  | 用户ID                 |
| `product_id` | `BIGINT`   | ❌        | ❌            | `NULL`                                                  | 商品ID                 |
| `quantity`   | `BIGINT`   | ❌        | ❌            | `0`                                                     | 商品数量               |
| `created_at` | `DATETIME` | ❌        | ✅            | `DEFAULT CURRENT_TIMESTAMP`                             | 商品创建时间           |
| `updated_at` | `DATETIME` | ❌        | ✅            | `DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP` | 商品信息修改时间       |

+ 索引策略

  | 索引名称           | 索引字段                | 索引类型  | 是否唯一 | 说明                                     |
  | :----------------- | ----------------------- | --------- | :------: | ---------------------------------------- |
  | `pk_id`            | `id`                    | `PRIMARY` |    ✅     | 主键索引                                 |
  | `idx_user_product` | `user_id`, `product_id` | `Regular` |    ❌     | 复合索引，用于加速用户购物车中商品的查询 |

### 3.3 `product`🛍️

- **用途**：存储商品的基本信息

- **引擎**：`InnoDB`

- **字符集**：`utf8mb4`

- **主键策略**：自增主键

| 字段名           | 数据类型   | 是否主键 | 是否允许为空 | 默认值                                                  | 备注                             |
| ---------------- | ---------- | -------- | ------------ | ------------------------------------------------------- | -------------------------------- |
| `id`             | `BIGINT`   | ✅        | ❌            | `AUTO_INCREMENT`                                        | 主键，自增                       |
| `name`           | `LONGTEXT` | ❌        | ❌            | `NULL`                                                  | 商品名称                         |
| `description`    | `LONGTEXT` | ❌        | ✅            | `NULL`                                                  | 商品描述                         |
| `price`          | `BIGINT`   | ❌        | ❌            | `NULL`                                                  | 商品价格                         |
| `original_price` | `BIGINT`   | ❌        | ✅            | `NULL`                                                  | 商品原价                         |
| `images`         | `LONGTEXT` | ❌        | ✅            | `NULL`                                                  | 商品图片                         |
| `status`         | `INT`      | ❌        | ❌            | `1`                                                     | 商品状态，1表示启用，0表示未启用 |
| `created_at`     | `DATETIME` | ❌        | ✅            | `DEFAULT CURRENT_TIMESTAMP`                             | 商品创建时间                     |
| `updated_at`     | `DATETIME` | ❌        | ✅            | `DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP` | 商品信息修改时间                 |
| `deleted_at`     | `DATETIME` | ❌        | ✅            | `NULL`                                                  | 商品删除时间                     |
| `stock`          | `INT`      | ❌        | ❌            | `0`                                                     | 商品库存数量                     |

+ 索引策略

  | 索引名称     | 索引字段 | 索引类型  | 是否唯一 | 说明                             |
  | ------------ | -------- | --------- | :------: | -------------------------------- |
  | `pk_id`      | `id`     | `PRIMARY` |    ✅     | 主键索引，确保每个商品ID唯一     |
  | `idx_name`   | `name`   | `Regular` |    ❌     | 商品名称索引，便于按名称查询商品 |
  | `idx_status` | `status` | `Regular` |    ❌     | 商品状态索引，便于查询商品状态   |
  | `idx_stock`  | `stock`  | `Regular` |    ❌     | 商品库存索引，便于查询库存数量   |

### 3.4 `category`📂

- **用途**：存储商品类别的基本信息

- **引擎**：`InnoDB`

- **字符集**：`utf8mb4`

- **主键策略**：自增主键

  | 字段名       | 数据类型      | 是否主键 | 是否允许为空 | 默认值                                                  | 备注                                 |
  | ------------ | ------------- | -------- | ------------ | ------------------------------------------------------- | ------------------------------------ |
  | `id`         | `BIGINT`      | ✅        | ❌            | `AUTO_INCREMENT`                                        | 主键，自增                           |
  | `name`       | `VARCHAT(50)` | ❌        | ❌            | `NULL`                                                  | 商品类别名称                         |
  | `status`     | `TINYINT`     | ❌        | ❌            | `1`                                                     | 商品类别状态，1表示启用，0表示未启用 |
  | `created_at` | `DATETIME`    | ❌        | ✅            | `DEFAULT CURRENT_TIMESTAMP`                             | 商品类别创建时间                     |
  | `updated_at` | `DATETIME`    | ❌        | ✅            | `DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP` | 商品类别更新时间                     |
  | `deleted_at` | `DATETIME`    | ❌        | ✅            | `NULL`                                                  | 商品类别删除时间                     |

+ 索引策略

  | 索引名称     | 索引字段 | 索引类型  | 是否唯一 | 说明             |
  | ------------ | -------- | --------- | -------- | ---------------- |
  | `pk_id`      | `id`     | `PRIMARY` | ✅        | 主键索引         |
  | `idx_name`   | `name`   | `Regular` | ❌        | 商品类别名称索引 |
  | `idx_status` | `status` | `Regular` | ❌        | 商品类别状态索引 |

### 3.5 `category_product`🔗

- **用途**：建立 categories（商品类别）和 products（商品）之间的多对多关系

- **引擎**：`InnoDB`

- **字符集**：`utf8mb4`

- **主键策略**：自增主键

  | 字段名        | 数据类型 | 是否主键 | 是否允许为空 | 默认值           | 备注                                 |
  | ------------- | -------- | -------- | ------------ | ---------------- | ------------------------------------ |
  | `id`          | `BIGINT` | ✅        | ❌            | `AUTO_INCREMENT` | 主键，自增                           |
  | `category_id` | `BIGINT` | ❌        | ❌            | `NULL`           | 引用categories表的id字段，表示类别ID |
  | `product_id`  | `BIGINT` | ❌        | ❌            | `NULL`           | 引用products表的id字段，表示商品ID   |

+ 索引策略

  | 索引名称       | 索引字段      | 索引类型  | 是否唯一 | 说明                             |
  | -------------- | ------------- | --------- | -------- | -------------------------------- |
  | `pk_id`        | `id`          | `PRIMARY` | ✅        | 主键索引                         |
  | `idx_category` | `category_id` | `Regular` | ❌        | 类别ID索引，便于根据类别查询商品 |
  | `idx_product`  | `product_id`  | `Regular` | ❌        | 商品ID索引，便于根据商品查询类别 |

### 3.6 `order`📝

- **用途**：订单表
- **引擎**：`InnoDB`
- **字符集**：`utf8mb4`
- **主键策略**：自增主键

| 字段名            | 数据类型       | 是否主键 | 是否允许为空 | 默认值                                                  | 备注                                                         |
| ----------------- | -------------- | -------- | ------------ | ------------------------------------------------------- | ------------------------------------------------------------ |
| `id`              | `CHAR(36)`     | ✅        | ❌            | `UUID()`                                                | 订单ID，UUID                                                 |
| `user_id`         | `BIGINT`       | ❌        | ❌            | `NULL`                                                  | 外键，关联`users`表的`id`字段，表示购买用户ID                |
| `total_price`     | `BIGINT`       | ❌        | ❌            | `NULL`                                                  | 订单总金额                                                   |
| `discount_price`  | `BIGINT`       | ❌        | ✅            | `0`                                                     | 优惠金额                                                     |
| `actual_price`    | `BIGINT`       | ❌        | ❌            | `NULL`                                                  | 实际支付金额                                                 |
| `order_status`    | `TINYINT`      | ❌        | ✅            | `0`                                                     | 订单状态（0-待支付, 1-已支付, 2-已发货, 3-已完成, 4-已取消） |
| `payment_status`  | `TINYINT`      | ❌        | ✅            | `0`                                                     | 支付状态（0-未支付, 1-已支付, 2-支付失败, 3-退款中, 4-已退款） |
| `payment_method`  | `VARCHAR(20)`  | ❌        | ✅            | `NULL`                                                  | 支付方式（如微信、支付宝等）                                 |
| `address_id`      | `BIGINT`       | ❌        | ❌            | `NULL`                                                  | 收货地址                                                     |
| `recipient_name`  | `VARCHAR(255)` | ❌        | ❌            | `NULL`                                                  | 收件人姓名                                                   |
| `phone_number`    | `VARCHAR(20)`  | ❌        | ❌            | `NULL`                                                  | 收件人电话号码                                               |
| `shipping_status` | `TINYINT`      | ❌        | ✅            | `0`                                                     | 物流状态（0-未发货, 1-已发货, 2-已签收）                     |
| `paid_at`         | `DATETIME`     | ❌        | ✅            | `NULL`                                                  | 订单支付时间                                                 |
| `shipped_at`      | `DATETIME`     | ❌        | ✅            | `NULL`                                                  | 发货时间                                                     |
| `completed_at`    | `DATETIME`     | ❌        | ✅            | `NULL`                                                  | 订单完成时间                                                 |
| `canceled_at`     | `DATETIME`     | ❌        | ✅            | `NULL`                                                  | 订单取消时间                                                 |
| `remark`          | `TEXT`         | ❌        | ✅            | `NULL`                                                  | 订单备注                                                     |
| `created_at`      | `DATETIME`     | ❌        | ✅            | `DEFAULT CURRENT_TIMESTAMP`                             | 订单创建时间                                                 |
| `updated_at`      | `DATETIME`     | ❌        | ✅            | `DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP` | 订单更新时间                                                 |

+ 索引策略

  | 索引名称       | 索引字段          | 索引类型  | 是否唯一 | 说明                                   |
  | -------------- | ----------------- | --------- | -------- | -------------------------------------- |
  | `pk_id`        | `id`              | `PRIMARY` | ✅        | 主键索引                               |
  | `idx_user_id`  | `user_id`         | `Regular` | ❌        | 用户ID索引，便于查找某个用户的所有订单 |
  | `idx_status`   | `order_status`    | `Regular` | ❌        | 订单状态索引，便于按状态查询订单       |
  | `idx_shipping` | `shipping_status` | `Regular` | ❌        | 物流状态索引，便于按物流状态查询订单   |
  | `idx_payment`  | `payment_status`  | `Regular` | ❌        | 支付状态索引，便于按支付状态查询订单   |

### 3.7 `payment`💳

- **用途**：支付表
- **引擎**：`InnoDB`
- **字符集**：`utf8mb4`

| 字段名           | 数据类型      | 是否主键 | 是否允许为空 | 默认值                                                  | 备注                                       |
| ---------------- | ------------- | -------- | ------------ | ------------------------------------------------------- | ------------------------------------------ |
| `id`             | `BIGINT`      | ✅        | ❌            | `AUTO_INCREMENT`                                        | 支付记录ID，自增主键                       |
| `user_id`        | `BIGINT`      | ❌        | ❌            | `NULL`                                                  | 外键，关联`users`表的`id`字段，表示用户ID  |
| `order_id`       | `CHAR(36)`    | ❌        | ❌            | `NULL`                                                  | 外键，关联`orders`表的`id`字段，表示订单ID |
| `transaction_id` | `VARCHAR(36)` | ❌        | ❌            | `NULL`                                                  | 交易ID（UUID）                             |
| `amount`         | `BIGINT`      | ❌        | ❌            | `NULL`                                                  | 交易金额                                   |
| `pay_at`         | `DATETIME`    | ❌        | ❌            | `DEFAULT CURRENT_TIMESTAMP`                             | 交易时间                                   |
| `created_at`     | `DATETIME`    | ❌        | ❌            | `DEFAULT CURRENT_TIMESTAMP`                             | 创建时间                                   |
| `updated_at`     | `DATETIME`    | ❌        | ❌            | `DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP` | 更新时间                                   |
| `deleted_at`     | `DATETIME`    | ❌        | ✅            | `NULL`                                                  | 删除时间                                   |

+ 索引策略

  | 索引名称             | 索引字段         | 索引类型  | 是否唯一 | 说明                                   |
  | -------------------- | ---------------- | --------- | -------- | -------------------------------------- |
  | `pk_id`              | `id`             | `PRIMARY` | ✅        | 主键索引，确保支付记录ID唯一           |
  | `idx_user_id`        | `user_id`        | `Regular` | ❌        | 用户ID索引，便于查询用户的支付记录     |
  | `idx_order_id`       | `order_id`       | `Unique`  | ❌        | 订单ID索引，便于查询某个订单的支付记录 |
  | `idx_transaction_id` | `transaction_id` | `Unique`  | ✅        | 交易ID唯一索引，确保交易记录的唯一性   |

### 3.8 `area`🌍

+ **用途**：地区表
+ **引擎**：`MyISAM `
+ **字符集**：`utf8mb4`

| 字段名  | 数据类型           | 是否主键 | 是否允许为空 | 默认值 | 备注                                     |
| ------- | ------------------ | -------- | ------------ | ------ | ---------------------------------------- |
| `id`    | `int(10) unsigned` | ✅        | ❌            | `NULL` | 区域ID，自增主键                         |
| `pid`   | `int(10) unsigned` | ❌        | ✅            | `NULL` | 父级区域ID，指向上级区域（用于树形结构） |
| `node`  | `varchar(64)`      | ❌        | ✅            | `NULL` | 区域节点（例如，表示区域的某个标识符）   |
| `name`  | `varchar(32)`      | ❌        | ❌            | `NULL` | 区域名称                                 |
| `level` | `tinyint(4)`       | ❌        | ❌            | `NULL` | 区域级别（如省、市、区等）               |
| `lat`   | `double(8,2)`      | ❌        | ❌            | `NULL` | 纬度值，存储区域的地理坐标               |
| `lng`   | `double(8,2)`      | ❌        | ❌            | `NULL` | 经度值，存储区域的地理坐标               |

+ 索引策略

  | 索引名称             | 索引字段     | 索引类型  | 是否唯一 | 说明                                               |
  | -------------------- | ------------ | --------- | -------- | -------------------------------------------------- |
  | `pk_id`              | `id`         | `PRIMARY` | ✅        | 主键索引，确保区域ID唯一                           |
  | `area_lat_lng_index` | `lat`, `lng` | `Regular` | ❌        | 经纬度索引，便于按地理坐标查询区域                 |
  | `area_pid_index`     | `pid`        | `Regular` | ❌        | 父级区域ID索引，用于树形结构查询                   |
  | `area_name_index`    | `name`       | `Regular` | ❌        | 区域名称索引，便于根据名称查询区域                 |
  | `area_level_index`   | `level`      | `Regular` | ❌        | 区域级别索引，便于按级别查询区域（如省、市、区等） |

### 3.9 `address`📍

- **用途**：地址表
- **引擎**：`InnoDB`
- **字符集**：`utf8mb4`

| 字段名         | 数据类型           | 是否主键 | 是否允许为空 | 默认值                                                  | 备注                                                       |
| -------------- | ------------------ | -------- | ------------ | ------------------------------------------------------- | ---------------------------------------------------------- |
| `id`           | `BIGINT`           | ✅        | ❌            | `AUTO_INCREMENT`                                        | 地址ID，自增主键                                           |
| `user_id`      | `BIGINT`           | ❌        | ❌            | `NULL`                                                  | 关联`users`表的`id`字段，表示该地址所属用户ID              |
| `area_id`      | `int(10) unsigned` | ❌        | ❌            | `NULL`                                                  | 关联`area`表的`id`字段，表示该地址的区域（如省、市、区等） |
| `address`      | `VARCHAR(255)`     | ❌        | ❌            | `NULL`                                                  | 详细地址信息                                               |
| `recipient`    | `VARCHAR(255)`     | ❌        | ❌            | `NULL`                                                  | 收件人姓名                                                 |
| `phone_number` | `VARCHAR(20)`      | ❌        | ❌            | `NULL`                                                  | 收件人电话号码                                             |
| `is_default`   | `TINYINT`          | ❌        | ✅            | `0`                                                     | 是否为默认地址，0表示不是，1表示是                         |
| `created_at`   | `DATETIME`         | ❌        | ✅            | `DEFAULT CURRENT_TIMESTAMP`                             | 地址创建时间                                               |
| `updated_at`   | `DATETIME`         | ❌        | ✅            | `DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP` | 地址更改时间                                               |

+ 索引策略

  | 索引名称         | 索引字段     | 索引类型  | 是否唯一 | 说明                                   |
  | ---------------- | ------------ | --------- | -------- | -------------------------------------- |
  | `pk_id`          | `id`         | `PRIMARY` | ✅        | 主键索引，确保地址ID唯一               |
  | `idx_user_id`    | `user_id`    | `Regular` | ❌        | 用户ID索引，便于查询某个用户的所有地址 |
  | `idx_area_id`    | `area_id`    | `Regular` | ❌        | 区域ID索引，便于根据区域查询地址       |
  | `idx_is_default` | `is_default` | `Regular` | ❌        | 默认地址索引，便于快速查找默认地址     |

------

## **4. 建表脚本**💻

```sql
-- 创建数据库
CREATE DATABASE IF NOT EXISTS MyGoMall CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE MyGoMall;

-- 1. 创建 user 表
CREATE TABLE `user` (
  `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自增',
  `email` VARCHAR(255) DEFAULT NULL COMMENT '用户邮箱（可选）',
  `username` VARCHAR(100) NOT NULL COMMENT '用户名',
  `password_hash` VARCHAR(255) NOT NULL COMMENT '加密存储的用户密码',
  `phone_number` VARCHAR(20) NOT NULL UNIQUE COMMENT '手机号码，唯一索引',
  `address_id` BIGINT DEFAULT NULL COMMENT '用户地址id（可选）',
  `role` TINYINT NOT NULL DEFAULT 0 COMMENT '用户角色（0-普通用户, 1-管理员）',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '账户状态（0-正常, 1-禁用, 2-待审核）',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '账户创建时间',
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '账户更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 索引策略
CREATE INDEX `idx_email` ON `user`(`email`);
CREATE INDEX `idx_phone_number` ON `user`(`phone_number`);

-- 2. 创建 cart 表
CREATE TABLE `cart` (
  `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键，用于区分每条数据',
  `user_id` BIGINT NOT NULL COMMENT '用户ID',
  `product_id` BIGINT NOT NULL COMMENT '商品ID',
  `quantity` BIGINT NOT NULL DEFAULT 0 COMMENT '商品数量',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '商品创建时间',
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '商品信息修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 索引策略
CREATE INDEX `idx_user_product` ON `cart`(`user_id`, `product_id`);

-- 3. 创建 product 表
CREATE TABLE `product` (
  `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '商品ID，自增',
  `name` LONGTEXT NOT NULL COMMENT '商品名称',
  `description` LONGTEXT DEFAULT NULL COMMENT '商品描述',
  `price` BIGINT NOT NULL COMMENT '商品价格',
  `original_price` BIGINT DEFAULT NULL COMMENT '商品原价',
  `images` LONGTEXT DEFAULT NULL COMMENT '商品图片',
  `status` INT NOT NULL DEFAULT 1 COMMENT '商品状态，1表示启用，0表示未启用',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '商品创建时间',
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '商品信息修改时间',
  `deleted_at` DATETIME DEFAULT NULL COMMENT '商品删除时间',
  `stock` INT NOT NULL DEFAULT 0 COMMENT '商品库存数量'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 索引策略
CREATE INDEX `idx_name` ON `product`(`name`);
CREATE INDEX `idx_status` ON `product`(`status`);
CREATE INDEX `idx_stock` ON `product`(`stock`);

-- 4. 创建 category 表
CREATE TABLE `category` (
  `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '商品类别ID，自增',
  `name` VARCHAR(50) NOT NULL COMMENT '商品类别名称',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '商品类别状态，1表示启用，0表示未启用',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '商品类别创建时间',
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '商品类别更新时间',
  `deleted_at` DATETIME DEFAULT NULL COMMENT '商品类别删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 索引策略
CREATE INDEX `idx_name` ON `category`(`name`);
CREATE INDEX `idx_status` ON `category`(`status`);

-- 5. 创建 category_product 表
CREATE TABLE `category_product` (
  `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自增',
  `category_id` BIGINT NOT NULL COMMENT '类别ID，引用category表的id字段',
  `product_id` BIGINT NOT NULL COMMENT '商品ID，引用product表的id字段'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 索引策略
CREATE INDEX `idx_category` ON `category_product`(`category_id`);
CREATE INDEX `idx_product` ON `category_product`(`product_id`);

-- 6. 创建 order 表
CREATE TABLE `order` (
  `id` CHAR(36) NOT NULL PRIMARY KEY COMMENT '订单ID，UUID',
  `user_id` BIGINT NOT NULL COMMENT '用户ID，外键，关联users表的id字段',
  `total_price` BIGINT NOT NULL COMMENT '订单总金额',
  `discount_price` BIGINT DEFAULT 0 COMMENT '优惠金额',
  `actual_price` BIGINT NOT NULL COMMENT '实际支付金额',
  `order_status` TINYINT NOT NULL DEFAULT 0 COMMENT '订单状态（0-待支付, 1-已支付, 2-已发货, 3-已完成, 4-已取消）',
  `payment_status` TINYINT NOT NULL DEFAULT 0 COMMENT '支付状态（0-未支付, 1-已支付, 2-支付失败, 3-退款中, 4-已退款）',
  `payment_method` VARCHAR(20) DEFAULT NULL COMMENT '支付方式（如微信、支付宝等）',
  `address_id` BIGINT DEFAULT NULL COMMENT '收货地址',
  `recipient_name` VARCHAR(255) DEFAULT NULL COMMENT '收件人姓名',
  `phone_number` VARCHAR(20) DEFAULT NULL COMMENT '收件人电话号码',
  `shipping_status` TINYINT NOT NULL DEFAULT 0 COMMENT '物流状态（0-未发货, 1-已发货, 2-已签收）',
  `paid_at` DATETIME DEFAULT NULL COMMENT '订单支付时间',
  `shipped_at` DATETIME DEFAULT NULL COMMENT '发货时间',
  `completed_at` DATETIME DEFAULT NULL COMMENT '订单完成时间',
  `canceled_at` DATETIME DEFAULT NULL COMMENT '订单取消时间',
  `remark` TEXT DEFAULT NULL COMMENT '订单备注',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '订单创建时间',
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '订单更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 索引策略
CREATE INDEX `idx_user_id` ON `order`(`user_id`);
CREATE INDEX `idx_status` ON `order`(`order_status`);
CREATE INDEX `idx_shipping` ON `order`(`shipping_status`);
CREATE INDEX `idx_payment` ON `order`(`payment_status`);

-- 7. 创建 payment 表
CREATE TABLE `payment` (
  `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '支付记录ID，自增主键',
  `user_id` BIGINT NOT NULL COMMENT '外键，关联users表的id字段，表示用户ID',
  `order_id` CHAR(36) NOT NULL COMMENT '外键，关联orders表的id字段，表示订单ID',
  `transaction_id` VARCHAR(36) DEFAULT NULL COMMENT '交易ID（UUID）',
  `amount` BIGINT NOT NULL COMMENT '交易金额',
  `pay_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '交易时间',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 索引策略
CREATE INDEX `idx_user_id` ON `payment`(`user_id`);
CREATE INDEX `idx_order_id` ON `payment`(`order_id`);
CREATE INDEX `idx_transaction_id` ON `payment`(`transaction_id`);

-- 8. 创建 area 表
CREATE TABLE `area` (
  `id` INT(10) UNSIGNED NOT NULL PRIMARY KEY COMMENT '区域ID，自增主键',
  `pid` INT(10) UNSIGNED DEFAULT NULL COMMENT '父级区域ID，指向上级区域（用于树形结构）',
  `node` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '区域节点（例如，表示区域的某个标识符）',
  `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '区域名称',
  `level` TINYINT(4) NOT NULL COMMENT '区域级别（如省、市、区等）',
  `lat` DOUBLE(8,2) NOT NULL COMMENT '纬度值，存储区域的地理坐标',
  `lng` DOUBLE(8,2) NOT NULL COMMENT '经度值，存储区域的地理坐标'
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 索引策略
CREATE INDEX `area_lat_lng_index` ON `area`(`lat`, `lng`);
CREATE INDEX `area_pid_index` ON `area`(`pid`);
CREATE INDEX `area_name_index` ON `area`(`name`);
CREATE INDEX `area_level_index` ON `area`(`level`);

-- 9. 创建 address 表
CREATE TABLE `address` (
  `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '地址ID，自增主键',
  `user_id` BIGINT NOT NULL COMMENT '关联users表的id字段，表示该地址所属用户ID',
  `area_id` INT(10) UNSIGNED NOT NULL COMMENT '关联area表的id字段，表示该地址的区域（如省、市、区等）',
  `address` VARCHAR(255) NOT NULL COMMENT '详细地址信息',
  `recipient` VARCHAR(255) NOT NULL COMMENT '收件人姓名',
  `phone_number` VARCHAR(20) NOT NULL COMMENT '收件人电话号码',
  `is_default` TINYINT NOT NULL DEFAULT 0 COMMENT '是否为默认地址，0表示不是，1表示是',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '地址创建时间',
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '地址更改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 索引策略
CREATE INDEX `idx_user_id` ON `address`(`user_id`);
CREATE INDEX `idx_area_id` ON `address`(`area_id`);
CREATE INDEX `idx_is_default` ON `address`(`is_default`);

```

------

## **5. 关联关系**🔗

| 表名               | 关联字段      | 关系类型 | 说明                                                   |
| ------------------ | ------------- | -------- | ------------------------------------------------------ |
| `user`             | `address_id`  | 逻辑外键 | 关联 `address` 表的 `id` 字段，表示用户的默认地址      |
| `cart`             | `user_id`     | 逻辑外键 | 关联 `user` 表的 `id` 字段，表示购物车属于哪个用户     |
| `cart`             | `product_id`  | 逻辑外键 | 关联 `product` 表的 `id` 字段，表示购物车中包含的商品  |
| `category_product` | `category_id` | 逻辑外键 | 关联 `category` 表的 `id` 字段，表示该商品属于哪个类别 |
| `category_product` | `product_id`  | 逻辑外键 | 关联 `product` 表的 `id` 字段，表示类别下的商品        |
| `order`            | `user_id`     | 逻辑外键 | 关联 `user` 表的 `id` 字段，表示订单属于哪个用户       |
| `order`            | `address_id`  | 逻辑外键 | 关联 `address` 表的 `id` 字段，表示订单使用的收货地址  |
| `payment`          | `user_id`     | 逻辑外键 | 关联 `user` 表的 `id` 字段，表示支付记录所属的用户     |
| `payment`          | `order_id`    | 逻辑外键 | 关联 `order` 表的 `id` 字段，表示支付记录对应的订单    |
| `area`             | `pid`         | 逻辑外键 | 关联同表的 `id` 字段，表示父区域，形成树形结构         |
| `address`          | `user_id`     | 逻辑外键 | 关联 `user` 表的 `id` 字段，表示地址所属的用户         |
| `address`          | `area_id`     | 逻辑外键 | 关联 `area` 表的 `id` 字段，表示该地址的具体区域       |

------

## **6. 性能优化⚡**

#### 6.1 **索引优化** 🔍

通过合理的 **索引策略**，可以显著提高数据库查询的速度。我们在以下字段上创建了索引，以加速查询操作：

- **`user_id`**、**`product_id`**、**`order_id`**：这些字段经常用于查询用户、商品和订单的数据，因此在这些字段上创建了索引。
- **`email`**、**`phone_number`**：为了加速用户查询，我们在用户表的 `email` 和 `phone_number` 字段上建立了索引，提升了查找用户信息的效率。

#### 6.2 **表结构优化** 🏗️

- **分区表**：对于某些数据量较大的表（如订单表、支付记录表等），我们可以采用 **分区表**（Partitioning）策略，将数据按时间或其他维度分区存储，以减少查询时扫描的行数，提升查询性能。
- **冗余字段**：避免表中的冗余字段，确保数据库的正常化，同时通过合适的 **查询缓存** 机制减少不必要的计算。

#### 6.3 **查询优化** 🚀

- **预查询**：对于复杂的查询操作（如涉及多个表的联接查询），我们通过使用 **预查询** 和 **视图**（Views）来减少计算开销，并提高查询效率。
- **LIMIT 和 OFFSET**：在涉及大量数据的查询时，使用分页查询（例如，`LIMIT` 和 `OFFSET`）来分批加载数据，避免一次性加载过多数据导致性能瓶颈。

#### 6.4 **事务和锁优化** 🔒

- **事务管理**：为确保数据一致性，在涉及多个表的操作时使用 **事务管理**，同时采用合适的隔离级别以提高并发性能。
- **行级锁**：采用 **行级锁**（如 InnoDB 的行锁）来避免 **表级锁**，从而提升高并发环境下的数据库性能。

#### 6.5 **数据库连接池** 💧

- 为了应对高并发的请求，MyGoMall 使用了`Kitex`框架内的 **数据库连接池** 来管理数据库连接。通过连接池复用数据库连接，减少了频繁创建和销毁连接的开销，提高了系统的并发处理能力。

#### 6.6 **存储引擎优化** 🛠️

选择合适的存储引擎对数据库性能至关重要。MyGoMall 系统的数据库表使用了不同的存储引擎，主要使用 **InnoDB** 引擎，以下是优化策略：

- **InnoDB 存储引擎**：作为 MySQL 默认的存储引擎，InnoDB 提供了 **ACID** 事务支持、**行级锁** 和 **外键约束**，它非常适合用于需要高并发和事务管理的应用，如订单处理和支付系统。我们将大多数表（如 `order`、`payment`、`user` 等）使用 InnoDB 引擎，以确保数据的完整性和一致性。
- **MyISAM 存储引擎**：对于数据量较少且查询频繁的表（如 `area` 表），选择 **MyISAM 存储引擎**，因为它提供了更高的读取性能，尤其是在只读查询较多的情况下。MyISAM 引擎支持 **表级锁**，可以在特定场景下获得更高的查询效率。
- **优化存储引擎参数**：针对 **InnoDB** 存储引擎，我们还优化了其参数配置，例如：
  - **`innodb_buffer_pool_size`**：增加缓冲池的大小，提升 InnoDB 的数据处理能力。
  - **`innodb_log_file_size`**：调整日志文件大小，减少磁盘 I/O 操作，提升事务性能。
  - **`innodb_flush_log_at_trx_commit`**：调整日志刷写策略，根据系统需求权衡数据的持久性和性能。

通过合理选择和配置存储引擎，MyGoMall 数据库能够根据不同表的特点和访问模式提供最佳的性能支持。
