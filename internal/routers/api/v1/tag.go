package v1

import (
	"github.com/gin-gonic/gin"
	"web-gin/pkg/app"
	"web-gin/pkg/err"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @summary 查询多个标签
// @Product json
// @Param id query int true "标签id"
// @Success 200 {object} app.SOut "成功"
// @Failure 500 {object} err.Resp "失败"
// @Router /api/v1/tags/{id} [GET]
func (t Tag) Get(c *gin.Context) {
	app.NewResp(c).ErrResp(err.NewError(505, ""))
	return
}

func (t Tag) List(c *gin.Context) {

}

// @Summary 创建标签
// @Produce
func (t Tag) Create(c *gin.Context) {

}
func (t Tag) Update(c *gin.Context) {

}
func (t Tag) Delete(c *gin.Context) {

}
