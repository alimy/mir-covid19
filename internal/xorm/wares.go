// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

// Dataware data service interface
type Dataware interface {
	GetInfoBatch()
	GetChinaTotal()
	GetAreaInfo() *AreaInfo
	GetForeignTotal()
	GetForeignInfo()
	GetForeignHistory()
	GetHubeiInfo()
	GetRate()
	GetCityInfoByCode()
	GetContents()
	GetAreaContents()
}
