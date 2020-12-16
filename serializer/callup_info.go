package serializer

import (
	"call-up/model"
)

// CallupInfo 召集令信息序列化器
type CallupInfo struct {
	Type        uint8  `json:"type"`
	Name        string `json:"name"`
	Description string `json:"descrpt"`
	Capacity    uint   `json:"cap"`
	SponsorID   uint   `json:"sponsor_id"`
	SponsorName string `json:"sponsor_name"`
	Province    string `json:"province"`
	City        string `json:"city"`
	EndDate     int64  `json:"end_date"`
	Status      uint8  `json:"status"`
}

// BuildCallupInfoResponse 序列化召集令信息响应
func BuildCallupInfoResponse(callup model.Callup, sponsor model.User) Response {
	return Response{
		Code: CodeSuccess,
		Data: CallupInfo{
			Type:        callup.Type,
			Name:        callup.Name,
			Description: callup.Description,
			Capacity:    callup.Capacity,
			SponsorID:   callup.SponsorID,
			SponsorName: sponsor.Name,
			Province:    sponsor.Province,
			City:        sponsor.City,
			EndDate:     callup.EndDate.Unix(),
			Status:      callup.Status(),
		},
	}
}
