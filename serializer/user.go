package serializer

import "call-up/model"

// User 用户序列化器
type User struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		UserID: user.UserID,
		//Username:  admin.Username,
		//CreatedAt: admin.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
