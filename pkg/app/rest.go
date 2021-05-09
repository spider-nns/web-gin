package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"web-gin/global/enum"
)

var (
	DefaultSuccessMessage = "success"
)

type Resp struct {
	Ctx *gin.Context
}

func NewResp(ctx *gin.Context) *Resp {
	return &Resp{ctx}
}

// SOut Global Standard Json Out
type SOut struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
type PageResp struct {
	PageNo    int `json:"PageNo"`
	PageSize  int `json:"pageSize"`
	TotalRows int `json:"totalRows"`
}

func (r *Resp) AllResp(code int, msg string, data interface{}) {
	content := SOut{
		code,
		data,
		msg,
	}
	body, _ := json.Marshal(content)
	r.Ctx.JSON(content.Code, body)
}

func (r *Resp) RespWithData(data interface{}) {
	content := SOut{
		enum.Success.Code(),
		data,
		DefaultSuccessMessage,
	}
	body, _ := json.Marshal(content)
	r.Ctx.JSON(content.Code, body)
}