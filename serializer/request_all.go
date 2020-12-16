package serializer

// RequestAllItem 全部接令请求列表元素序列化器
type RequestAllItem struct {
	ID            uint   `json:"id"`
	CallupID      uint   `json:"callup_id"`
	CallupName    string `json:"callup_name"`
	RequesterID   uint   `json:"requester_id"`
	RequesterName string `json:"requester_name"`
	Status        uint8  `json:"status"`
}
