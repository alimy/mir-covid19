// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

import (
	"github.com/alimy/mir-covid19/internal/config"
	_ "github.com/jinzhu/gorm"
)

// InitWith initial gorm from config
func InitWith(conf *config.Config) {
	// TODO
}
