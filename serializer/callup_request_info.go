package serializer

import (
	"call-up/model"
)

// CallupRequestInfo 召集令请求信息序列化器
type CallupRequestInfo struct {
	RequesterID   uint   `json:"requester_id"`
	RequesterName string `json:"requester_name"`
	Description   string `json:"descrpt"`
	Status        uint8  `json:"status"`
}

// BuildCallupRequestInfoResponse 序列化召集令请求查询响应
func BuildCallupRequestInfoResponse(request model.Request, requester model.User) Response {
	return Response{
		Code: CodeSuccess,
		Data: CallupRequestInfo{
			RequesterID:   requester.ID,
			RequesterName: requester.Name,
			Description:   request.Description,
			Status:        request.Status,
		},
	}
}
