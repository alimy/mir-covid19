// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package model

import "github.com/jinzhu/gorm"

// SummaryForeign summary foreign table
//CREATE TABLE `t_summary_foreign` (
//`id` int(11) NOT NULL AUTO_INCREMENT,
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确认',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heavy` int(11) NOT NULL DEFAULT '0' COMMENT '重症',
//`f_now_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '现有确诊',
//`f_now_confirm_desc` varchar(1024) NOT NULL COMMENT '现有确诊描述',
//`f_now_heavy` int(11) NOT NULL DEFAULT '0' COMMENT '现有重症',
//`f_now_heavy_desc` varchar(1024) NOT NULL COMMENT '现有重症描述',
//`f_add_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '新增确认',
//`f_add_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '新增疑似',
//`f_add_dead` int(11) NOT NULL DEFAULT '0' COMMENT '新增死亡',
//`f_add_heal` int(11) NOT NULL DEFAULT '0' COMMENT '新增治愈',
//`f_add_heavy` int(11) NOT NULL DEFAULT '0' COMMENT '新增重症',
//`f_add_now_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '新增现有确诊',
//`f_add_confirm_desc` varchar(1024) DEFAULT NULL COMMENT '新增确诊病例描述',
//`f_add_suspect_desc` varchar(1024) DEFAULT NULL COMMENT '新增疑似病例描述',
//`f_add_dead_desc` varchar(1024) DEFAULT NULL COMMENT '新增死亡病例描述',
//`f_add_heal_desc` varchar(1024) DEFAULT NULL COMMENT '新增治愈病例描述',
//`f_add_heavy_desc` varchar(1024) DEFAULT NULL COMMENT '新增重症病例描述',
//`f_add_now_confirm_desc` varchar(1024) DEFAULT NULL COMMENT '新增现有确诊描述',
//`f_data_time` varchar(128) NOT NULL DEFAULT '' COMMENT '统计时间',
//`f_is_show_add` int(2) NOT NULL DEFAULT '0' COMMENT '是否展示新增数据',
//`f_add_no_show_text` varchar(1024) DEFAULT NULL COMMENT '新增不展示提示文案（通用）',
//`f_add_no_show_text_mini` varchar(1024) DEFAULT NULL COMMENT '新增不展示提示文案（小程序）',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '本条数据是否自动更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='海外数据汇总';
type SummaryForeign struct {
	gorm.Model
}

// SummaryForeignHistory summary foreign history table
//CREATE TABLE `t_summary_foreign_history` (
//`f_date` varchar(32) NOT NULL DEFAULT '' COMMENT '日期 20200309',
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确诊',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_add_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '新增确诊',
//`f_dead_rate` varchar(8) NOT NULL DEFAULT '0' COMMENT '死亡率',
//`f_heal_rate` varchar(8) NOT NULL DEFAULT '0' COMMENT '治愈率',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '是否自动更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(120) NOT NULL DEFAULT '' COMMENT '修改人',
//PRIMARY KEY (`f_date`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='海外历史数据';
type SummaryForeignHistory struct {
	gorm.Model
}

// SummaryChina summary china table
//CREATE TABLE `t_summery_china` (
//`id` int(20) NOT NULL AUTO_INCREMENT,
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确认',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_import` int(11) NOT NULL DEFAULT '0' COMMENT '境外输入',
//`f_import_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '境外输入描述',
//`f_now_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '现有确诊',
//`f_now_confirm_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '现有确诊描述',
//`f_now_heavy` int(11) NOT NULL DEFAULT '0' COMMENT '现有重症',
//`f_now_heavy_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '现有重症描述',
//`f_data_time` varchar(128) NOT NULL DEFAULT '' COMMENT '统计时间',
//`f_virus_info` text NOT NULL COMMENT '病毒信息',
//`f_data_from` text NOT NULL COMMENT '数据来源',
//`f_area_data_autoupdate` int(11) NOT NULL DEFAULT '0' COMMENT '省份数据是否自动更新',
//`f_is_show_add` int(11) NOT NULL DEFAULT '0' COMMENT '是否展示新增数据',
//`f_add_no_show_text` varchar(128) NOT NULL DEFAULT '' COMMENT '新增数据不展示时的文案',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '本条数据是否自动更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(32) DEFAULT NULL COMMENT '修改人',
//`f_is_auto_daily_update` int(1) NOT NULL DEFAULT '0' COMMENT '每日新增是否自动更新',
//`f_notice` text NOT NULL COMMENT '通告',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='全国数据';
type SummaryChina struct {
	gorm.Model
}

func (t *SummaryForeign) TableName() string {
	return "summary_foreign"
}

func (t *SummaryForeignHistory) TableName() string {
	return "summary_foreign_history"
}

func (t *SummaryChina) TableName() string {
	return "summary_china"
}
