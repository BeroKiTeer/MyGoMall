## **1. 项目背景**

该数据表用于存储支付（交易）的基本信息

## **2. 设计原则**

- 

## **3. 数据库表结构**

### **3.1. 订单表 (`orders`)**

这是订单的主表，存储每个订单的基本信息。

| 字段名             | 数据类型     | 约束                                   | 说明                                                         |
| ------------------ | ------------ | -------------------------------------- | ------------------------------------------------------------ |
| `id`               | BIGINT       | 主键、自增                             | 订单ID                                                       |
| `user_id`          | BIGINT       | 非空、索引                             | 购买用户ID                                                   |
| `total_price`      | BIGINT       | 非空                                   | 订单总金额                                                   |
| `discount_price`   | BIGINT       | 默认 0                                 | 优惠金额                                                     |
| `actual_price`     | BIGINT       | 非空                                   | 实际支付金额                                                 |
| `order_status`     | TINYINT      | 默认 0                                 | 订单状态（0-待支付, 1-已支付, 2-已发货, 3-已完成, 4-已取消） |
| `payment_status`   | TINYINT      | 默认 0                                 | 支付状态（0-未支付, 1-已支付, 2-支付失败, 3-退款中, 4-已退款） |
| `payment_method`   | VARCHAR(20)  | 可为空                                 | 支付方式（微信、支付宝、银行卡等）                           |
| `shipping_address` | VARCHAR(255) | 非空                                   | 收货地址                                                     |
| `recipient_name`   | VARCHAR(255) | 非空                                   | 收件人姓名                                                   |
| `phone_number`     | VARCHAR(20)  | 非空                                   | 收件人电话号码                                               |
| `shipping_status`  | TINYINT      | 默认 0                                 | 物流状态（0-未发货, 1-已发货, 2-已签收）                     |
| `created_at`       | DATETIME     | 默认当前时间                           | 订单创建时间                                                 |
| `paid_at`          | DATETIME     | 可为空                                 | 订单支付时间                                                 |
| `shipped_at`       | DATETIME     | 可为空                                 | 发货时间                                                     |
| `completed_at`     | DATETIME     | 可为空                                 | 订单完成时间                                                 |
| `canceled_at`      | DATETIME     | 可为空                                 | 订单取消时间                                                 |
| `updated_at`       | DATETIME     | 默认当前时间，修改订单时为订单更新时间 | 订单更新时间                                                 |
| `remark`           | TEXT         | 可为空                                 | 订单备注                                                     |

------

### **3.2. 订单明细表 (`order_item`)**

用于存储订单中的商品明细信息，每个订单可能包含多个商品。

| 字段名           | 数据类型      | 约束         | 说明                       |
| ---------------- | ------------- | ------------ | -------------------------- |
| `id`             | BIGINT        | 主键、自增   | 订单项ID                   |
| `order_id`       | BIGINT        | 非空、索引   | 关联的订单ID               |
| `product_id`     | BIGINT        | 非空、索引   | 商品ID                     |
| `product_name`   | VARCHAR(255)  | 非空         | 商品名称（冗余存储）       |
| `price`  | BIGINT | 非空         | 商品单价（冗余存储）             |
| `quantity`       | INT           | 非空         | 购买数量                   |
| `created_at`     | DATETIME      | 默认当前时间 | 创建时间                   |
| `updated_at`       | DATETIME     | 默认当前时间，修改订单时为订单更新时间       | 更新时间                                   |
------

### **3.3. 支付表 (`payments`)**

用于记录支付信息，支持订单的支付、退款等操作。

- **用途**：存储支付（交易）的基本信息

- **引擎**：`InnoDB`

- **字符集**：`utf8mb4`

- **主键策略**：AUTO

  | 字段名           | 数据类型       | 是否主键 | 是否允许为空 | 默认值                                          | 备注                   |
  | ---------------- | -------------- | -------- | ------------ | ----------------------------------------------- | ---------------------- |
  | `id`             | `BIGINT`       | ✅        | ❌            | `AUTO_INCREMENT`                                | 主键，用于区分每条数据 |
  | `user_id`        | `BIGINT`       | ❌        | ❌            | `NULL`                                          | 用户ID                 |
  | `order_id`       | `VARCHAR(255)` | ❌        | ❌            | `NULL`                                          | 订单ID                 |
  | `transaction_id` | `VARCHAR(36)`  | ❌        | ❌            | `NULL`                                          | 交易ID(uuid)           |
  | `amount`         | `BIGINT`       | ❌        | ❌            | `NULL`                                          | 交易金额               |
  | `pay_at`         | `DATETIME`     | ❌        | ❌            | `CURRENT_TIMESTAMP`                             | 交易时间               |
  | `created_at`     | `DATETIME`     | ❌        | ❌            | `CURRENT_TIMESTAMP`                             | 创建时间               |
  | `updated_at`     | `DATETIME`     | ❌        | ❌            | `CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP` | 更新时间               |
  | `deleted_at`     | `DATETIME`     | ❌        | ✅            | `NULL`                                          | 删除时间               |

------

### **3.4. 订单状态流转日志 (`order_logs`)**

记录订单状态的变更历史，用于审计和追踪。

| 字段名            | 数据类型   | 约束         | 说明                     |
| ----------------- | ---------- | ------------ | ------------------------ |
| `id`              | `BIGINT`   | 主键、自增   | 日志ID                   |
| `order_id`        | `BIGINT`   | 非空、索引   | 关联订单                 |
| `previous_status` | `TINYINT`  | 非空         | 变更前状态               |
| `current_status`  | `TINYINT`  | 非空         | 变更后状态               |
| `changed_by`      | `BIGINT`   | 非空         | 操作人（用户ID 或 系统） |
| `changed_at`      | `DATETIME` | 默认当前时间 | 状态变更时间             |

------

### **3.5. 订单扩展表 (`order_metadata`)**

存储订单的额外信息，如发票信息、优惠券信息等。

| 字段名       | 数据类型      | 约束         | 说明       |
| ------------ | ------------- | ------------ | ---------- |
| `id`         | `BIGINT`      | 主键、自增   | 记录ID     |
| `order_id`   | `BIGINT`      | 非空、索引   | 订单ID     |
| `key`        | `VARCHAR(50)` | 非空         | 扩展字段名 |
| `value`      | `TEXT`        | 非空         | 扩展字段值 |
| `created_at` | `DATETIME`    | 默认当前时间 | 创建时间   |
| `updated_at` | DATETIME     | 默认当前时间，修改订单时为订单更新时间       | 更新时间                                   |

### 3.2 建表语句

```sql
CREATE TABLE orders (
  id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '订单ID',
  user_id BIGINT NOT NULL COMMENT '购买用户ID',
  total_price BIGINT NOT NULL COMMENT '订单总金额',
  discount_price BIGINT DEFAULT 0 COMMENT '优惠金额',
  actual_price BIGINT NOT NULL COMMENT '实际支付金额',
  order_status TINYINT DEFAULT 0 COMMENT '订单状态（0-待支付, 1-已支付, 2-已发货, 3-已完成, 4-已取消）',
  payment_status TINYINT DEFAULT 0 COMMENT '支付状态（0-未支付, 1-已支付, 2-支付失败, 3-退款中, 4-已退款）',
  payment_method VARCHAR(20) DEFAULT NULL COMMENT '支付方式（微信、支付宝、银行卡等）',
  shipping_address VARCHAR(255) NOT NULL COMMENT '收货地址',
  recipient_name VARCHAR(255) NOT NULL COMMENT '收件人姓名',
  phone_number VARCHAR(20) NOT NULL COMMENT '收件人电话号码',
  shipping_status TINYINT DEFAULT 0 COMMENT '物流状态（0-未发货, 1-已发货, 2-已签收）',
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '订单创建时间',
  paid_at DATETIME DEFAULT NULL COMMENT '订单支付时间',
  shipped_at DATETIME DEFAULT NULL COMMENT '发货时间',
  completed_at DATETIME DEFAULT NULL COMMENT '订单完成时间',
  canceled_at DATETIME DEFAULT NULL COMMENT '订单取消时间',
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '订单更新时间',
  remark TEXT DEFAULT NULL COMMENT '订单备注',
  INDEX idx_user (user_id)
) COMMENT = '订单表，存储每个订单的基本信息';


CREATE TABLE order_items (
  id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '订单项ID',
  order_id BIGINT NOT NULL COMMENT '关联的订单ID',
  product_id BIGINT NOT NULL COMMENT '商品ID',
  product_name VARCHAR(255) NOT NULL COMMENT '商品名称（冗余存储）',
  product_price BIGINT NOT NULL COMMENT '商品单价',
  quantity INT NOT NULL COMMENT '购买数量',
  subtotal_price BIGINT GENERATED ALWAYS AS (product_price * quantity) STORED COMMENT '商品小计',
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  INDEX idx_order (order_id),
  INDEX idx_product (product_id)
) COMMENT = '订单明细表，存储每个订单的商品明细信息';


CREATE TABLE payments (
  id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '支付记录ID',
  user_id BIGINT NOT NULL COMMENT '用户ID',
  order_id BIGINT NOT NULL COMMENT '订单ID',
  transaction_id VARCHAR(36) NOT NULL COMMENT '交易ID(uuid)',
  amount BIGINT NOT NULL COMMENT '交易金额',
  pay_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '交易时间',
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
  INDEX idx_user (user_id),
  INDEX idx_order (order_id)
) COMMENT = '支付表，存储支付（交易）的基本信息';


CREATE TABLE order_logs (
  id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '日志ID',
  order_id BIGINT NOT NULL COMMENT '关联订单ID',
  previous_status TINYINT NOT NULL COMMENT '变更前状态',
  current_status TINYINT NOT NULL COMMENT '变更后状态',
  changed_by BIGINT NOT NULL COMMENT '操作人（用户ID 或 系统）',
  changed_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '状态变更时间',
  INDEX idx_order (order_id)
) COMMENT = '订单状态流转日志，记录订单状态的变更历史';


CREATE TABLE order_metadata (
  id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '记录ID',
  order_id BIGINT NOT NULL COMMENT '订单ID',
  key VARCHAR(50) NOT NULL COMMENT '扩展字段名',
  value TEXT NOT NULL COMMENT '扩展字段值',
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  INDEX idx_order (order_id)
) COMMENT = '订单扩展表，存储订单的额外信息，如发票信息、优惠券信息等';

```

------

## **4. 索引设计**

| **索引名称** | **索引字段** | **索引类型** | **是否唯一** | **说明** |
| ------------ | ------------ | ------------ | ------------ | -------- |
| `pk_id`      | `id`         | `PRIMARY`    | ✅            | 主键索引 |
| `idx_user`   | `user_id`    | `Regular`    | ❌            | 用户索引 |
| `idx_order`  | `order_id`   | `Regular`    | ❌            | 订单索引 |

```sql
CREATE INDEX idx_user    ON payment(user_id);
CREATE INDEX idx_order   ON payment(order_id);
```

------

## **5. 业务约束**

- `id` 仅用于区分每条不同的数据，按照先后顺序。

------

## **6. 数据示例**

```sql

```

------

## **7. 可能的扩展方案**

- 

------

## **8. 关联关系**

| **表名** | **关联字段** | **关系类型** | **说明** |
| -------- | ------------ | ------------ | -------- |
|          |              |              |          |
|          |              |              |          |
|          |              |              |          |

------

## **9. 性能优化**

- 

------

## **10. 备份与恢复策略**

- 