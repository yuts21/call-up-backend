package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// UserList 用户列表服务
type UserList struct {
	Offset int `form:"offset" json:"offset"`
	Limit  int `form:"limit" json:"limit"`
}

// List 用户列表
func (service *UserList) List(c *gin.Context) serializer.Response {
	if service.Limit == 0 {
		service.Limit = 10
	}

	var total int64 = 0
	if err := model.DB.Model(&model.User{}).Count(&total).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求列表查询失败", err)
	}

	users := []model.User{}
	if err := model.DB.Limit(service.Limit).Offset(service.Offset).Find(&users).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令列表查询失败", err)
	}

	resp := serializer.BuildUserListResponse(users, total)
	resp.Msg = "查询成功"
	return resp
}
