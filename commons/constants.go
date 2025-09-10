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

const (
	YYYY           = "2006"
	YYYYMM         = "2006-01"
	YYYYMMDD       = "2006-01-02"
	YYYYMMDDHH     = "2006-01-02 15"
	YYYYMMDDHHMM   = "2006-01-02 15:04"
	YYYYMMDDHHMMSS = "2006-01-02 15:04:05"
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

const (
	MaxIDLength = 20 // ID最大长度
)

const (
	StatusUp = 1 // 状态正常
)

const (
	GroupTypeSystem  = 0 //系统分组
	GroupTypeDefault = 1 //默认分组
	GroupTypeCustom  = 2 //自定义分组
)

const (
	EditableNotPermission = 0 //不可修改
	EditablePermission    = 1 //可修改
)
