CREATE DATABASE recite_english;
use recite_english;
CREATE TABLE `english_word` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `word` varchar(100) DEFAULT '' COMMENT '英文单词',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  `chinese` varchar(100) DEFAULT '' COMMENT '中文',
  `word_type` int(16) unsigned DEFAULT '0' COMMENT '词性',
  `word_class` int(16) unsigned DEFAULT '0' COMMENT '词汇分级，如四六级单词',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='英语单词';

CREATE TABLE `rc_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='认证管理';

CREATE TABLE `rc_queue` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `word_id` int(10) unsigned DEFAULT '0' COMMENT '单词_id',  
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
   PRIMARY KEY (`id`) USING BTREE ,
   unique index (`created_by`,`word_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='背诵队列';

 