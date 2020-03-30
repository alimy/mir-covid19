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
	GetInfoBatch() (interface{}, error)
	GetChinaTotal() (interface{}, error)
	GetAreaInfo(*AreaInfoArg) (*AreaInfo, error)
	GetForeignTotal() (interface{}, error)
	GetForeignInfo() (interface{}, error)
	GetForeignHistory() (interface{}, error)
	GetHubeiInfo() (interface{}, error)
	GetRate() (interface{}, error)
	GetCityInfoByCode() (interface{}, error)
	GetContents() (interface{}, error)
	GetAreaContents() (interface{}, error)
	Shutdown()
}
