// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package model

import "github.com/jinzhu/gorm"

type SummaryForeign struct {
	gorm.Model
}

type SummaryForeignHistory struct {
	gorm.Model
}

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
