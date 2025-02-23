## **1. 项目背景**

该数据表用于存储支付（交易）的基本信息

## **2. 设计原则**

- 

## **3. 数据库表结构**

### **3.1 表名：`payments`**

- **用途**：存储支付（交易）的基本信息

- **引擎**：`InnoDB`

- **字符集**：`utf8mb4`

- **主键策略**：AUTO

  | 字段名              | 数据类型           | 是否主键 | 是否允许为空 | 默认值                                             | 备注    |
   |------------------|----------------| ------- | ------------ |-------------------------------------------------|-------|
  | `id`             | `BIGINT`       | ✅       | ❌            | `AUTO_INCREMENT`                                | 主键，用于区分每条数据 |
  | `user_id`        | `BIGINT`       | ❌       | ❌            | `NULL`                                          | 用户ID  |
  | `order_id`       | `VARCHAR(255)` | ❌       | ❌            | `NULL`                                          | 订单ID  |
  | `transaction_id` | `VARCHAR(36)`  | ❌       | ❌            | `NULL`                                          | 交易ID(uuid) |
  | `amount`         | `BIGINT`       | ❌       | ❌            | `NULL`                                          | 交易金额  |
  | `pay_at`         | `DATETIME`     | ❌       | ❌            | `CURRENT_TIMESTAMP`                             | 交易时间  |
  | `created_at`     | `DATETIME`     | ❌        | ❌            | `CURRENT_TIMESTAMP`                             | 创建时间  |
  | `updated_at`     | `DATETIME`     | ❌        | ❌            | `CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP` | 更新时间  |
  | `deleted_at`     | `DATETIME`     | ❌        | ✅            | `NULL`                                          | 删除时间  |

### 3.2 建表语句

```sql
CREATE TABLE payments
(
  id              BIGINT     AUTO_INCREMENT  PRIMARY KEY  COMMENT '主键，区分数据',
  user_id         BIGINT        NOT NULL                  COMMENT '用户ID',
  order_id        VARCHAR(255)  NOT NULL                  COMMENT '订单ID',
  transaction_id  VARCHAR(36)   NOT NULL                  COMMENT '交易ID(uuid)',
  amount          BIGINT        NOT NULL                  COMMENT '交易金额',
  pay_at   DATETIME NOT NULL    DEFAULT CURRENT_TIMESTAMP COMMENT '交易时间',
  created_at   DATETIME NOT NULL   DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at   DATETIME NOT NULL   DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  deleted_at   DATETIME          DEFAULT NULL            COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储支付（交易）的基本信息';

```

------

## **4. 索引设计**

| **索引名称**    | **索引字段**    | **索引类型** | **是否唯一** | **说明** |
|-------------|-------------| ------------ | ------------ |--------|
| `pk_id`     | `id`        | `PRIMARY`    | ✅            | 主键索引   |
| `idx_user`  | `user_id`   | `Regular`    | ❌            | 用户索引   |
| `idx_order` | `order_id`  | `Regular`    | ❌            | 订单索引   |

```sql
CREATE INDEX idx_user    ON payments(user_id);
CREATE INDEX idx_order   ON payments(order_id);
```

------

## **5. 业务约束**

- `id` 仅用于区分每条不同的数据，按照先后顺序。
- 不存在 `update_at` ，因为支付信息不可被修改。

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