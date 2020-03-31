// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package model

import "github.com/jinzhu/gorm"

// ForeignCity foreign city table
//CREATE TABLE `t_foreign_city` (
//`f_country` varchar(128) NOT NULL DEFAULT '' COMMENT '国家',
//`f_city` varchar(128) NOT NULL DEFAULT '' COMMENT '城市',
//`f_city_en` varchar(64) NOT NULL DEFAULT '' COMMENT '城市英文名',
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确认',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heavy` int(11) NOT NULL DEFAULT '0' COMMENT '重症',
//`f_now_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '现有确诊',
//`f_now_confirm_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '现有确诊描述',
//`f_now_heavy` int(11) NOT NULL DEFAULT '0' COMMENT '现有重症',
//`f_now_heavy_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '现有重症描述',
//`f_add_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '新增确认',
//`f_add_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '新增疑似',
//`f_add_dead` int(11) NOT NULL DEFAULT '0' COMMENT '新增死亡',
//`f_add_heal` int(11) NOT NULL DEFAULT '0' COMMENT '新增治愈',
//`f_add_heavy` int(11) NOT NULL DEFAULT '0' COMMENT '新增重症',
//`f_add_confirm_desc` varchar(64) DEFAULT '' COMMENT '新增确诊病例描述',
//`f_add_suspect_desc` varchar(64) DEFAULT '' COMMENT '新增疑似病例描述',
//`f_add_dead_desc` varchar(64) DEFAULT NULL COMMENT '新增死亡病例描述',
//`f_add_heal_desc` varchar(64) DEFAULT '' COMMENT '新增治愈病例描述',
//`f_add_heavy_desc` varchar(64) DEFAULT '' COMMENT '新增重症病例描述',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '本条数据是否自动更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//PRIMARY KEY (`f_country`,`f_city`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='海外城市数据';
type ForeignCity struct {
	gorm.Model
}

// ForeignCountry foreign country table
//CREATE TABLE `t_foreign_country` (
//`f_country` varchar(64) NOT NULL DEFAULT '' COMMENT '国家',
//`f_continent` varchar(16) NOT NULL DEFAULT '' COMMENT '洲际名称',
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '累计确诊',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '累计疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '累计死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '累计治愈',
//`f_confirm_modify` int(11) NOT NULL DEFAULT '0' COMMENT '新增确诊',
//`f_suspect_modify` int(11) NOT NULL DEFAULT '0' COMMENT '新增疑似',
//`f_dead_modify` int(11) NOT NULL DEFAULT '0' COMMENT '新增死亡',
//`f_heal_modify` int(11) NOT NULL DEFAULT '0' COMMENT '新增治愈',
//`f_confirm_modify_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '新增确诊描述',
//`f_suspect_modify_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '新增疑似描述',
//`f_dead_modify_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '新增死亡描述',
//`f_heal_modify_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '新增治愈描述',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '是否更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(120) NOT NULL DEFAULT '' COMMENT '修改人',
//`f_short_name` varchar(64) DEFAULT NULL COMMENT '国家简称',
//`f_is_sub` int(1) DEFAULT '0' COMMENT '海外疫情消息订阅是否推送',
//PRIMARY KEY (`f_country`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
type ForeignCountry struct {
	gorm.Model
}

// ForeignHistory foreign history table
//CREATE TABLE `t_foreign_history` (
//`f_country` varchar(64) NOT NULL DEFAULT '' COMMENT '国家',
//`f_continent` varchar(16) NOT NULL DEFAULT '' COMMENT '洲际名称',
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确诊',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_confirm_modify` int(11) NOT NULL DEFAULT '0' COMMENT '新增确诊',
//`f_suspect_modify` int(11) NOT NULL DEFAULT '0' COMMENT '新增疑似',
//`f_dead_modify` int(11) NOT NULL DEFAULT '0' COMMENT '新增死亡',
//`f_heal_modify` int(11) NOT NULL DEFAULT '0' COMMENT '新增治愈',
//`f_confirm_modify_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '新增确诊描述',
//`f_suspect_modify_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '新增疑似描述',
//`f_dead_modify_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '新增死亡描述',
//`f_heal_modify_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '新增治愈描述',
//`f_date` varchar(32) NOT NULL DEFAULT '' COMMENT '日期',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '是否自动更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(120) NOT NULL DEFAULT '' COMMENT '修改人',
//PRIMARY KEY (`f_country`,`f_date`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
type ForeignHistory struct {
	gorm.Model
}

// ForeignHistoryFromNews foreign history from News table
//CREATE TABLE `t_foreign_history_fromnews` (
//`f_country` varchar(64) NOT NULL DEFAULT '' COMMENT '国家',
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确诊',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_add_confirm` int(11) NOT NULL DEFAULT '-999999' COMMENT '新增确诊',
//`f_date` varchar(32) NOT NULL DEFAULT '' COMMENT '日期',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '是否自动更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(120) NOT NULL DEFAULT '' COMMENT '修改人',
//PRIMARY KEY (`f_country`,`f_date`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
type ForeignHistoryFromNews struct {
	gorm.Model
}

func (t *ForeignCity) TableName() string {
	return "foreign_city"
}

func (t *ForeignCountry) TableName() string {
	return "foreign_country"
}

func (t *ForeignHistory) TableName() string {
	return "foreign_history"
}

func (t *ForeignHistoryFromNews) TableName() string {
	return "foreign_history_fromnews"
}
