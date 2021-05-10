package service

import (
	"context"
	"web-gin/global"
	"web-gin/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	ser := Service{ctx: ctx}
	ser.dao = dao.New(global.DBEngine)
	return ser
}
