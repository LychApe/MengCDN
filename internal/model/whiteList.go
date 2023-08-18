package model

import "gorm.io/gorm"

// WhiteList 白名单表
type WhiteList struct {
	gorm.Model
	Key   string `gorm:"column:key; type:varchar(255);" json:"key"`
	Value string `gorm:"column:value; type:longtext;" json:"value"`
}

func (table WhiteList) TableName() string {
	return "WhiteList"
}
