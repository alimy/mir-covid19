// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cache

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

func newCacheFake() Cache {
	return &cacheFake{}
}
