package service

import (
	"HH_ADMIN/internal/consts"
	"HH_ADMIN/internal/model/entity"
	"HH_ADMIN/utility/rr"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gvalid"
	"reflect"
	"time"
)

var Middleware sMiddleware

type sMiddleware struct{}

// Description 鉴权方法
// Author daixk
// Date 2023-09-07 19:15:49
func (s *sMiddleware) PermissionAuthentication(r *ghttp.Request, user *entity.SysUser) {
	fmt.Println("phone:" + gconv.String(user.Phone))
	fmt.Println("url:" + r.URL.Path)
	fmt.Println("method:" + r.Method)

	// 超级管理员
	if user.UserType == 1 {
		fmt.Println("-----------------------超级管理员不用鉴权-----------------------")
		r.Middleware.Next()
		return
	}

	fmt.Println("-----------------------非超级管理员需要鉴权-----------------------")
	// 鉴权方法
	enforce, err := consts.Casbin.Enforce(user.UserName, r.URL.Path, r.Method)
	if err != nil {
		rr.FailedJsonExitAll(r)
		return
	}
	if !enforce {
		rr.FailedJsonWithCodeAndMessageExitAll(r, consts.CODENOPERMISSIONS, consts.CODENOPERMISSIONSMESSAGE)
		return
	}

	// 放行
	r.Middleware.Next()
}

// Description 统一返回
// Author daixk
// Date 2023-09-07 19:27:49
func (s *sMiddleware) HandlerResponseMiddleware(r *ghttp.Request) {
	nowTime := time.Now()

	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	if r.GetError() != nil {
		// 判断是否为参数校验异常
		if _, ok := r.GetError().(gvalid.Error); ok {
			rr.FailedJsonWithMessageExitAll(r, r.GetError().Error())
			return
		}
		g.Log().Line().Errorf(gctx.New(), "HandlerResponse err:%s", r.GetError().Error())
		//rr.FailedJsonWithMessageExitAll(r, r.GetError().Error())
		rr.FailedJsonExitAll(r)
		return
	}

	if !reflect.ValueOf(r.GetHandlerResponse()).IsValid() || reflect.ValueOf(r.GetHandlerResponse()).IsNil() {
		rr.SuccessJsonExitAll(r)
	} else {
		//r.Response.WriteJsonExit(r.GetHandlerResponse())
		// 请求时间
		v := new(rr.CommonRes)
		if err := gconv.Struct(r.GetHandlerResponse(), &v); err != nil {
			r.Response.WriteJson(r.GetHandlerResponse())
			return
		}
		v.DataTime = time.Since(nowTime).Seconds()
		r.Response.WriteJson(v)
	}
}

// Description 跨域
// Author daixk
// Data 2023-08-17 22:00:21
func (s *sMiddleware) CORSMiddleware(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
