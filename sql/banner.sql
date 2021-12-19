CREATE TABLE  `hc_banner`(
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