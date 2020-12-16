package service

import (
	"call-up/model"
	"call-up/serializer"
	"time"

	"github.com/gin-gonic/gin"
)

// AgencyProfitList 中介收益信息服务
type AgencyProfitList struct {
	StartDate int64  `form:"start_date" json:"start_date" binding:"required"`
	EndDate   int64  `form:"end_date" json:"end_date" binding:"required"`
	Province  string `form:"province" json:"province" binding:"required"`
	City      string `form:"city" json:"city" binding:"required"`
	Type      uint8  `form:"type" json:"type" binding:"required"`
}

// List 中介收益信息函数
func (service *AgencyProfitList) List(c *gin.Context) serializer.Response {
	var agencyProfits []model.AgencyProfit
	if err := model.DB.
		Where("success_date >= ? and success_date <= ?", time.Unix(service.StartDate, 0), time.Unix(service.EndDate, 0)).
		Find(&agencyProfits).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "中介收益查询失败", err)
	}

	resp := serializer.BuildAgencyProfitListResponse(agencyProfits)
	resp.Msg = "查询成功"
	return resp
}
