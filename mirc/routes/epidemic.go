// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package routes

import (
	"github.com/alimy/mir/v2"
	"github.com/alimy/mir/v2/engine"
)

func init() {
	engine.AddEntry(new(Epidemic))
}

// Epidemic contents fetch interface define
type Epidemic struct {
	Group             mir.Group `mir:"ncovh5api/THPneumoniaOuterDataService"`
	GetInfoBatch      mir.Post  `mir:"getInfoBatch"`
	GetChinaTotal     mir.Post  `mir:"getChinaTotal"`
	GetAreaInfo       mir.Post  `mir:"getAreaInfo"`
	GetForeignTotal   mir.Post  `mir:"getForeignTotal"`
	GetForeignInfo    mir.Post  `mir:"getForeignInfo"`
	GetForeignHistory mir.Post  `mir:"getForeignHistory"`
	GetHubeiInfo      mir.Post  `mir:"getHubeiInfo"`
	GetRate           mir.Post  `mir:"getRate"`
	GetCityInfoByCode mir.Post  `mir:"getCityInfoByCode"`
}
