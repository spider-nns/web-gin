package service

import "web-gin/internal/model"

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=2,max=55"`
	CreatedBy string `form:"createdBy" binding:"required,min"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}
type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
	PageNo int `form:"pageNo,default=1"`
	PageSize int `form:"pageSize,default=10"`
}

func (ser *Service) CreateTag(param *CreateTagRequest) error {
	return ser.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (ser *Service) CountTag(param *TagListRequest) (int, error) {
	return ser.dao.CountTag(param.Name, param.State)
}

func (ser *Service) GetTagList(param *TagListRequest) ([]*model.Tag, error) {
	return ser.dao.GetTagList(param.Name, param.State,param.PageNo,param.PageSize)
}
