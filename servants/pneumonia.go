// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"github.com/alimy/mir-covid19/internal/xorm"
	"github.com/gin-gonic/gin"

	ps "github.com/alimy/mir-covid19/mirc/gen/api/ncovh5api/THPneumoniaService"
)

type pneumonia struct {
	*baseServant
}

func (e *pneumonia) GetContents(c *gin.Context) {
	p := &xorm.ContentsArg{}
	e.handle(c, "Contents", func() (interface{}, error) {
		data, err := e.dw.GetContents(p)
		return data, err
	})
}

func (e *pneumonia) GetAreaContents(c *gin.Context) {
	p := &xorm.AreaContentsArg{}
	e.handle(c, "AreaContents", func() (interface{}, error) {
		data, err := e.dw.GetAreaContents(p)
		return data, err
	})
}

// NewEpidemic return an Pneumonia service instance
func NewPneumonia() ps.Pneumonia {
	return &pneumonia{
		baseServant: myBaseServant(),
	}
}
