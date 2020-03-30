// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"github.com/alimy/mir-covid19/internal/xorm"
	"github.com/gin-gonic/gin"

	ds "github.com/alimy/mir-covid19/mirc/gen/api/ncovh5api/THPneumoniaOuterDataService"
)

type epidemic struct {
	*baseServant
}

func (e *epidemic) GetInfoBatch(c *gin.Context) {
	p := &xorm.InfoBatchArg{}
	e.handle(c, "InfoBatch", func() (interface{}, error) {
		data, err := e.dw.GetInfoBatch(p)
		return data, err
	})
}

func (e *epidemic) GetChinaTotal(c *gin.Context) {
	p := &xorm.ChinaTotalArg{}
	e.handle(c, "ChinaTotal", func() (interface{}, error) {
		data, err := e.dw.GetChinaTotal(p)
		return data, err
	})
}

func (e *epidemic) GetAreaInfo(c *gin.Context) {
	p := &xorm.AreaInfoArg{}
	e.handle(c, "AreaInfo", func() (interface{}, error) {
		data, err := e.dw.GetAreaInfo(p)
		return data, err
	})
}

func (e *epidemic) GetForeignTotal(c *gin.Context) {
	p := &xorm.ForeignTotalArg{}
	e.handle(c, "ForeignTotal", func() (interface{}, error) {
		data, err := e.dw.GetForeignTotal(p)
		return data, err
	})
}

func (e *epidemic) GetForeignInfo(c *gin.Context) {
	p := &xorm.ForeignInfoArg{}
	e.handle(c, "ForeignInfo", func() (interface{}, error) {
		data, err := e.dw.GetForeignInfo(p)
		return data, err
	})
}

func (e *epidemic) GetForeignHistory(c *gin.Context) {
	p := &xorm.ForeignHistoryArg{}
	e.handle(c, "ForeignHistory", func() (interface{}, error) {
		data, err := e.dw.GetForeignHistory(p)
		return data, err
	})
}

func (e *epidemic) GetHubeiInfo(c *gin.Context) {
	p := &xorm.HubeiInfoArg{}
	e.handle(c, "HubeiInfo", func() (interface{}, error) {
		data, err := e.dw.GetHubeiInfo(p)
		return data, err
	})
}

func (e *epidemic) GetRate(c *gin.Context) {
	p := &xorm.RateArg{}
	e.handle(c, "Rate", func() (interface{}, error) {
		data, err := e.dw.GetRate(p)
		return data, err
	})
}

func (e *epidemic) GetCityInfoByCode(c *gin.Context) {
	p := &xorm.CityInfoByCodeArg{}
	e.handle(c, "CityInfoByCode", func() (interface{}, error) {
		data, err := e.dw.GetCityInfoByCode(p)
		return data, err
	})
}

// NewEpidemic return an Epidemic service instance
func NewEpidemic() ds.Epidemic {
	return &epidemic{
		baseServant: myBaseServant(),
	}
}
