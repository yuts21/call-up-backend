package serializer

import (
	"time"
)

//CallupProfitRequestListItem 召集令收益接令列表元素序列化器
type CallupProfitRequestListItem struct {
	ID            uint   `json:"id"`
	RequesterID   uint   `json:"requester_id"`
	RequesterName string `json:"requester_name"`
}

// CallupProfitListItem 召集令收益列表元素序列化器
type CallupProfitListItem struct {
	CallupID          uint                          `json:"callup_id"`
	CallupName        string                        `json:"callup_name"`
	SponsorID         uint                          `json:"sponsor_id"`
	SponsorName       string                        `json:"sponsor_name"`
	Date              int64                         `json:"date"`
	Province          string                        `json:"province"`
	City              string                        `json:"city"`
	Type              uint8                         `json:"type"`
	SponsorProfit     uint                          `json:"sponsor_profit"`
	ParticipantProfit uint                          `json:"participant_profit"`
	Requests          []CallupProfitRequestListItem `json:"requests"`
}

// CallupProfitListItemData 召集令收益列表元素数据
type CallupProfitListItemData struct {
	CallupID          uint
	CallupName        string
	SponsorID         uint
	SponsorName       string
	Date              time.Time
	Province          string
	City              string
	Type              uint8
	SponsorProfit     uint
	ParticipantProfit uint
}

func BuildCallupProfitList(data []CallupProfitListItemData) []CallupProfitListItem {
	var results []CallupProfitListItem
	for _, item := range data {
		resultItem := CallupProfitListItem{
			CallupID:          item.CallupID,
			CallupName:        item.CallupName,
			SponsorID:         item.SponsorID,
			SponsorName:       item.SponsorName,
			Date:              item.Date.Unix(),
			Province:          item.Province,
			City:              item.City,
			Type:              item.Type,
			SponsorProfit:     item.SponsorProfit,
			ParticipantProfit: item.ParticipantProfit,
		}
		results = append(results, resultItem)
	}
	return results
}
