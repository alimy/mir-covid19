// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cache

import (
	"time"

	"github.com/alimy/mgo/json"
	"github.com/alimy/mir-covid19/internal/config"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
)

type cacheRedis struct {
	client *redis.Client
}

func (c *cacheRedis) Get(key string) (string, bool) {
	if data, err := c.client.Get(key).Result(); err == nil {
		return data, true
	}
	return "", false
}

func (c *cacheRedis) Put(key string, value interface{}) {
	if data, err := json.Marshal(value); err == nil {
		c.client.Set(key, data, 15*time.Minute)
	} else {
		logrus.Errorf("cache[%s] failure: %s", key, err)
	}
}

func newCacheRedis() Cache {
	conf := config.MyConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
	if _, err := client.Ping().Result(); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("initial redis finish")
	return &cacheRedis{
		client: client,
	}
}
