package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"web-gin/global"
	"web-gin/internal/routers"
)

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
	s.ListenAndServe()
}

func init() {
	if err := global.ParseConfig(); err != nil {
		log.Fatalf("parse config yaml err: %v", err)
	}
}
