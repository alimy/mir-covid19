package model

import "github.com/jinzhu/gorm"

// Area area table schema
type Area struct {
	gorm.Model
	// TODO
}

type AreaAddHistory struct {
	gorm.Model
}

type AreaHistory struct {
	gorm.Model
}

func (t *Area) TableName() string {
	return "area"
}

func (t *AreaAddHistory) TableName() string {
	return "area_add_history"
}

func (t *AreaHistory) TableName() string {
	return "area_history"
}
