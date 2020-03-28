// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

type dwFake struct {
	// just empty
}

func (d *dwFake) GetInfoBatch() {
	// TODO
}

func (d *dwFake) GetChinaTotal() {
	// TODO
}

func (d *dwFake) GetAreaInfo() *AreaInfo {
	// TODO
	return nil
}

func (d *dwFake) GetForeignTotal() {
	// TODO
}

func (d *dwFake) GetForeignInfo() {
	// TODO
}

func (d *dwFake) GetForeignHistory() {
	// TODO
}

func (d *dwFake) GetHubeiInfo() {
	// TODO
}

func (d *dwFake) GetRate() {
	// TODO
}

func (d *dwFake) GetCityInfoByCode() {
	// TODO
}

func (d *dwFake) GetContents() {
	// TODO
}

func (d *dwFake) GetAreaContents() {
	// TODO
}

func newDwFake() Dataware {
	return &dwFake{}
}
