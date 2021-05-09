package global

import (
	"web-gin/pkg/logger"
	"web-gin/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DataBaseSetting *setting.DataBaseSetting
	Log             *logger.Logger
)
