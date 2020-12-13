package serializer

import (
	"call-up/model"
)

// RequestInfo 接令请求信息序列化器
type RequestInfo struct {
	CallupName    string `json:"callup_name"`
	Description   string `json:"descrpt"`
	Status        uint8  `json:"status"`
}

// BuildRequestInfoResponse 序列化接令请求查询响应
func BuildRequestInfoResponse(request model.Request, callupName string) Response {
	return Response{
		Code: CodeSuccess,
		Data: RequestInfo{
			CallupName:    callupName,
			Description:   request.Description,
			Status:        request.Status,
		},
	}
}
