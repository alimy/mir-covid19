// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"sync"

	"github.com/alimy/mir-covid19/internal/xorm"
	"github.com/jinzhu/gorm"
)

var (
	singleServant *baseServant
	onceInit      = &sync.Once{}
)

type baseServant struct {
	db *gorm.DB
}

func myBaseServant() *baseServant {
	onceInit.Do(func() {
		singleServant = &baseServant{
			db: xorm.MyDB(),
		}
	})
	return singleServant
}
