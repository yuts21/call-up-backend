package serializer

import (
	"call-up/model"
)

// Request 接令请求序列化器
type Request struct {
	RequestID   uint   `json:"request_id"`
	CallupID    uint   `json:"callup_id"`
	RequesterID uint   `json:"requester_id"`
	Description string `json:"desc"`
	Status      uint8  `json:"status"`
}

// BuildRequestShow 序列化接令请求查询
func BuildRequestShow(request model.Request) Request {
	return Request{
		RequestID:   request.ID,
		CallupID:    request.CallupID,
		RequesterID: request.RequesterID,
		Description: request.Description,
		Status:      request.Status,
	}
}

// BuildRequestShowResponse 序列化接令请求查询响应
func BuildRequestShowResponse(request model.Request) Response {
	return Response{
		Code: CodeSuccess,
		Data: BuildRequestShow(request),
	}
}

// BuildRequestList 序列化接令请求列表
func BuildRequestList(items []model.Request) (requests []Request) {
	for _, item := range items {
		request := BuildRequestShow(item)
		requests = append(requests, request)
	}
	return requests
}

// BuildRequestListResponse 序列化接令请求列表响应
func BuildRequestListResponse(items []model.Request) Response {
	return Response{
		Code: CodeSuccess,
		Data: BuildRequestList(items),
	}
}
