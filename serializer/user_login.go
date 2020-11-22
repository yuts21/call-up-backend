package serializer

import "time"

// UserLogin 用户登录序列化器
type UserLogin struct {
	ID    uint `json:"UID"`
	Token  string `json:"token"`
	Expire int64 `json:"expire"`
}

// BuildUserLoginResponse 序列化用户响应
func BuildUserLoginResponse(ID uint, token string, expire time.Time) Response {
	return Response{
		Code: CodeSuccess,
		Data: UserLogin{
			ID: ID,
			Token: token,
			Expire: expire.Unix(),
		},
	}
}
