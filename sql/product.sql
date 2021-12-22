-- ----------------------------
-- Table structure for banner
-- ----------------------------

CREATE TABLE IF NOT EXISTS  `hc_banner`(
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `product_id` int unsigned NOT NULL COMMENT '商品ID',
    `url` varchar(255) NOT NULL COMMENT '商品链接',
    `order` tinyint DEFAULT 1 COMMENT '排序',
    `deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否注销',
    `created_at` int(32) COMMENT '创建时间',
    `updated_at` int(32) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `banner_product_id` (`product_id`) USING BTREE
    
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商城轮播图表';

-- ----------------------------
-- Table structure for product
-- ----------------------------

CREATE TABLE IF NOT EXISTS `hc_product`(
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `detail_id` int unsigned DEFAULT NULL COMMENT '详情表id',
    `comment_id` int unsigned DEFAULT NULL COMMENT '评价表id',
    `url` varchar(255) NOT NULL COMMENT '商品链接',
    `des` varchar(255) NOT NULL COMMENT '描述',
    `name` varchar(255) NOT NULL COMMENT '名称',
    `short_name` varchar(255) NOT NULL COMMENT '简称',
    `type` tinyint DEFAULT 1 COMMENT '商品类别：1.',
    `price` varchar(255) DEFAULT NULL COMMENT '价格',
    `sale` int unsigned DEFAULT NULL COMMENT '销量',
    `inventory` int unsigned NOT NULL COMMENT '库存',
    `score` varchar(2) DEFAULT '' COMMENT '评分',
    `discount` varchar(2) DEFAULT '' COMMENT '折扣',
    `deleted` tinyint(1)  DEFAULT 0 COMMENT '是否注销',
    `created_at` int(32) COMMENT '创建时间',
    `updated_at` int(32) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `product_detail_id` (`detail_id`) USING BTREE,
    KEY `product_comment_id` (`comment_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商城产品表'