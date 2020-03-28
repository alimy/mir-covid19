// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"sync"

	"github.com/alimy/mir-covid19/internal/xorm"
)

var (
	singleServant *baseServant
	onceInit      = &sync.Once{}
)

type baseServant struct {
	dw xorm.Dataware
}

func myBaseServant() *baseServant {
	onceInit.Do(func() {
		singleServant = &baseServant{
			dw: xorm.NewDataWare(),
		}
	})
	return singleServant
}
