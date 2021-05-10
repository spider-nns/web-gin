package service

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=2,max=55"`
	CreatedBy string `form:"createdBy" binding:"required,min"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

func (ser *Service) CreateTag(param *CreateTagRequest) error {
	return ser.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}
