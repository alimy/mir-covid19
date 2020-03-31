// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package model

import "github.com/jinzhu/gorm"

// City city table schema
//CREATE TABLE `t_city` (
//`f_country` varchar(64) NOT NULL DEFAULT '' COMMENT '国家',
//`f_area` varchar(64) NOT NULL DEFAULT '' COMMENT '省份',
//`f_city` varchar(64) NOT NULL DEFAULT '' COMMENT '城市',
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确认',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '是否更新',
//`f_is_show_heal` int(11) NOT NULL DEFAULT '0' COMMENT '是否展示治愈',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(120) NOT NULL DEFAULT '' COMMENT '修改人',
//PRIMARY KEY (`f_country`,`f_area`,`f_city`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='城市表';
type City struct {
	gorm.Model
}

// CityAddHistory city add history table
//CREATE TABLE `t_city_add_history` (
//`f_area` varchar(64) NOT NULL DEFAULT '' COMMENT '省份',
//`f_city` varchar(64) NOT NULL DEFAULT '' COMMENT '城市',
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确诊',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_heavy` int(11) NOT NULL DEFAULT '0' COMMENT '重症',
//`f_date` varchar(32) NOT NULL DEFAULT '' COMMENT '日期',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '是否自动更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(120) NOT NULL DEFAULT '' COMMENT '修改人',
//PRIMARY KEY (`f_area`,`f_city`,`f_date`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
type CityAddHistory struct {
	gorm.Model
}

// CityCode city code table
//CREATE TABLE `t_city_code` (
//`f_code` varchar(64) NOT NULL DEFAULT '' COMMENT 'code',
//`f_city_code` varchar(64) NOT NULL DEFAULT '' COMMENT 'city code',
//`f_province_code` varchar(64) NOT NULL DEFAULT '' COMMENT 'province code',
//`f_area` varchar(128) NOT NULL DEFAULT '' COMMENT 'area名称',
//PRIMARY KEY (`f_code`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='code转换';
type CityCode struct {
	gorm.Model
}

// CityDaily city daily table
//CREATE TABLE `t_city_daily` (
//`f_country` varchar(64) NOT NULL DEFAULT '' COMMENT '国家',
//`f_area` varchar(64) NOT NULL DEFAULT '' COMMENT '省份',
//`f_city` varchar(64) NOT NULL DEFAULT '' COMMENT '城市',
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确认',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_data_time` varchar(32) NOT NULL DEFAULT '' COMMENT '统计时间',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(120) NOT NULL DEFAULT '' COMMENT '修改人',
//PRIMARY KEY (`f_country`,`f_area`,`f_city`,`f_data_time`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='城市每日统计表';
type CityDaily struct {
	gorm.Model
}

// CityHistory city history table
//CREATE TABLE `t_city_history` (
//`f_area` varchar(64) NOT NULL DEFAULT '' COMMENT '省份',
//`f_city` varchar(64) NOT NULL DEFAULT '' COMMENT '城市',
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确诊',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_heavy` int(11) NOT NULL DEFAULT '0' COMMENT '重症',
//`f_date` varchar(32) NOT NULL DEFAULT '' COMMENT '日期',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '是否自动更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(120) NOT NULL DEFAULT '' COMMENT '修改人',
//PRIMARY KEY (`f_area`,`f_city`,`f_date`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
type CityHistory struct {
	gorm.Model
}

// CityMap city map table
//CREATE TABLE `t_city_map` (
//`f_area` varchar(64) NOT NULL DEFAULT '' COMMENT '省份',
//`f_city` varchar(64) NOT NULL DEFAULT '' COMMENT '城市',
//`f_city_1` varchar(64) NOT NULL DEFAULT '' COMMENT '城市别名1',
//`f_city_2` varchar(64) NOT NULL DEFAULT '' COMMENT '城市别名2',
//`f_city_3` varchar(64) NOT NULL COMMENT '城市别名3',
//`f_city_4` varchar(64) NOT NULL COMMENT '城市别名4',
//`f_city_5` varchar(64) NOT NULL COMMENT '城市别名5',
//`f_city_6` varchar(6) NOT NULL COMMENT '城市别名6',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(120) NOT NULL DEFAULT '' COMMENT '修改人',
//`f_city_code` varchar(64) NOT NULL DEFAULT '' COMMENT '城市code',
//`f_province_code` varchar(64) NOT NULL DEFAULT '' COMMENT '省份code',
//`f_screen` int(2) NOT NULL DEFAULT '0' COMMENT '是否屏蔽城市数据，只展示省份',
//PRIMARY KEY (`f_area`,`f_city`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='城市映射表';
type CityMap struct {
	gorm.Model
}

// CityModify city modify table
//CREATE TABLE `t_city_modify` (
//`f_country` varchar(64) NOT NULL DEFAULT '' COMMENT '国家',
//`f_area` varchar(64) NOT NULL DEFAULT '' COMMENT '省份',
//`f_city` varchar(64) NOT NULL DEFAULT '' COMMENT '城市',
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确认',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_import` int(11) NOT NULL DEFAULT '0' COMMENT '境外输入',
//`f_now_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '现有确诊',
//`f_now_heavy` int(11) NOT NULL DEFAULT '0' COMMENT '现有重症',
//`f_is_updated` int(11) NOT NULL DEFAULT '0' COMMENT '今日是已经更新',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '是否更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(120) NOT NULL DEFAULT '' COMMENT '修改人',
//`f_confirm_desc` varchar(64) DEFAULT '' COMMENT '新增确诊病例描述',
//`f_suspect_desc` varchar(64) DEFAULT '' COMMENT '新增疑似病例描述',
//`f_dead_desc` varchar(64) DEFAULT NULL COMMENT '新增死亡病例描述',
//`f_heal_desc` varchar(64) DEFAULT '' COMMENT '新增治愈病例描述',
//`f_import_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '境外输入描叙',
//`f_now_confirm_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '现有确诊描述',
//`f_now_heavy_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '现有重症描述',
//PRIMARY KEY (`f_country`,`f_area`,`f_city`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='城市表';
type CityModify struct {
	gorm.Model
}

// CityUpdate city update table
//CREATE TABLE `t_city_update` (
//`f_country` varchar(64) NOT NULL DEFAULT '' COMMENT '国家',
//`f_area` varchar(64) NOT NULL DEFAULT '' COMMENT '省份',
//`f_city` varchar(64) NOT NULL DEFAULT '' COMMENT '城市',
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确认',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '是否更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(120) NOT NULL DEFAULT '' COMMENT '修改人',
//PRIMARY KEY (`f_country`,`f_area`,`f_city`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='城市更新表';
type CityUpdate struct {
	gorm.Model
}

func (t *City) TableName() string {
	return "city"
}

func (t *CityAddHistory) TableName() string {
	return "city_add_history"
}

func (t *CityCode) TableName() string {
	return "city_code"
}

func (t *CityDaily) TableName() string {
	return "city_daily"
}

func (t *CityHistory) TableName() string {
	return "city_history"
}

func (t *CityMap) TableName() string {
	return "city_map"
}

func (t *CityModify) TableName() string {
	return "city_modify"
}

func (t *CityUpdate) TableName() string {
	return "city_update"
}
