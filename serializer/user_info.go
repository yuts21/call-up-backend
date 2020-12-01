package serializer

import "call-up/model"

// UserInfo 用户信息序列化器
type UserInfo struct {
	UserID       string `json:"user"`
	Type         bool   `json:"type"`
	Name         string `json:"name"`
	IDType       uint8  `json:"id_type"`
	IDNumber     string `json:"id_number"`
	Phone        string `json:"phone"`
	Level        uint8  `json:"level"`
	Introduction string `json:"intro"`
	RegCity      string `json:"city"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

// BuildUserInfoResponse 序列化用户信息响应
func BuildUserInfoResponse(user model.User) Response {
	return Response{
		Code: CodeSuccess,
		Data: UserInfo{
			UserID:       user.UserID,
			Type:         user.Type,
			Name:         user.Name,
			IDType:       user.IDType,
			IDNumber:     user.IDNumber,
			Phone:        user.Phone,
			Level:        user.Level,
			Introduction: user.Introduction,
			RegCity:      user.RegCity,
			CreatedAt:    user.CreatedAt.Unix(),
			UpdatedAt:    user.UpdatedAt.Unix(),
		},
	}
}
