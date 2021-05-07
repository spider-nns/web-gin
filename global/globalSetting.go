package global

import (
	"log"
	"time"
	"web-gin/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DataBaseSetting *setting.DataBaseSetting
)

func ParseConfig() error {
	setting, err := setting.Parse()
	if err != nil {
		log.Fatalf("parse viper config errr: %v", err)
	}
	if err := setting.ReadConfig("Server", &ServerSetting); err != nil {
		return err
	}
	if err := setting.ReadConfig("App", &AppSetting); err != nil {
		return err
	}
	if err := setting.ReadConfig("DataBase", &DataBaseSetting); err != nil {
		return err
	}
	ServerSetting.ReadTimeOut *= time.Second
	ServerSetting.WriteTimeout *= time.Second
	return nil
}
