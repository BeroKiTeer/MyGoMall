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
                           `name` varchar(40) NOT NULL COMMENT '商品名称',
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

-- 7 订单明细表 (order_item)
CREATE TABLE `order_item` (
                              `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '订单项ID，自增',
                              `order_id` CHAR(36) NOT NULL COMMENT '关联的订单ID，表示该订单的具体商品项',
                              `product_id` BIGINT NOT NULL COMMENT '商品ID，关联商品表',
                              `product_name` VARCHAR(255) NOT NULL COMMENT '商品名称（冗余存储）',
                              `price` BIGINT NOT NULL COMMENT '商品单价（冗余存储）',
                              `quantity` INT NOT NULL DEFAULT 0 COMMENT '购买数量',
                              `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 索引策略
CREATE INDEX `idx_order_id` ON `order_item`(`order_id`);
CREATE INDEX `idx_product_id` ON `order_item`(`product_id`);

-- 8. 创建 payment 表
CREATE TABLE `payment` (
                           `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '支付记录ID，自增主键',
                           `user_id` BIGINT NOT NULL COMMENT '外键，关联users表的id字段，表示用户ID',
                           `order_id` CHAR(36) NOT NULL COMMENT '外键，关联orders表的id字段，表示订单ID',
                           `transaction_id` VARCHAR(36) DEFAULT NULL COMMENT '交易ID（UUID）',
                           `amount` BIGINT NOT NULL COMMENT '交易金额',
                           `pay_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '交易时间',
                           `way` VARCHAR(255) NOT NULL COMMENT '支付方式',
                           `status` BIGINT NOT NULL COMMENT '支付状态',
                           `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 索引策略
CREATE INDEX `idx_user_id` ON `payment`(`user_id`);
CREATE INDEX `idx_order_id` ON `payment`(`order_id`);
CREATE INDEX `idx_transaction_id` ON `payment`(`transaction_id`);

-- 9. 创建 area 表
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

-- 10. 创建 address 表
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
