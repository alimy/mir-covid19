// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package model

import "github.com/jinzhu/gorm"

type City struct {
	gorm.Model
}

type CityAddHistory struct {
	gorm.Model
}

type CityCode struct {
	gorm.Model
}

type CityDaily struct {
	gorm.Model
}

type CityHistory struct {
	gorm.Model
}

type CityMap struct {
	gorm.Model
}

type CityModify struct {
	gorm.Model
}

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
