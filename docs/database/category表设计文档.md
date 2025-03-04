## **1. 项目背景**

该数据表用于存储商品类别的基本信息

## **2. 设计原则**

-

## **3. 数据库表结构**

### **3.1 表名：`categories`**

- **用途**：存储商品类别的基本信息

- **引擎**：`InnoDB`

- **字符集**：`utf8mb4`

- **主键策略**：自增主键

  | 字段名          | 数据类型          | 是否主键 | 是否允许为空 | 默认值                                             | 备注                  |
    |--------------|---------------|------|--------|-------------------------------------------------|---------------------|
  | `id`         | `BIGINT`      | ✅    | ❌      | `AUTO_INCREMENT`                                | 主键，自增               |
  | `name`       | `VARCHAT(50)` | ❌    | ❌      | `NULL`                                          | 商品类别名称              |
  | `status`     | `TINYINT`     | ❌    | ❌      | `1`                                             | 商品类别状态，1表示启用，0表示未启用 |
  | `created_at` | `TIMESTAMP`   | ❌    | ✅      | `CURRENT_TIMESTAMP`                             | 商品类别创建时间            |
  | `updated_at` | `TIMESTAMP`   | ❌    | ✅      | `CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP` | 商品类别更新时间            |
  | `deleted_at` | `TIMESTAMP`   | ❌    | ✅      | `NULL`                                          | 商品类别删除时间            |

### 3.2 建表语句

```mysql
create table categories
(
  id         bigint auto_increment
    primary key,
  name       varchar(50)                         not null,
  status     tinyint   default 1                 null comment '1: active, 0: inactive',
  created_at timestamp default CURRENT_TIMESTAMP null,
  updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
  deleted_at timestamp                           null,
  constraint categories_pk
    unique (name)
)
  collate = utf8mb4_unicode_ci
  row_format = DYNAMIC;


```



------

## **4. 索引设计**

| **索引名称**                  | **索引字段**              | **索引类型**  | **是否唯一**    | **说明**     |
|---------------------------|-----------------------|-----------|-------------|------------|
| `pk_id`                   | `id`                  | `PRIMARY` | ✅           | 主键索引       |
| `idx_status`              | `status`              | `Regular` | ❌           | 商品类别状态加速查询 |

```mysql
create index idx_status
  on category (status);
```

------

## **5. 业务约束**

- `id` 字段必须唯一，确保每个商品类别被唯一标识。
- `status` 0为未启用，1为启用

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