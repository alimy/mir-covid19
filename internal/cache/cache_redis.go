// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cache

import (
	"sync"
	"time"

	"github.com/alimy/mgo/json"
	"github.com/alimy/mir-covid19/internal/config"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
)

type cacheRedis struct {
	mu        *sync.Mutex
	client    *redis.Client
	preCaches map[string]struct{}
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

func (c *cacheRedis) PreCache(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exist := c.preCaches[key]
	if !exist {
		c.preCaches[key] = struct{}{}
	}
	return !exist
}

func (c *cacheRedis) PostCache(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.preCaches, key)
}

func (c *cacheRedis) SetNX(key string, value interface{}, expiration time.Duration) bool {
	return c.client.SetNX(key, value, expiration).Val()
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
		client:    client,
		mu:        &sync.RWMutex{},
		preCaches: make(map[string]struct{}, 16),
	}
}
