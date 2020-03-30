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

func (d *dwGorm) GetInfoBatch(p *InfoBatchArg) (*InfoBatchRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetChinaTotal(p *ChinaTotalArg) (*ChinaTotalRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetAreaInfo(p *AreaInfoArg) (*AreaInfoRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetForeignTotal(p *ForeignTotalArg) (*ForeignTotalRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetForeignInfo(p *ForeignInfoArg) (*ForeignInfoRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetForeignHistory(p *ForeignHistoryArg) (*ForeignHistoryRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetHubeiInfo(p *HubeiInfoArg) (*HubeiInfoRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetRate(p *RateArg) (*RateRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetCityInfoByCode(p *CityInfoByCodeArg) (*CityInfoByCodeRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetContents(p *ContentsArg) (*ContentsRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwGorm) GetAreaContents(p *AreaContentsArg) (*AreaContentsRes, error) {
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
