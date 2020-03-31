// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

import (
	"github.com/alimy/mir-covid19/internal/config"
	"github.com/alimy/mir-covid19/internal/xorm/model"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// NewDataWare return an initialed db
func NewDataWare() Dataware {
	conf := config.MyConfig()
	if conf.Develop.DatawareFake {
		logrus.Info("in dataware fake mode so use faked dataware")
		return newDwFake()
	}

	initGorm()
	dialect, dsn := conf.Database.Dsn()
	db, err := gorm.Open(dialect, dsn)
	if err != nil {
		logrus.Fatal(err)
	}
	initDB(db)
	logrus.Infof("connect database(%s) by dsn: %s", dialect, dsn)

	return newDwGorm(db)
}

func initGorm() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}

	ns := gorm.TheNamingStrategy
	ns.Column = func(s string) string {
		return "f_" + ns.Column(s)
	}
	gorm.AddNamingStrategy(ns)
}

func initDB(db *gorm.DB) {
	// Disable table name's pluralization
	db.SingularTable(true)
	db.AutoMigrate(model.AllModels()...)
}
