// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Area area table schema
//CREATE TABLE `t_area` (
//`f_area` varchar(64) NOT NULL DEFAULT '' COMMENT '省份',
//`f_confirm` int(11) NOT NULL DEFAULT '0' COMMENT '确诊',
//`f_suspect` int(11) NOT NULL DEFAULT '0' COMMENT '疑似',
//`f_dead` int(11) NOT NULL DEFAULT '0' COMMENT '死亡',
//`f_heal` int(11) NOT NULL DEFAULT '0' COMMENT '治愈',
//`f_heavy` int(11) NOT NULL DEFAULT '0' COMMENT '重症',
//`f_confirm_add` int(11) NOT NULL DEFAULT '0' COMMENT '新增确诊',
//`f_suspect_add` int(11) NOT NULL DEFAULT '0' COMMENT '新增疑似',
//`f_dead_add` int(11) NOT NULL DEFAULT '0' COMMENT '新增死亡',
//`f_heal_add` int(11) NOT NULL DEFAULT '0' COMMENT '新增治愈',
//`f_heavy_add` int(11) NOT NULL DEFAULT '0' COMMENT '新增重症',
//`f_confirm_add_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '新增确诊描述',
//`f_is_updated` int(11) NOT NULL DEFAULT '0' COMMENT '今日是否已经更新',
//`f_is_show_heal` int(11) NOT NULL DEFAULT '0' COMMENT '是否展示治愈',
//`f_status` int(11) NOT NULL DEFAULT '0' COMMENT '是否自动更新',
//`f_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`f_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`f_rtx` varchar(120) NOT NULL DEFAULT '' COMMENT '修改人',
//PRIMARY KEY (`f_area`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
type Area struct {
	Area           string    `gorm:"primary_key;size:64;not null;comment:省份"`
	Confirm        int       `gorm:"not null;comment:确诊"`
	Suspect        int       `gorm:"not null;comment:疑似"`
	Dead           int       `gorm:"not null;comment:死亡"`
	Heal           int       `gorm:"not null;comment:治愈"`
	Heavy          int       `gorm:"not null;comment:重症"`
	ConfirmAdd     int       `gorm:"not null;comment:新增确诊"`
	SuspectAdd     int       `gorm:"not null;comment:新增疑似"`
	DeadAdd        int       `gorm:"not null;comment:新增死亡"`
	HealAdd        int       `gorm:"not null;comment:新增治愈"`
	HeavyAdd       int       `gorm:"not null;comment:新增重症"`
	ConfirmAddDesc string    `gorm:"size:64;not null;comment:新增确诊描述"`
	IsUpdated      int       `gorm:"not null;comment:今日是否已经更新"`
	IsShowHeal     int       `gorm:"not null;comment:是否展示治愈"`
	Status         int       `gorm:"not null;comment:是否自动更新"`
	Rtx            string    `grom:"size:120;not null;comment:修改人"`
	CreatedAt      time.Time `gorm:"not null;defualt:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt      time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

// AreaAddHistory area add history table schema
//CREATE TABLE `t_area_add_history` (
//`f_area` varchar(64) NOT NULL DEFAULT '' COMMENT '省份',
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
//PRIMARY KEY (`f_area`,`f_date`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
type AreaAddHistory struct {
	gorm.Model
}

// AreaHistory area history table schema
//CREATE TABLE `t_area_history` (
//`f_area` varchar(64) NOT NULL DEFAULT '' COMMENT '省份',
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
//PRIMARY KEY (`f_area`,`f_date`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
type AreaHistory struct {
	gorm.Model
}

func (t *Area) TableName() string {
	return "area"
}

func (t *AreaAddHistory) TableName() string {
	return "area_add_history"
}

func (t *AreaHistory) TableName() string {
	return "area_history"
}
