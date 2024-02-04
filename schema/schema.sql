create database if not exists postbox;

CREATE TABLE IF NOT EXISTS postbox.notification (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',
  `notifi_no` varchar(32) NOT NULL DEFAULT '' COMMENT 'notification no',
  `user_no` varchar(32) NOT NULL DEFAULT '' COMMENT 'user no',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'title',
  `message` varchar(1000) NOT NULL DEFAULT '' COMMENT 'message',
  `status` varchar(10) NOT NULL DEFAULT 'INIT' COMMENT 'Status: INIT, OPENED',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'when the record is created',
  `create_by` varchar(255) NOT NULL DEFAULT '' COMMENT 'who created this record',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'when the record is updated',
  `update_by` varchar(255) NOT NULL DEFAULT '' COMMENT 'who updated this record',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '0-normal, 1-deleted',
  PRIMARY KEY (`id`),
  KEY `user_no_status_idx` (`user_no`, `status`)
) ENGINE=InnoDB COMMENT='Platform Notification';