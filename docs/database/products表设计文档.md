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

  | 字段名              | 数据类型          | 是否主键     | 是否允许为空   | 默认值              | 备注       |
    |------------------|---------------|----------|----------|------------------|----------|
  | `id`             | `BIGINT`      | ✅        | ❌        | `AUTO_INCREMENT` | 主键，自增    |
  | `name`           | `LONGTEXT`    | ❌        | ❌        | `NULL`           | 商品名称     |
  | `description`    | `LONGTEXT`    | ❌        | ✅        | `NULL`           | 商品描述     |
  | `price`          | `FLOAT`       | ❌        | ❌        | `NULL`           | 商品价格     |
  | `original_price` | `FLOAT`       | ❌        | ❌        | `NULL`           | 商品原价     |
  | `images`         | `LONGTEXT`    | ❌        | ❌        | `NULL`           | 商品图片     |
  | `status`         | `BIGINT`      | ❌        | ❌        | `NULL`           | 商品状态     |
  | `created_at`     | `DATATIME(3)` | ❌        | ✅        | `NULL`           | 商品创建时间   |
  | `updated_at`     | `DATETIME(3)` | ❌        | ✅        | `NULL`           | 商品信息修改时间 |
  | `deleted_at`     | `DATETIME(3)` | ❌        | ✅        | `NULL`           | 商品删除时间   |
  | `stock`          | `INT`         | ❌        | ❌        | `NULL`           | 商品库存数量   |

### 3.2 建表语句

```mysql
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
  created_at     datetime(3) null,
  updated_at     datetime(3) null,
  deleted_at     datetime(3) null,
  stock          int         null
)
  collate = utf8mb4_unicode_ci
    row_format = DYNAMIC;

```



------

## **4. 索引设计**

| **索引名称**                  | **索引字段**              | **索引类型**  | **是否唯一**    | **说明**   |
|---------------------------|-----------------------|-----------|-------------|----------|
| `pk_id`                   | `id`                  | `PRIMARY` | ✅           | 主键索引     |
| `idx_products_deleted_at` | `products_deleted_at` | `Regular` | ❌           | 删除记录加速查询 |
| `idx_status`              | `status`              | `Regular` | ❌           | 商品状态加速查询 |
| `idx_updated_at`          | `updated_at`          | `Regular` | ❌           | 修改时间加速查询 |

```mysql
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
- `status` 0为已下架，1为上架
- `category_id` 需要关联 `categories` 表的 `id`，外键可选。

------

## **6. 数据示例**

```mysql
INSERT INTO go_test.products (id, name, description, price, original_price, images, status, created_at, updated_at, deleted_at, stock) VALUES (1, '苹果', '一种水果', 2.35, 2.04, 'https://loremflickr.com/400/400?lock=4735001455941464', 1, '2025-02-10 16:50:57.000', '2025-02-10 16:51:06.000', null, 50);
INSERT INTO go_test.products (id, name, description, price, original_price, images, status, created_at, updated_at, deleted_at, stock) VALUES (2, '香蕉', '另一种水果', 4.3, 2.2, null, 1, '2025-02-10 16:50:57.000', '2025-02-10 16:51:06.000', null, 30);

```

------

## **7. 可能的扩展方案**

-

------

## **8. 关联关系**

| **表名**    | **关联字段**      | **关系类型**      | **说明**    |
|-----------|---------------|---------------|-----------|
|           |               |               |           |
|           |               |               |           |
|           |               |               |           |

------

## **9. 性能优化**

-

------

## **10. 备份与恢复策略**

- 