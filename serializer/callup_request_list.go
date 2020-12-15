package serializer

// CallupRequestListItem 召集令请求列表元素序列化器
type CallupRequestListItem struct {
	RequestID     uint   `json:"request_id"`
	RequesterID   uint   `json:"requester_id"`
	RequesterName string `json:"requester_name"`
	Status        uint8  `json:"status"`
}
