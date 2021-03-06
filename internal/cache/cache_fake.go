// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cache

import (
	"time"
)

// cacheFake used for test cache
type cacheFake struct {
	// just empty
}

func (c *cacheFake) Get(_key string) (string, bool) {
	return "", false
}

func (c *cacheFake) Put(_key string, _value interface{}) {
	// do nothing
}

func (c *cacheFake) SetNX(_key string, _value interface{}, _expiration time.Duration) bool {
	return true
}

func (c *cacheFake) PreCache(key string) (<-chan struct{}, bool) {
	return nil, true
}

func (c *cacheFake) PostCache(key string) {
	// do nothing
}

func newCacheFake() Cache {
	return &cacheFake{}
}
