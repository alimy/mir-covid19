// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package model

import "github.com/jinzhu/gorm"

// AreaHistory base data table schema
//CREATE TABLE `t_base_data` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
//`f_base` varchar(32) NOT NULL COMMENT '维度',
//`f_confirm_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '确诊描述',
//`f_suspect_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '疑似描述',
//`f_dead_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '死亡描述',
//`f_heal_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '治愈描述',
//`f_heavy_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '重症描述',
//`f_now_confirm_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '现有确诊描述',
//`f_now_heavy_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '现有重症描述',
//`f_confirm_add_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '新增确诊描述',
//`f_suspect_add_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '新增疑似描述',
//`f_dead_add_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '新增死亡描述',
//`f_heal_add_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '新增治愈描述',
//`f_heavy_add_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '新增重症描述',
//`f_now_confirm_add_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '现有确诊新增描述',
//`f_now_heavy_add_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '现有重症新增描述',
//`f_is_show_add` int(11) NOT NULL DEFAULT '0' COMMENT '是否展示新增',
//`f_add_no_show_text` varchar(128) NOT NULL DEFAULT '' COMMENT '新增不展示时的文案',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '是否自动更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(32) NOT NULL DEFAULT '' COMMENT '修改人',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='维度数据(目前仅非湖北数据)';
type BaseData struct {
	gorm.Model
}

// Config config table schema
//CREATE TABLE `t_config` (
//`f_day` int(11) NOT NULL DEFAULT '0' COMMENT '日期',
//`f_time` int(11) NOT NULL DEFAULT '0' COMMENT '时间',
//`f_country` varchar(32) NOT NULL DEFAULT '' COMMENT '国家',
//`f_area` varchar(32) NOT NULL DEFAULT '' COMMENT '省份',
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确认',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(120) NOT NULL DEFAULT '' COMMENT '修改人',
//PRIMARY KEY (`f_day`,`f_country`,`f_area`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='配置表';
type Config struct {
	gorm.Model
}

// Content content table schema
//CREATE TABLE `t_content` (
//`id` int(20) NOT NULL AUTO_INCREMENT,
//`ctype` varchar(32) NOT NULL DEFAULT '' COMMENT '类型  shishitongbao-实时通报; ruheyufang-如何预防; paichazhiyin-排查指引',
//`area` varchar(32) NOT NULL DEFAULT '' COMMENT '内容相关省份-地区',
//`public_time` datetime NOT NULL DEFAULT '2020-01-21 19:55:00' COMMENT '发布时间',
//`src` varchar(256) NOT NULL DEFAULT '' COMMENT '来源',
//`title` varchar(1024) NOT NULL DEFAULT '' COMMENT '标题',
//`text` varchar(2048) NOT NULL COMMENT '描叙',
//`img_url` varchar(2048) NOT NULL DEFAULT '' COMMENT '缩略图',
//`url` text NOT NULL COMMENT 'URL',
//`url_type` int(11) NOT NULL DEFAULT '0' COMMENT '0 - 无效; 1 - web view; 2 - 小程序; 3 - 小程序内部跳转',
//`url_appid` varchar(64) NOT NULL DEFAULT '' COMMENT '跳转的小程序APPID',
//`url_appver` varchar(32) NOT NULL DEFAULT 'release' COMMENT '跳转小程序的版本',
//`priority` int(11) NOT NULL DEFAULT '0' COMMENT '优先级',
//`status` int(11) NOT NULL DEFAULT '1' COMMENT '状态 1-有效, 2-失效',
//`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
//`oper` varchar(32) NOT NULL DEFAULT '操作人',
//PRIMARY KEY (`id`),
//KEY `idx_area` (`area`)
//) ENGINE=InnoDB AUTO_INCREMENT=3919 DEFAULT CHARSET=utf8 COMMENT='内容';
type Content struct {
	gorm.Model
}

// DataSchema data schema table
//CREATE TABLE `t_data_schema` (
//`id` int(11) NOT NULL AUTO_INCREMENT,
//`name` varchar(32) NOT NULL COMMENT 'name',
//`dataschema` text NOT NULL COMMENT 'dataschema',
//`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
//PRIMARY KEY (`id`),
//UNIQUE KEY `uni_key` (`name`)
//) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='data schema';
type DataSchema struct {
	gorm.Model
}

// Rate rate table schema
//CREATE TABLE `t_rate` (
//`f_base` varchar(64) NOT NULL DEFAULT '' COMMENT '维度',
//`f_dead_rate_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '病死率描述',
//`f_heavy_rate_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '重症率描述',
//`f_heal_rate_desc` varchar(32) NOT NULL DEFAULT '' COMMENT '治愈率描述',
//`f_seq` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
//`f_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '描述',
//`f_rtx` varchar(32) NOT NULL DEFAULT '' COMMENT '修改人',
//PRIMARY KEY (`f_base`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
type Rate struct {
	gorm.Model
}

func (t *BaseData) TableName() string {
	return "base_data"
}

func (t *Config) TableName() string {
	return "config"
}

func (t *Content) TableName() string {
	return "content"
}

func (t *DataSchema) TableName() string {
	return "data_schema"
}

func (t *Rate) TableName() string {
	return "rate"
}
