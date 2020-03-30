package model

import "github.com/jinzhu/gorm"

type TotalHistory struct {
	gorm.Model
}

type TotalAddHistory struct {
	gorm.Model
}

type TotalDaily struct {
	gorm.Model
}

func (t *TotalHistory) TableName() string {
	return "total_history"
}

func (t *TotalAddHistory) TableName() string {
	return "total_add_history"
}

func (t *TotalDaily) TableName() string {
	return "total_daily"
}
