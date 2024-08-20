package service

import (
	"service/test/testUtils"
	"service/test/testUtils/httpClient"
)

func EventCount() {
	//
	res := AddThemeEventCount("jBurger", "ThemePageCount")
	testUtils.CheckSuccess(res, true)

	//
	res = AddGlobalEventCount("MainPageCount")
	testUtils.CheckSuccess(res, true)
}

func AddThemeEventCount(themeType string, eventName string) string {
	route := "AddThemeEventCount"
	formFields := map[string]string{
		"themeType": themeType,
		"eventName": eventName,
	}

	header := httpClient.GetPrivateKeyHeader()

	return httpClient.SendPOSTRequest(UrlBase+route, formFields, header)
}

func AddGlobalEventCount(eventName string) string {
	route := "AddGlobalEventCount"
	formFields := map[string]string{
		"eventName": eventName,
	}

	header := httpClient.GetPrivateKeyHeader()

	return httpClient.SendPOSTRequest(UrlBase+route, formFields, header)
}

// switch eventName {
// case "MainPageCount":
// 	eventDayCountUpdate.MainPageCount += 1
// case "ThemePageCount":
// 	eventDayCountUpdate.ThemePageCount += 1
// case "PictureCount":
// 	eventDayCountUpdate.PictureCount += 1
// case "RecordCount":
// 	eventDayCountUpdate.RecordCount += 1
// case "SaveCount":
// 	eventDayCountUpdate.SaveCount += 1
// }

// themeMap := map[string]string{
// 	"ayuan":          "一顆ㄚ圓",
// 	"jBurger":        "JBURGER",
// 	"manLaoDie":      "鰻老爹",
// 	"yinLou":         "隱樓",
// 	"deluna":         "德魯納韓式",
// 	"ganDaMen":       "甘答門",
// 	"daKaCoffee":     "打卡咖啡",
// 	"guShanXingChen": "古山星辰",
// 	"taiYan":         "泰雁",
// 	"jinSeSanMai":    "金色三麥",
// }
