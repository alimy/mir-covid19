package model

import "github.com/jinzhu/gorm"

type SummaryForeign struct {
	gorm.Model
}

type SummaryForeignHistory struct {
	gorm.Model
}

type SummaryChina struct {
	gorm.Model
}

func (t *SummaryForeign) TableName() string {
	return "summary_foreign"
}

func (t *SummaryForeignHistory) TableName() string {
	return "summary_foreign_history"
}

func (t *SummaryChina) TableName() string {
	return "summary_china"
}
