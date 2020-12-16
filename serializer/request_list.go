package serializer

// RequestListItem 接令请求列表元素序列化器
type RequestListItem struct {
	ID         uint   `json:"id"`
	CallupID   uint   `json:"callup_id"`
	CallupName string `json:"callup_name"`
	Status     uint8  `json:"status"`
}
