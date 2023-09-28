package rr

import (
	"HH_ADMIN/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// CommonRes 数据返回通用JSON数据结构
type CommonRes struct {
	Code     int         `json:"code"`      // 提示码
	Message  string      `json:"message"`   // 提示信息
	Data     interface{} `json:"data"`      // 返回数据(业务接口定义具体数据结构)
	DataTime float64     `json:"data_time"` // 请求时间
}

// Json 返回标准JSON数据。
func Json(r *ghttp.Request, code int, message string, data interface{}) {
	r.Response.WriteJson(CommonRes{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// SuccessJsonExitAll 成功返回
func SuccessJsonExitAll(r *ghttp.Request) {
	Json(r, consts.CODEOK, consts.CODEOKMESSAGE, g.Array{})
	r.ExitAll()
}

// SuccessJsonWithMessageExitAll 带有提示的成功返回
func SuccessJsonWithMessageExitAll(r *ghttp.Request, message string) {
	Json(r, consts.CODEOK, message, g.Array{})
	r.ExitAll()
}

// SuccessJsonWithDataExitAll 带有数据的成功返回
func SuccessJsonWithDataExitAll(r *ghttp.Request, data interface{}) {
	Json(r, consts.CODEOK, consts.CODEOKMESSAGE, data)
	r.ExitAll()
}

// SuccessJsonWithMessageAndData 带有数据和提示成功返回
func SuccessJsonWithMessageAndData(r *ghttp.Request, message string, data interface{}) {
	Json(r, consts.CODEOK, message, data)
	r.ExitAll()
}

// FailedJsonExitAll 失败返回
func FailedJsonExitAll(r *ghttp.Request) {
	Json(r, consts.CODESERVERERROR, consts.CODESERVERBUSYMESSAGE, g.Array{})
	r.ExitAll()
}

// FailedJsonWithMessageExitAll 带有提示的失败返回
func FailedJsonWithMessageExitAll(r *ghttp.Request, message string) {
	Json(r, consts.CODESERVERERROR, message, g.Array{})
	r.ExitAll()
}

// FailedJsonWithCodeAndMessageExitAll 带有提示码和提示的失败返回
func FailedJsonWithCodeAndMessageExitAll(r *ghttp.Request, code int, message string) {
	Json(r, code, message, g.Array{})
	r.ExitAll()
}

// Success 统一成功返回
func Success() *CommonRes {
	return &CommonRes{
		Code:    consts.CODEOK,
		Message: consts.CODEOKMESSAGE,
		Data:    g.Array{},
	}
}

// SuccessWithMessage 带有提示的统一成功返回
func SuccessWithMessage(message string) *CommonRes {
	return &CommonRes{
		Code:    consts.CODEOK,
		Message: message,
		Data:    g.Array{},
	}
}

// SuccessWithData 带有数据的统一成功返回
func SuccessWithData(data interface{}) *CommonRes {
	return &CommonRes{
		Code:    consts.CODEOK,
		Message: consts.CODEOKMESSAGE,
		Data:    data,
	}
}

// SuccessWithMessageAndData 带有数据和提示的统一成功返回
func SuccessWithMessageAndData(message string, data interface{}) *CommonRes {
	return &CommonRes{
		Code:    consts.CODEOK,
		Message: message,
		Data:    data,
	}
}

// Failed 统一失败返回
func Failed() *CommonRes {
	return &CommonRes{
		Code:    consts.CODESERVERERROR,
		Message: consts.CODESERVERBUSYMESSAGE,
		Data:    g.Array{},
	}
}

// FailedWithMessage 带有提示的统一失败返回
func FailedWithMessage(message string) *CommonRes {
	return &CommonRes{
		Code:    consts.CODESERVERERROR,
		Message: message,
		Data:    g.Array{},
	}
}

// FailedWithCodeAndMessage 带有提示码和提示的统一失败返回
func FailedWithCodeAndMessage(code int, message string) *CommonRes {
	return &CommonRes{
		Code:    code,
		Message: message,
		Data:    g.Array{},
	}
}
