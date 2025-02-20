## **1. 项目背景**

该数据表用于存储支付（交易）的基本信息

## **2. 设计原则**

- 

## **3. 数据库表结构**

### **3.1. 订单表 (`orders`)**

这是订单的主表，存储每个订单的基本信息。

| 字段名                | 数据类型         | 约束                  | 说明                                       |
|--------------------|--------------|---------------------|------------------------------------------|
| `id`               | BIGINT       | 主键、自增               | 订单ID                                     |
| `user_id`          | BIGINT       | 非空、索引               | 购买用户ID                                   |
| `total_price`      | VARCHAR(30)  | 非空                  | 订单总金额                                    |
| `discount_price`   | VARCHAR(30)  | 默认 0.00             | 优惠金额                                     |
| `actual_price`     | VARCHAR(30)  | 非空                  | 实际支付金额                                   |
| `order_status`     | TINYINT      | 默认 0                | 订单状态（0-待支付, 1-已支付, 2-已发货, 3-已完成, 4-已取消）  |
| `payment_status`   | TINYINT      | 默认 0                | 支付状态（0-未支付, 1-已支付, 2-支付失败, 3-退款中, 4-已退款） |
| `payment_method`   | VARCHAR(20)  | 可为空                 | 支付方式（微信、支付宝、银行卡等）                        |
| `shipping_address` | VARCHAR(255) | 非空                  | 收货地址                                     |
| `recipient_name`   | VARCHAR(255) | 非空                  | 收件人姓名                                    |
| `phone_number`     | VARCHAR(20)  | 非空                  | 收件人电话号码                                  |
| `shipping_status`  | TINYINT      | 默认 0                | 物流状态（0-未发货, 1-已发货, 2-已签收）                |
| `created_at`       | DATETIME     | 默认当前时间              | 订单创建时间                                   |
| `paid_at`          | DATETIME     | 可为空                 | 订单支付时间                                   |
| `shipped_at`       | DATETIME     | 可为空                 | 发货时间                                     |
| `completed_at`     | DATETIME     | 可为空                 | 订单完成时间                                   |
| `canceled_at`      | DATETIME     | 可为空                 | 订单取消时间                                   |
| `updated_at`       | DATETIME     | 默认当前时间，修改订单时为订单更新时间 | 订单更新时间                                   |
| `remark`           | TEXT         | 可为空                 | 订单备注                                     |

------

### **3.2. 订单明细表 (`order_items`)**

用于存储订单中的商品明细信息，每个订单可能包含多个商品。

| 字段名           | 数据类型      | 约束         | 说明                       |
| ---------------- | ------------- | ------------ | -------------------------- |
| `id`             | BIGINT        | 主键、自增   | 订单项ID                   |
| `order_id`       | BIGINT        | 非空、索引   | 关联的订单ID               |
| `product_id`     | BIGINT        | 非空、索引   | 商品ID                     |
| `product_name`   | VARCHAR(255)  | 非空         | 商品名称（冗余存储）       |
| `product_price`  | VARCHAR(30)   | 非空         | 商品单价                   |
| `quantity`       | INT           | 非空         | 购买数量                   |
| `subtotal_price` | DECIMAL(10,2) | 计算字段     | `product_price * quantity` |
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
  | `amount`         | `VARCHAR(36)`  | ❌        | ❌            | `NULL`                                          | 交易金额               |
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

```

------

## **4. 索引设计**

| **索引名称** | **索引字段** | **索引类型** | **是否唯一** | **说明** |
| ------------ | ------------ | ------------ | ------------ | -------- |
| `pk_id`      | `id`         | `PRIMARY`    | ✅            | 主键索引 |
| `idx_user`   | `user_id`    | `Regular`    | ❌            | 用户索引 |
| `idx_order`  | `order_id`   | `Regular`    | ❌            | 订单索引 |

```sql
CREATE INDEX idx_user    ON payments(user_id);
CREATE INDEX idx_order   ON payments(order_id);
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