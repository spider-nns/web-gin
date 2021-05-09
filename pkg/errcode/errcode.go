package errcode

import (
	"fmt"
	"net/http"
	"web-gin/global"
)

type Resp struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Resp {
	defer func() {
		if err := recover(); err != nil {
			global.Log.PanicF("捕获异常:%s", err)
		}
	}()
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在,请更换一个", code))
	}
	codes[code] = msg
	return &Resp{code: code, msg: msg}
}

func (e Resp) Error() string {
	return fmt.Sprintf("错误码 %d,错误信息: %s", e.Code(), e.Msg())
}

func (e *Resp) Code() int {
	return e.code
}
func (e *Resp) Msg() string {
	return e.msg
}
func (e *Resp) MsgF(args []interface{}) string {
	return fmt.Sprintf(e.msg, args)
}
func (e *Resp) Details() []string {
	return e.details
}

func (e *Resp) WithDetails(details ...string) *Resp {
	e.details = []string{}
	for _, d := range details {
		e.details = append(e.details, d)
	}
	return e
}

func (e *Resp) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	default:
		return http.StatusInternalServerError
	}
}
