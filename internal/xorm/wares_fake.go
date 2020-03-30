// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

type dwFake struct {
	// just empty
}

func (d *dwFake) GetInfoBatch(p *InfoBatchArg) (*InfoBatchRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetChinaTotal(p *ChinaTotalArg) (*ChinaTotalRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetAreaInfo(p *AreaInfoArg) (*AreaInfoRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetForeignTotal(p *ForeignTotalArg) (*ForeignTotalRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetForeignInfo(p *ForeignInfoArg) (*ForeignInfoRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetForeignHistory(p *ForeignHistoryArg) (*ForeignHistoryRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetHubeiInfo(p *HubeiInfoArg) (*HubeiInfoRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetRate(p *RateArg) (*RateRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetCityInfoByCode(p *CityInfoByCodeArg) (*CityInfoByCodeRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetContents(p *ContentsArg) (*ContentsRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) GetAreaContents(p *AreaContentsArg) (*AreaContentsRes, error) {
	// TODO
	return nil, errNotReady
}

func (d *dwFake) Shutdown() {
	// just empty
}

func newDwFake() Dataware {
	return &dwFake{}
}
