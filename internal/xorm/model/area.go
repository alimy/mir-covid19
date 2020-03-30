// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package model

import "github.com/jinzhu/gorm"

// Area area table schema
type Area struct {
	gorm.Model
	// TODO
}

type AreaAddHistory struct {
	gorm.Model
}

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
