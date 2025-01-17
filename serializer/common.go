package serializer

import "github.com/gin-gonic/gin"

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"res,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"err,omitempty"`
}

// DataList 基础列表结构
type DataList struct {
	Items interface{} `json:"items"`
	Total int64       `json:"total"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// CodeSuccess 成功
	CodeSuccess = 200
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	// CodeCacheError 缓存操作失败
	CodeCacheError = 50003
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
	// CodeFileUploadError 文件上传错误
	CodeFileUploadError = 40002
)

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Code: errCode,
		Msg:  msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// Success 通用成功处理
func Success(msg string) Response {
	res := Response{
		Code: CodeSuccess,
		Msg:  msg,
	}
	return res
}

// BuildListResponse 序列列表响应
func BuildListResponse(items interface{}, total int64) Response {
	return Response{
		Code: CodeSuccess,
		Data: DataList{
			Items: items,
			Total: total,
		},
	}
}
