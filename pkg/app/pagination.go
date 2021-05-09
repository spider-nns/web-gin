package app

import (
	"github.com/gin-gonic/gin"
	"web-gin/global"
	"web-gin/pkg/convert"
)

func GetPage(c *gin.Context) int {
	pageNo := convert.StrTo(c.Query("pageNo")).MustInt()
	if pageNo <= 0 {
		return global.AppSetting.DefaultPageNo
	}
	return global.AppSetting.DefaultPageNo
}
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("pageSize")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
