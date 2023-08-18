package model

import (
	"gorm.io/gorm"
)

type SystemConfig struct {
	gorm.Model
	Key   string `gorm:"column:key; type:varchar(255);" json:"key"`
	Value string `gorm:"column:value; type:longtext;" json:"value"`
}

func (table SystemConfig) TableName() string {
	return "SystemConfig"
}

func CdnSW(id string, mk string) (code any) {
	var sc SystemConfig
	var err error
	mk0 := mk + "SW0"
	mk1 := mk + "SW1"
	if id == "0" {
		err = DB.Where("Key = ?", mk0).First(&sc).Error
		if err != nil {
			return 500
		} else {
			return sc.Value
		}
	} else if id == "1" {
		err = DB.Where("Key = ?", mk1).First(&sc).Error
		if err != nil {
			return 500
		} else {
			return sc.Value
		}
	} else {
		return 500
	}
}

func PUTCdnSW(id string, mk string, sw string) (code int) {
	var sc SystemConfig
	var sw1 string
	mk0 := mk + "SW0"
	mk1 := mk + "SW1"
	var maps = make(map[string]interface{})
	if sw == "1" {
		sw1 = "1"
	} else {
		sw1 = "0"
	}
	maps["value"] = sw1

	if id == "0" {
		err := DB.Model(&sc).Where("key = ?", mk0).Updates(&maps).Error
		if err != nil {
			return 500
		}
		return 200
	} else if id == "1" {
		err := DB.Model(&sc).Where("key = ?", mk1).Updates(&maps).Error
		if err != nil {
			return 500
		}
		return 200
	} else {
		return 500
	}
}

func CdnWL(mk string) (data string) {
	var sc WhiteList
	var err error
	mk0 := mk + "WL"
	err = DB.Where("Key = ?", mk0).First(&sc).Error
	if err != nil {
		return "白名单不存在！"
	} else {
		return sc.Value
	}
}

func CdnWLS(mk string) (data string) {
	var sc WhiteList
	DB.Where("Key = ?", mk).First(&sc)
	return sc.Value
}

func PUTCdnWL(mk string, data *WhiteList) (code int) {
	var wl WhiteList
	mk0 := mk + "WL"
	var maps = make(map[string]interface{})
	maps["value"] = data.Value
	err := DB.Model(&wl).Where("Key = ?", mk0).Updates(&maps).Error
	if err != nil {
		return 500
	} else {
		return 200
	}
}
