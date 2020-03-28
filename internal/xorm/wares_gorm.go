// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

import "github.com/jinzhu/gorm"

type dwGorm struct {
	*gorm.DB
}

func (d *dwGorm) GetInfoBatch() {
	// TODO
}

func (d *dwGorm) GetChinaTotal() {
	// TODO
}

func (d *dwGorm) GetAreaInfo() *AreaInfo {
	// TODO
	return nil
}

func (d *dwGorm) GetForeignTotal() {
	// TODO
}

func (d *dwGorm) GetForeignInfo() {
	// TODO
}

func (d *dwGorm) GetForeignHistory() {
	// TODO
}

func (d *dwGorm) GetHubeiInfo() {
	// TODO
}

func (d *dwGorm) GetRate() {
	// TODO
}

func (d *dwGorm) GetCityInfoByCode() {
	// TODO
}

func (d *dwGorm) GetContents() {
	// TODO
}

func (d *dwGorm) GetAreaContents() {
	// TODO
}

func newDwGorm(db *gorm.DB) Dataware {
	return &dwGorm{
		DB: db,
	}
}
