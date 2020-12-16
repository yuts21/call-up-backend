package serializer

// CallupRequestListItem 召集令请求列表元素序列化器
type CallupRequestListItem struct {
	ID            uint   `json:"id"`
	RequesterID   uint   `json:"requester_id"`
	RequesterName string `json:"requester_name"`
	Status        uint8  `json:"status"`
}
