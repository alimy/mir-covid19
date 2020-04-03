// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

import (
	"database/sql"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

var _ = Describe("Repository", func() {
	var wares Dataware
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var db *sql.DB
		var err error

		// db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		// use equal matcher
		// mock sql.DB
		db, mock, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())

		// open mysql db
		gdb, err := gorm.Open("mysql", db)
		Expect(err).ShouldNot(HaveOccurred())

		wares = newDwGorm(gdb)
	})

	AfterEach(func() {
		// make sure all expectations were met
		err := mock.ExpectationsWereMet()
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("get info batch", func() {
		It("found", func() {
			p := &ForeignInfoArg{
				// TODO
			}
			// TODO
			_, _ = wares.GetForeignInfo(p)
		})
	})
})
