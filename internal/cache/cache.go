// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cache

import (
	"time"

	"github.com/alimy/mir-covid19/internal/config"
	"github.com/sirupsen/logrus"
)

// Cache cache service interface
type Cache interface {
	Get(key string) (string, bool)
	Put(key string, value interface{})
	SetNX(key string, value interface{}, expiration time.Duration) bool
	PreCache(key string) bool
	PostCache(key string)
}

// NewCache return a cache instance
func NewCache() Cache {
	conf := config.MyConfig()
	if conf.Develop.CacheFake {
		logrus.Info("in cache fake mode so use faked cache")
		return newCacheFake()
	}
	return newCacheRedis()
}
