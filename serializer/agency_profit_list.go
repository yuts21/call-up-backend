package serializer

import (
	"call-up/model"
)

// AgencyProfit 中介收益列表序列化器
type AgencyProfit struct {
	SuccessDate  int64 `json:"date"`
	CompletedNum uint  `json:"completed_num"`
	Profit       uint  `json:"profit"`
}

// BuildAgencyProfitInfo 序列化中介收益查询
func BuildAgencyProfitInfo(agencyProfit model.AgencyProfit) AgencyProfit {
	return AgencyProfit{
		SuccessDate:  (agencyProfit.SuccessDate).Unix(),
		CompletedNum: agencyProfit.CompletedNum,
		Profit:       agencyProfit.Profit,
	}
}

// BuildAgencyProfitInfoResponse 序列化中介收益查询响应
func BuildAgencyProfitInfoResponse(angencyProfit model.AgencyProfit) Response {
	return Response{
		Code: CodeSuccess,
		Data: BuildAgencyProfitInfo(angencyProfit),
	}
}

// BuildAgencyProfitList 序列化中介收益列表
func BuildAgencyProfitList(items []model.AgencyProfit) (angencyProfits []AgencyProfit) {
	for _, item := range items {
		angencyProfit := BuildAgencyProfitInfo(item)
		angencyProfits = append(angencyProfits, angencyProfit)
	}
	return angencyProfits
}

// BuildAgencyProfitListResponse 序列化中介收益列表响应
func BuildAgencyProfitListResponse(angencyProfits []model.AgencyProfit) Response {
	return Response{
		Code: CodeSuccess,
		Data: BuildAgencyProfitList(angencyProfits),
	}
}
