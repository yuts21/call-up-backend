package serializer

import "call-up/model"

// CallupListItem 召集令列表元素序列化器
type CallupListItem struct {
	ID       uint   `json:"id"`
	Type     uint8  `json:"type"`
	Name     string `json:"name"`
	Capacity uint   `json:"cap"`
	EndDate  int64  `json:"end_date"`
	Status   uint8  `json:"status"`
}

func BuildCallupListResponse(callups []model.Callup, total int64) Response {
	var callupList []CallupListItem
	for _, callup := range callups {
		callupListItem := CallupListItem{
			ID: callup.ID,
			Type: callup.Type,
			Name: callup.Name,
			Capacity: callup.Capacity,
			EndDate: callup.EndDate.Unix(),
			Status: callup.Status(),
		}
		callupList = append(callupList, callupListItem)
	}
	return BuildListResponse(callupList, total)
}