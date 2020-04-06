// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"errors"
	"net/http"
	"sync"
	"time"

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
	name := s.keysPool.Get(key)
	err, exist := s.cacheWrite(c, name)
	if exist && err != nil {
		logrus.Error(err)
		return
	}

	// avoid cache breakdown
	if !exist && s.cache.SetNX(key, "", 5*time.Second) {
		if data, err := fetchData(); err == nil {
			s.success(c, key, data)
		} else {
			s.failure(c, err)
		}
	} else {
		time.Sleep(5 * time.Second)
		if err, _ := s.cacheWrite(c, name); err != nil {
			s.failure(c, err)
			logrus.Error(err)
		}
	}
}

func (s *baseServant) cacheWrite(c *gin.Context, name string) (error, bool) {
	if resp, _ := s.cache.Get(name); resp != "" {
		_, err := c.Writer.WriteString(resp)
		return err, true
	}
	return errors.New("not exist result"), false
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
