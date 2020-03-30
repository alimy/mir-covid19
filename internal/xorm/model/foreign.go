package model

import "github.com/jinzhu/gorm"

type ForeignCity struct {
	gorm.Model
}

type ForeignCountry struct {
	gorm.Model
}

type ForeignHistory struct {
	gorm.Model
}

type ForeignHistoryFromNews struct {
	gorm.Model
}

func (t *ForeignCity) TableName() string {
	return "foreign_city"
}

func (t *ForeignCountry) TableName() string {
	return "foreign_country"
}

func (t *ForeignHistory) TableName() string {
	return "foreign_history"
}

func (t *ForeignHistoryFromNews) TableName() string {
	return "foreign_history_fromnews"
}
