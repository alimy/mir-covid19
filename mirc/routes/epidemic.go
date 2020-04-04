// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Epidemic))
}

// Epidemic contents fetch interface define
type Epidemic struct {
	Group             Group `mir:"ncovh5api/THPneumoniaOuterDataService"`
	GetInfoBatch      Post  `mir:"getInfoBatch"`
	GetChinaTotal     Post  `mir:"getChinaTotal"`
	GetAreaInfo       Post  `mir:"getAreaInfo"`
	GetForeignTotal   Post  `mir:"getForeignTotal"`
	GetForeignInfo    Post  `mir:"getForeignInfo"`
	GetForeignHistory Post  `mir:"getForeignHistory"`
	GetHubeiInfo      Post  `mir:"getHubeiInfo"`
	GetRate           Post  `mir:"getRate"`
	GetCityInfoByCode Post  `mir:"getCityInfoByCode"`
}
