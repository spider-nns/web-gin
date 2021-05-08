package global

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"web-gin/internal/model"
	"web-gin/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DataBaseSetting *setting.DataBaseSetting
	DBEngine        *gorm.DB
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

func InitDB() error {
	var err error
	DBEngine, err = model.NewDBEngine(DataBaseSetting)
	if err != nil {
		return err
	}
	return nil
}
