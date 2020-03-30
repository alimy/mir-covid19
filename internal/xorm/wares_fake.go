// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

type dwFake struct {
	// just empty
}

func (d *dwFake) GetInfoBatch() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetChinaTotal() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetAreaInfo(p *AreaInfoArg) (*AreaInfo, error) {
	return nil, errNotReady
}

func (d *dwFake) GetForeignTotal() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetForeignInfo() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetForeignHistory() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetHubeiInfo() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetRate() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetCityInfoByCode() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetContents() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetAreaContents() (interface{}, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) Shutdown() {
	// just empty
}

func newDwFake() Dataware {
	return &dwFake{}
}
