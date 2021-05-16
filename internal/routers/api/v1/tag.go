package v1

import (
	"github.com/gin-gonic/gin"
	"web-gin/global"
	"web-gin/internal/service"
	"web-gin/pkg/app"
	"web-gin/pkg/convert"
	"web-gin/pkg/errenum"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @summary 查询标签
// @Produce json
// @Param id query int true "标签id"
// @Success 200 {object} app.SOut "成功"
// @Failure 500 {object} errenum.Resp "失败"
// @Router /api/v1/tags/{id} [GET]
func (t Tag) Get(c *gin.Context) {
	app.NewResp(c).ErrResp(errenum.NewError(505, ""))
	return
}

// @summary 查询标签列表
// @produce json
// @param name body string false "标签"
// @param state body int false "状态"
// @param pageNo body int false "页，默认1"
// @param pageSize body int false "页大小，默认10"
// @Router /api/v1/tags [GET]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	resp := app.NewResp(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Log.ErrorF("app.BindAndValid errs: %v", errs)
		resp.ErrResp(errenum.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	//service
	ser := service.New(c.Request.Context())
	tagCount, countErr := ser.CountTag(&param)
	if countErr != nil {
		global.Log.ErrorF("ser.countTag errenum: %v", countErr)
		resp.ErrResp(errenum.ErrorCountTagFail)
		return
	}
	list, listErr := ser.GetTagList(&param)
	if listErr != nil {
		global.Log.ErrorF("ser.getTagList errenum: %v", listErr)
		resp.ErrResp(errenum.ErrorGetTagListFail)
		return
	}
	sOut := app.PageResp{
		PageNo:    app.GetPage(c),
		PageSize:  app.GetPageSize(c),
		Data:      list,
		TotalRows: tagCount,
	}
	resp.RespWithData(sOut)
	return
}

// @Summary 创建标签
// @Produce json
// @Param id query int true "标签id"
// @Success 200 {object} app.SOut "成功"
// @Failure 500 {object} errenum.Resp "失败"
// @Router /api/v1/tags/ [POST]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	resp := app.NewResp(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Log.ErrorF("app.BindAndValid errs: %v", errors)
		resp.ErrResp(errenum.InvalidParams.WithDetails(errors.Errors()...))
		return
	}
	ser := service.New(c.Request.Context())
	err := ser.CreateTag(&param)
	if err != nil {
		global.Log.ErrorF("ser.CreateTag err: %v", err)
		resp.ErrResp(errenum.ErrorCreateTagFail)
		return
	}
	resp.RespWithData(gin.H{})
	return
}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {object} app.SOut "成功"
// @Failure 500 {object} errenum.Resp "失败"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{
		ID:         convert.StrTo(c.Param("id")).MustUint32(),
		Name:       c.Param("name"),
		ModifiedBy: c.Param("modifiedBy"),
		State:      convert.StrTo(c.Param("state")).MustUint8(),
	}
	resp := app.NewResp(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Log.ErrorF("app.BindAndValid errs: %v", errors)
		resp.ErrResp(errenum.InvalidParams.WithDetails(errors.Errors()...))
		return
	}
	ser := service.New(c.Request.Context())
	err := ser.UpdateTag(&param)
	if err != nil {
		global.Log.ErrorF("ser.UpdateTag err: %v", err)
		resp.ErrResp(errenum.ErrorUpdateTagFail)
		return
	}
	resp.RespWithData(gin.H{})
	return
}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签ID"
// @Success 200 {string} string "成功"
// @Success 200 {object} app.SOut "成功"
// @Failure 500 {object} errenum.Resp "失败"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResp(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Log.ErrorF("app.BindAndValid errs: %v", errors)
		response.ErrResp(errenum.InvalidParams.WithDetails(errors.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Log.ErrorF("svc.DeleteTag err: %v", err)
		response.ErrResp(errenum.ErrorDeleteTagFail)
		return
	}
	response.RespWithData(gin.H{})
	return
}
