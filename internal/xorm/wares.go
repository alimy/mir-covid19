// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

import "errors"

var (
	errNotReady = errors.New("not ready")
)

// Dataware data service interface
type Dataware interface {
	GetInfoBatch(*InfoBatchArg) (*InfoBatchRes, error)
	GetChinaTotal(*ChinaTotalArg) (*ChinaTotalRes, error)
	GetAreaInfo(*AreaInfoArg) (*AreaInfoRes, error)
	GetForeignTotal(*ForeignTotalArg) (*ForeignTotalRes, error)
	GetForeignInfo(*ForeignInfoArg) (*ForeignInfoRes, error)
	GetForeignHistory(*ForeignHistoryArg) (*ForeignHistoryRes, error)
	GetHubeiInfo(*HubeiInfoArg) (*HubeiInfoRes, error)
	GetRate(*RateArg) (*RateRes, error)
	GetCityInfoByCode(*CityInfoByCodeArg) (*CityInfoByCodeRes, error)
	GetContents(*ContentsArg) (*ContentsRes, error)
	GetAreaContents(*AreaContentsArg) (*AreaContentsRes, error)
	Shutdown()
}
