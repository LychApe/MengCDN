package service

import (
	"MengCDN/internal/model"
	"strings"
)

func GetWLC(mk string) []string {
	WLData := model.CdnWLS(mk)
	WLText := strings.ReplaceAll(WLData, `"`, "")
	return strings.Split(WLText, ",")
}

func GetWLSWC(mk string) string {
	data := model.CdnSW("1", mk)
	if data == "0" {
		return "off"
	} else if data == "1" {
		return "on"
	} else {
		return "off"
	}
}

func GetCDNSWC(mk string) string {
	data := model.CdnSW("0", mk)
	if data == "0" {
		return "off"
	} else if data == "1" {
		return "on"
	} else {
		return "off"
	}
}
