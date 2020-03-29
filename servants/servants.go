// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"
	"sync"

	"github.com/alimy/mgo/keys"
	"github.com/alimy/mir-covid19/internal/cache"
	"github.com/alimy/mir-covid19/internal/xorm"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	singleServant *baseServant
	onceInit      = &sync.Once{}
)

type baseServant struct {
	dw       xorm.Dataware
	cache    cache.Cache
	keysPool keys.StrsPool
}

func (s *baseServant) handle(c *gin.Context, key string, fetchData func() (interface{}, error)) {
	if resp, exist := s.cache.Get(s.keysPool.Get(key)); exist {
		if _, err := c.Writer.WriteString(resp); err != nil {
			logrus.Error(err)
		}
		return
	}
	if data, err := fetchData(); err == nil {
		s.success(c, key, data)
	} else {
		s.failure(c, err)
	}
}

func (s *baseServant) success(c *gin.Context, key string, data interface{}) {
	resp := okResp(data)
	c.JSON(http.StatusOK, resp)
	s.cache.Put(s.keysPool.Get(key), resp)
}

func (s *baseServant) failure(c *gin.Context, err error) {
	resp := errResp(err)
	c.JSON(http.StatusInternalServerError, resp)
}

func myBaseServant() *baseServant {
	onceInit.Do(func() {
		singleServant = &baseServant{
			dw:       xorm.NewDataWare(),
			cache:    cache.NewCache(),
			keysPool: keys.NewStrsPool("covid19:", ":resp", 16),
		}
	})
	return singleServant
}
