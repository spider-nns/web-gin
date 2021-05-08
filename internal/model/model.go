package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"web-gin/global"
	"web-gin/pkg/setting"
)

type Model struct {
	ID         uint32 `json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(dbSetting *setting.DataBaseSetting) (*gorm.DB, error) {
	db, err := gorm.Open(dbSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)?charset=%s&parseTime=%t&loc=Local",
		dbSetting.UserName,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.DBName,
		dbSetting.Charset,
		dbSetting.ParseTime))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(dbSetting.MaxIdleCons)
	db.DB().SetMaxOpenConns(dbSetting.MaxOpenCons)
	return db, nil
}
