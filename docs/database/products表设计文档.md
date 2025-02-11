## **1. 项目背景**

该数据表用于存储商品的基本信息

## **2. 设计原则**

- 

## **3. 数据库表结构**

### **3.1 表名：`products`**

- **用途**：存储商品的基本信息

- **引擎**：`InnoDB` 

- **字符集**：`utf8mb4`

- **主键策略**：自增主键

  | 字段名              | 数据类型          | 是否主键 | 是否允许为空 | 默认值                                             | 备注       |
  |------------------|---------------| -------- | ------- |-------------------------------------------------|----------|
  | `id`             | `BIGINT`      | ✅        | ❌       | `AUTO_INCREMENT`                                | 主键，自增    |
  | `name`           | `LONGTEXT`    | ❌        | ❌       | `NULL`                                          | 商品名称     |
  | `description`    | `LONGTEXT`    | ❌        | ✅        | `NULL`                                          | 商品描述     |
  | `price`          | `FLOAT`       | ❌        | ❌       | `NULL`                                          | 商品价格     |
  | `original_price` | `FLOAT`       | ❌        | ❌        | `NULL`                                          | 商品原价     |
  | `images`         | `LONGTEXT`    | ❌        | ❌        | `NULL`                                          | 商品图片     |
  | `status`         | `BIGINT`      | ❌        | ❌       | `NULL`                                          | 商品状态     |
  | `created_at`     | `DATATIME(3)` | ❌        | ✅        | `CURRENT_TIMESTAMP`                             | 商品创建时间   |
  | `updated_at`     | `DATETIME(3)` | ❌        | ✅       | `CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP` | 商品信息修改时间 |
  | `deleted_at`     | `DATETIME(3)` | ❌        | ✅       | `CURRENT_TIMESTAMP`                             | 商品删除时间   |
  | `category_id`    | `BIGINT`      | ❌        | ❌        | `NULL`                                          | 商品类别ID   |
- | `stock`          | `INT`         | ❌        | ❌       | `NULL`                                          | 商品库存数量   |

### 3.2 建表语句

```sql
create table products
(
  id             bigint auto_increment
        primary key,
  name           longtext    null,
  description    longtext    null,
  price          float       null,
  original_price float       null,
  images         longtext    null,
  status         bigint      null,
  created_at     datetime(3) CURRENT_TIMESTAMP,
  updated_at     datetime(3) CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at     datetime(3) CURRENT_TIMESTAMP,
  category_id    bigint      null,
  stock          int         null
)
  collate = utf8mb4_unicode_ci
    row_format = DYNAMIC;

```



------

## **4. 索引设计**

| **索引名称**                  | **索引字段**   | **索引类型**  | **是否唯一** | **说明**   |
|---------------------------| -------------- |-----------| ------------ |----------|
| `pk_id`                   | `id`           | `PRIMARY` | ✅            | 主键索引     |
| `idx_products_deleted_at` | `products_deleted_at`        | `Regular` | ❌            | 删除记录加速查询 |
| `idx_status`               | `status` | `Regular` | ❌            | 商品状态加速查询 |
| `idx_updated_at`               | `updated_at` | `Regular` | ❌            | 修改时间加速查询 |

```sql
create index idx_products_deleted_at
  on go_test.products (deleted_at);

create index idx_status
  on go_test.products (status);

create index idx_updated_at
  on go_test.products (updated_at);
```

------

## **5. 业务约束**

- `id` 字段必须唯一，确保每个商品被唯一标识。
- `price` 不能为负数，防止数据错误。
- `original_price` 不能为负数，防止数据错误。
- `stock` 不能为负数，防止数据错误。
- `status` **??????**
- `category_id` 需要关联 `categories` 表的 `id`，外键可选。

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