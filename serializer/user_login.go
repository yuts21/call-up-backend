package serializer

import "time"

// UserLogin 用户登录序列化器
type UserLogin struct {
	ID     uint   `json:"UID"`
	Type   bool   `json:"type"`
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

// BuildUserLoginResponse 序列化用户登录响应
func BuildUserLoginResponse(ID uint, Type bool, Token string, Expire time.Time) Response {
	return Response{
		Code: CodeSuccess,
		Data: UserLogin{
			ID:     ID,
			Type:   Type,
			Token:  Token,
			Expire: Expire.Unix(),
		},
		Msg: "登录成功",
	}
}
