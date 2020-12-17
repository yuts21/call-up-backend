package serializer

import (
	"call-up/model"
)

// AgencyProfitListItem 中介收益列表元素序列化器
type AgencyProfitListItem struct {
	Date         int64  `json:"date"`
	Province     string `json:"province"`
	City         string `json:"city"`
	Type         uint8  `json:"type"`
	CompletedNum uint   `json:"completed_num"`
	Profit       uint   `json:"profit"`
}

// BuildAgencyProfitList 序列化中介收益列表
func BuildAgencyProfitList(items []model.AgencyProfit) []AgencyProfitListItem {
	agencyProfits := []AgencyProfitListItem{}
	for _, item := range items {
		agencyProfit := AgencyProfitListItem{
			Date:         item.SuccessDate.Unix(),
			Province:     item.Province,
			City:         item.City,
			Type:         item.Type,
			CompletedNum: item.CompletedNum,
			Profit:       item.Profit,
		}
		agencyProfits = append(agencyProfits, agencyProfit)
	}
	return agencyProfits
}

// BuildAgencyProfitListResponse 序列化中介收益列表响应
func BuildAgencyProfitListResponse(agencyProfits []model.AgencyProfit) Response {
	return Response{
		Code: CodeSuccess,
		Data: BuildAgencyProfitList(agencyProfits),
	}
}
