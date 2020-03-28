// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	ds "github.com/alimy/mir-covid19/mirc/gen/api/ncovh5api/THPneumoniaOuterDataService"
	"github.com/gin-gonic/gin"
)

type epidemic struct {
	*baseServant
}

func (e *epidemic) GetInfoBatch(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetChinaTotal(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetAreaInfo(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetForeignTotal(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetForeignInfo(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetForeignHistory(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetHubeiInfo(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetRate(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetCityInfoByCode(c *gin.Context) {
	// TODO
}

// NewEpidemic return an Epidemic service instance
func NewEpidemic() ds.Epidemic {
	return &epidemic{
		baseServant: myBaseServant(),
	}
}
