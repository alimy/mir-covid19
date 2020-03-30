// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package model

import "github.com/jinzhu/gorm"

type TotalHistory struct {
	gorm.Model
}

type TotalAddHistory struct {
	gorm.Model
}

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
