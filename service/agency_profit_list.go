package service

import (
	"call-up/model"
	"call-up/serializer"
	"time"

	"github.com/gin-gonic/gin"
)

// AgencyProfitList 中介收益信息服务
type AgencyProfitList struct {
	StartDate *int64  `form:"start_date" json:"start_date"`
	EndDate   *int64  `form:"end_date" json:"end_date"`
	Province  *string `form:"province" json:"province"`
	City      *string `form:"city" json:"city"`
	Type      *uint8  `form:"type" json:"type"`
}

// List 中介收益信息函数
func (service *AgencyProfitList) List(c *gin.Context) serializer.Response {
	var agencyProfits []model.AgencyProfit

	db := model.DB.Model(&model.AgencyProfit{})
	if service.StartDate != nil {
		date := time.Unix(*service.StartDate, 0)
		date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
		db = db.Where("success_date >= ?", date)
	}
	if service.EndDate != nil {
		date := time.Unix(*service.EndDate, 0)
		date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
		db = db.Where("success_date <= ?", date)
	}
	if service.Province != nil {
		db = db.Where("province = ?", *service.Province)
	}
	if service.City != nil {
		db = db.Where("city = ?", *service.City)
	}
	if service.Type != nil {
		db = db.Where("type = ?", *service.Type)
	}

	if err := db.Find(&agencyProfits).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "中介收益查询失败", err)
	}

	resp := serializer.BuildAgencyProfitListResponse(agencyProfits)
	resp.Msg = "查询成功"
	return resp
}
