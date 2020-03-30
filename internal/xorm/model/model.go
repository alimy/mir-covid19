// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package model

// AllModels return all models
func AllModels() []interface{} {
	return []interface{}{
		new(Area),
		new(AreaHistory),
		new(AreaAddHistory),
		new(BaseData),
		new(Config),
		new(Content),
		new(DataSchema),
		new(Rate),
		new(City),
		new(CityCode),
		new(CityDaily),
		new(CityHistory),
		new(CityAddHistory),
		new(CityMap),
		new(CityModify),
		new(CityUpdate),
		new(ForeignCity),
		new(ForeignCountry),
		new(ForeignHistory),
		new(ForeignHistoryFromNews),
		new(SummaryChina),
		new(SummaryForeign),
		new(SummaryForeignHistory),
		new(TotalHistory),
		new(TotalAddHistory),
		new(TotalDaily),
	}
}
