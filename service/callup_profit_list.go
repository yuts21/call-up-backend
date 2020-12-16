package service

import (
	"call-up/model"
	"call-up/serializer"
	"time"

	"github.com/gin-gonic/gin"
)

// CallupProfitList 召集令收益信息服务
type CallupProfitList struct {
	StartDate *int64  `form:"start_date" json:"start_date"`
	EndDate   *int64  `form:"end_date" json:"end_date"`
	Province  *string `form:"province" json:"province"`
	City      *string `form:"city" json:"city"`
	Type      *uint8  `form:"type" json:"type"`
	Offset    int     `form:"offset" json:"offset"`
	Limit     int     `form:"limit" json:"limit"`
}

// List 召集令收益信息函数
func (service *CallupProfitList) List(c *gin.Context) serializer.Response {
	if service.Limit == 0 {
		service.Limit = 10
	}

	db := model.DB.Model(&model.SuccessCallupDetail{}).
		Joins("join callups on success_callup_details.callup_id = callups.id").
		Joins("join users on callups.sponsor_id = users.id")
	if service.StartDate != nil {
		db = db.Where("success_callup_details.date >= ?", time.Unix(*service.StartDate, 0))
	}
	if service.EndDate != nil {
		db = db.Where("success_callup_details.date <= ?", time.Unix(*service.EndDate, 0))
	}
	if service.Province != nil {
		db = db.Where("users.province = ?", *service.Province)
	}
	if service.City != nil {
		db = db.Where("users.city = ?", *service.City)
	}
	if service.Type != nil {
		db = db.Where("callups.type = ?", *service.Type)
	}

	var total int64 = 0
	if err := db.Count(&total).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令收益查询失败", err)
	}

	var resultdata []serializer.CallupProfitListItemData
	if err := db.
		Select(`success_callup_details.callup_id as callup_id, callups.name as callup_name, 
		callups.sponsor_id as sponsor_id, users.name as sponsor_name, success_callup_details.date as date,
		users.province as province, users.city as city, callups.type as type, 
		success_callup_details.sponsor_profit as sponsor_profit, 
		success_callup_details.participant_profit as participant_profit`).
		Limit(service.Limit).Offset(service.Offset).Scan(&resultdata).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令收益查询失败", err)
	}
	results := serializer.BuildCallupProfitList(resultdata)

	for i := range results {
		if err := model.DB.Model(&model.Request{}).Select("requests.id as id, requests.requester_id as requester_id, users.name as requester_name").
			Joins("join users on requests.requester_id = users.id").
			Where("requests.callup_id = ? and requests.status = ?", results[i].CallupID, model.Agreed).
			Scan(&results[i].Requests).Error; err != nil {
				return serializer.Err(serializer.CodeDBError, "接令请求查询失败", err)
		}
	}

	resp := serializer.BuildListResponse(results, total)
	resp.Msg = "查询成功"
	return resp
}
