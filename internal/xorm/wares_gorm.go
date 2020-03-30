// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type dwGorm struct {
	*gorm.DB
}

func (d *dwGorm) GetInfoBatch() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetChinaTotal() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetAreaInfo(p *AreaInfoArg) (*AreaInfo, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetForeignTotal() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetForeignInfo() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetForeignHistory() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetHubeiInfo() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetRate() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetCityInfoByCode() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetContents() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetAreaContents() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) Shutdown() {
	if err := d.Close(); err != nil {
		logrus.Error(err)
	}
}

func newDwGorm(db *gorm.DB) Dataware {
	return &dwGorm{
		DB: db,
	}
}
