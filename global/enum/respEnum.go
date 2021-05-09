package enum

import (
	"web-gin/pkg/errcode"
)

var (
	Success                   = errcode.NewError(200, "成功")
	ServerError               = errcode.NewError(10000000, "服务内部错误")
	InvalidParams             = errcode.NewError(10000001, "入参错误")
	NotFound                  = errcode.NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = errcode.NewError(10000003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = errcode.NewError(10000004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = errcode.NewError(10000005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = errcode.NewError(10000006, "鉴权失败，Token 生成失败")
	TooManyRequests           = errcode.NewError(10000007, "请求过多")
)

