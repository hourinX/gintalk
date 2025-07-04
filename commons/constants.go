package commons

const (
	CodeSuccess        = 10000 // 操作成功
	CodeFail           = 10001 // 操作失败
	CodeParamError     = 10002 // 参数错误
	CodeUnauthorized   = 10003 // 未登录/未授权
	CodeForbidden      = 10004 // 没有权限
	CodeNotFound       = 10005 // 资源不存在
	CodeServerError    = 10006 // 服务端异常
	CodeConflict       = 10007 // 资源冲突
	CodeTooManyRequest = 10008 // 请求过于频繁
)

var CodeMap = map[string]int{
	"success":      CodeSuccess,
	"fail":         CodeFail,
	"param_error":  CodeParamError,
	"unauthorized": CodeUnauthorized,
	"forbidden":    CodeForbidden,
	"not_found":    CodeNotFound,
	"server_error": CodeServerError,
	"conflict":     CodeConflict,
	"too_many_req": CodeTooManyRequest,
}
