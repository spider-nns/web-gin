package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
	"web-gin/global"
	"web-gin/internal/model"
	"web-gin/internal/routers"
	"web-gin/pkg/logger"
	"web-gin/pkg/setting"
)

func init() {
	if err := parseConfig(); err != nil {
		log.Fatalf("parse config yaml errenum: %v", err)
	}
	err := initDB()
	if err != nil {
		log.Fatalf("Init DB errenum:%v", err)
	}
	err = initLog()
	if err != nil {
		log.Fatalf("Init Log errenum: %v", err)
	}
}

func parseConfig() error {
	setting, err := setting.Parse()
	if err != nil {
		log.Fatalf("parse viper config errr: %v", err)
	}
	if err := setting.ReadConfig("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err := setting.ReadConfig("App", &global.AppSetting); err != nil {
		return err
	}
	if err := setting.ReadConfig("DataBase", &global.DataBaseSetting); err != nil {
		return err
	}
	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func initDB() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DataBaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func initLog() error {
	global.Log = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   global.AppSetting.LogMaxSize,
		MaxAge:    global.AppSetting.LogMaxTime,
		LocalTime: true,
	}, global.AppSetting.LogPrefix, log.LstdFlags).WithCaller(2)
	return nil
}

// @title web-gin
// @version v1
// @description Go 语言编程之旅:一起用Go写项目
// @termsOfService https://github.com/spider-nns/web-gin
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeOut,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: global.ServerSetting.MaxHeaderBytes,
	}
	global.Log.InfoF("%s:web-gin/%s", "spider", "go")
	//e := errenum.NewError(200, "a")
	//ee := errenum.NewError(200, "b")
	//global.Log.Info(e.Msg())
	//global.Log.Info(ee.Msg())
	s.ListenAndServe()
}
