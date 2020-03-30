package model

import "github.com/jinzhu/gorm"

type BaseData struct {
	gorm.Model
}

type Config struct {
	gorm.Model
}

type Content struct {
	gorm.Model
}

type DataSchema struct {
	gorm.Model
}

type Rate struct {
	gorm.Model
}

func (t *BaseData) TableName() string {
	return "base_data"
}

func (t *Config) TableName() string {
	return "config"
}

func (t *Content) TableName() string {
	return "content"
}

func (t *DataSchema) TableName() string {
	return "data_schema"
}

func (t *Rate) TableName() string {
	return "rate"
}
