## **1. 项目背景**

该表是一个中间表，用于建立 categories（商品类别）和 products（商品）之间的多对多关系。其中，一个类别可以包含多个商品，同时一个商品也可以属于多个类别。
## **2. 设计原则**

-

## **3. 数据库表结构**

### **3.1 表名：`category_product`**

- **用途**：建立 categories（商品类别）和 products（商品）之间的多对多关系

- **引擎**：`InnoDB`

- **字符集**：`utf8mb4`

- **主键策略**：自增主键

  | 字段名           | 数据类型     | 是否主键     | 是否允许为空   | 默认值              | 备注                           |
  |---------------|----------|----------|----------|------------------|------------------------------|
  | `id`          | `BIGINT` | ✅        | ❌        | `AUTO_INCREMENT` | 主键，自增                        |
  | `category_id` | `BIGINT` | ❌        | ❌        | `NULL`           | 外键，引用categories表的id字段，表示类别ID |
  | `product_id`  | `BIGINT` | ❌        | ❌        | `NULL`           | 外键，引用products表的id字段，表示商品ID   |

### 3.2 建表语句

```mysql
create table category_product
(
  id          bigint auto_increment
    primary key,
  category_id bigint not null,
  product_id  bigint not null,
  constraint cp_c_id
    foreign key (category_id) references categories (id),
  constraint cp_p_id
    foreign key (product_id) references products (id)
);
```



------

## **4. 索引设计**

| **索引名称**                  | **索引字段**              | **索引类型**  | **是否唯一** | **说明**     |
|---------------------------|-----------------------|-----------|----------|------------|
| `pk_id`                   | `id`                  | `PRIMARY` | ✅        | 主键索引       |
| `idx_products_deleted_at` | `products_deleted_at` | `Regular` | ❌        | 商品删除记录加速查询 |
| `idx_status`              | `status`              | `Regular` | ❌        | 商品状态加速查询   |
| `idx_updated_at`          | `updated_at`          | `Regular` | ❌        | 修改时间加速查询   |
```mysql

create index idx_products_deleted_at
  on products (deleted_at);

create index idx_status
  on products (status);

create index idx_updated_at
  on products (updated_at);
```

------

## **5. 业务约束**

- `id` 字段必须唯一。
- `category_id` 字段是外键，确保category_id在categories表中存在。
- `product_id` 字段是外键，确保product_id在products表中存在。

------

## **6. 数据示例**

```mysql

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