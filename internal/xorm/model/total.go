// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package model

import "github.com/jinzhu/gorm"

// TotalHistory total history table
//CREATE TABLE `t_total_history` (
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确认',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_heavy` int(11) NOT NULL DEFAULT '-999999' COMMENT '重症',
//`f_import` int(11) NOT NULL DEFAULT '0' COMMENT '境外输入',
//`f_date` varchar(32) NOT NULL DEFAULT '' COMMENT '统计时间',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '是否自动更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(32) DEFAULT NULL COMMENT '修改人',
//PRIMARY KEY (`f_date`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='全国每日统计表，同步于inews';
type TotalHistory struct {
	gorm.Model
}

// TotalAddHistory total add history table
//CREATE TABLE `t_total_add_history` (
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确认',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_heavy` int(11) NOT NULL DEFAULT '-999999' COMMENT '重症',
//`f_import` int(11) NOT NULL DEFAULT '0' COMMENT '境外输入',
//`f_date` varchar(32) NOT NULL DEFAULT '' COMMENT '统计时间',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '是否自动更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(32) NOT NULL DEFAULT '' COMMENT '修改人',
//PRIMARY KEY (`f_date`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='全国新增每日统计表，同步于inews';
type TotalAddHistory struct {
	gorm.Model
}

// TotalDaily total daily table
//CREATE TABLE `t_total_daily` (
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确认',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_data_time` varchar(32) NOT NULL DEFAULT '' COMMENT '统计时间',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(32) DEFAULT NULL COMMENT '修改人',
//PRIMARY KEY (`f_data_time`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='全国每日统计表';
type TotalDaily struct {
	gorm.Model
}

func (t *TotalHistory) TableName() string {
	return "total_history"
}

func (t *TotalAddHistory) TableName() string {
	return "total_add_history"
}

func (t *TotalDaily) TableName() string {
	return "total_daily"
}
