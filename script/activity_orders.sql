CREATE TABLE `activity_orders` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `member_id` int DEFAULT '0' COMMENT '用户ID',
  `activity_id` int DEFAULT '0' COMMENT '活动ID',
  `status` int DEFAULT '0' COMMENT '0报名1签到2取消3超时失约',
  `seat_number` int DEFAULT '1' COMMENT '座位',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间(借书日期)',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT '2006-01-02 15:04:05' COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动订单表';