package serializer

import "call-up/model"

// UserListItem 用户列表元素序列化器
type UserListItem struct {
	ID       uint   `json:"id"`
	UserID   string `json:"user"`
	Type     bool   `json:"type"`
	Name     string `json:"name"`
	Level    uint8  `json:"level"`
	Province string `json:"province"`
	City     string `json:"city"`
}

func BuildUserListResponse(users []model.User, total int64) Response {
	userList := []UserListItem{}
	for _, user := range users {
		userListItem := UserListItem{
			ID:       user.ID,
			UserID:   user.UserID,
			Type:     user.Type,
			Name:     user.Name,
			Level:    user.Level,
			Province: user.Province,
			City:     user.City,
		}
		userList = append(userList, userListItem)
	}
	return BuildListResponse(userList, total)
}
