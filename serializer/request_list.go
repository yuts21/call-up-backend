package serializer

// RequestListItem 接令请求列表元素序列化器
type RequestListItem struct {
	RequestID  uint   `json:"request_id"`
	CallupName string `json:"callup_name"`
	Status     uint8  `json:"status"`
}
