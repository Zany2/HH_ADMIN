package cmd

import (
	"HH_ADMIN/internal/consts"
	"HH_ADMIN/internal/dao"
	"HH_ADMIN/internal/model/entity"
	"HH_ADMIN/internal/service"
	"HH_ADMIN/util"
	"HH_ADMIN/utility/rr"
	"context"
	"database/sql"
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// Description gtoken 配置启动
// Author daixk
// Date 2023-09-07 08:45:55
func StartBackendGToken(ctx context.Context) *gtoken.GfToken {
	gToken := &gtoken.GfToken{
		CacheMode:        2,
		ServerName:       "HH_ADMIN",
		LoginPath:        "/login",
		LoginBeforeFunc:  loginFunc,
		LoginAfterFunc:   loginAfterFunc,
		LogoutPath:       "/logout",
		LogoutBeforeFunc: logoutFunc,
		LogoutAfterFunc:  logoutAfterFunc,
		AuthPaths:        g.SliceStr{},
		AuthExcludePaths: g.SliceStr{"/api/v1/hello"},
		AuthAfterFunc:    authAfterFunc,
		MultiLogin:       true,
		Timeout:          10000000,
	}
	if err := gToken.Start(); err != nil {
		panic(err)
	}
	return gToken
}

// Description 自定义登录方法
// Author daixk
// Data 2023-09-05 15:02:37
func loginFunc(r *ghttp.Request) (string, interface{}) {
	userName := r.GetForm("user_name").String()
	password := r.GetForm("password").String()

	fmt.Println(userName)
	fmt.Println(password)

	if userName == "" || password == "" {
		dao.SysLoginLog.Ctx(gctx.New()).Data(entity.SysLoginLog{
			LoginName:     userName,
			Ipaddr:        util.GetRequestIp(r.Request),
			Browser:       r.Header.Get("User-Agent"),
			LoginLocation: "内网IP",
			Os:            r.Header.Get("Sec-Ch-Ua-Platform"),
			Status:        2,
			Msg:           consts.CODEUSERNAMEORPASSWORDERRORMESSAGE,
			LoginTime:     gtime.Now(),
			Module:        "系统后台",
		}).OmitEmpty().Insert()
		rr.FailedJsonWithMessageExitAll(r, consts.CODEUSERNAMEORPASSWORDERRORMESSAGE)
		return "", nil
	}

	user := new(entity.SysUser)
	if err := dao.SysUser.Ctx(gctx.New()).Where(dao.SysUser.Columns().UserName, userName).Where(dao.SysUser.Columns().DelFlag, 1).Scan(&user); err != nil {
		if err == sql.ErrNoRows {
			dao.SysLoginLog.Ctx(gctx.New()).Data(entity.SysLoginLog{
				LoginName:     userName,
				Ipaddr:        util.GetRequestIp(r.Request),
				Browser:       r.Header.Get("User-Agent"),
				LoginLocation: "内网IP",
				Os:            r.Header.Get("Sec-Ch-Ua-Platform"),
				Status:        2,
				Msg:           consts.CODEUSERNOTEXISTMESSAGE,
				LoginTime:     gtime.Now(),
				Module:        "系统后台",
			}).OmitEmpty().Insert()
			rr.FailedJsonWithMessageExitAll(r, consts.CODEUSERNOTEXISTMESSAGE)
			return "", nil
		} else {
			dao.SysLoginLog.Ctx(gctx.New()).Data(entity.SysLoginLog{
				LoginName:     userName,
				Ipaddr:        util.GetRequestIp(r.Request),
				Browser:       r.Header.Get("User-Agent"),
				LoginLocation: "内网IP",
				Os:            r.Header.Get("Sec-Ch-Ua-Platform"),
				Status:        2,
				Msg:           consts.CODESERVERBUSYMESSAGE,
				LoginTime:     gtime.Now(),
				Module:        "系统后台",
			}).OmitEmpty().Insert()
			rr.FailedJsonExitAll(r)
			return "", nil
		}
	}

	encrypt, err := gmd5.EncryptString(password)
	if err != nil {
		rr.FailedJsonExitAll(r)
		return "", nil
	}

	if user.Password != encrypt {
		dao.SysLoginLog.Ctx(gctx.New()).Data(entity.SysLoginLog{
			LoginName:     userName,
			Ipaddr:        util.GetRequestIp(r.Request),
			Browser:       r.Header.Get("User-Agent"),
			LoginLocation: "内网IP",
			Os:            r.Header.Get("Sec-Ch-Ua-Platform"),
			Status:        2,
			Msg:           consts.CODEUSERNAMEORPASSWORDERRORMESSAGE,
			LoginTime:     gtime.Now(),
			Module:        "系统后台",
		}).OmitEmpty().Insert()
		rr.FailedJsonWithMessageExitAll(r, consts.CODEUSERNAMEORPASSWORDERRORMESSAGE)
		return "", nil
	}

	if user.Status == 2 {
		dao.SysLoginLog.Ctx(gctx.New()).Data(entity.SysLoginLog{
			LoginName:     userName,
			Ipaddr:        util.GetRequestIp(r.Request),
			Browser:       r.Header.Get("User-Agent"),
			LoginLocation: "内网IP",
			Os:            r.Header.Get("Sec-Ch-Ua-Platform"),
			Status:        2,
			Msg:           "账号已停用",
			LoginTime:     gtime.Now(),
			Module:        "系统后台",
		}).OmitEmpty().Insert()
		rr.FailedJsonWithMessageExitAll(r, "账号已停用")
		return "", nil
	}

	dao.SysLoginLog.Ctx(gctx.New()).Data(entity.SysLoginLog{
		LoginName:     userName,
		Ipaddr:        util.GetRequestIp(r.Request),
		Browser:       r.Header.Get("User-Agent"),
		LoginLocation: "内网IP",
		Os:            r.Header.Get("Sec-Ch-Ua-Platform"),
		Status:        1,
		Msg:           "登录成功",
		LoginTime:     gtime.Now(),
		Module:        "系统后台",
	}).OmitEmpty().Insert()

	return consts.GTOKENLOGINPREFIX + gconv.String(user.UserId), user
}

// Description 自定义登录后方法
// Author daixk
// Data 2023-09-05 15:00:07
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Success() {
		userKey := respData.GetString(gtoken.KeyUserKey)
		userId := gstr.StrEx(userKey, consts.GTOKENLOGINPREFIX)

		// 这里查询角色信息与权限等信息
		user := new(entity.SysUser)
		err := dao.SysUser.Ctx(gctx.New()).Where(dao.SysUser.Columns().UserId, userId).Where(dao.SysUser.Columns().DelFlag, 1).Scan(&user)
		if err != nil {
			g.Log().Line().Errorf(gctx.New(), "loginAfterFunc err:%s", err.Error())
			rr.FailedJsonWithMessageExitAll(r, consts.CODESERVERBUSYMESSAGE)
			return
		}

		user.Password = ""
		rr.SuccessJsonWithDataExitAll(r, map[string]interface{}{
			"user_data": user,
			"token":     respData.GetString(gtoken.KeyToken),
		})
		return
	}
	rr.FailedJsonWithMessageExitAll(r, consts.CODESERVERBUSYMESSAGE)
	return
}

// Description 鉴权中间件
// Author daixk
// Data 2023-09-05 15:02:57
func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	fmt.Println("-----------------------走鉴权-----------------------")

	if !respData.Success() {
		rr.FailedJsonWithCodeAndMessageExitAll(r, consts.CODECODETOKENAUTHENTICATIONFAILED, consts.CODECODETOKENAUTHENTICATIONFAILEDMESSAGE)
		return
	}

	user := new(entity.SysUser)
	if err := gjson.Unmarshal([]byte(respData.GetString(gtoken.KeyData)), &user); err != nil {
		rr.FailedJsonWithCodeAndMessageExitAll(r, consts.CODECODETOKENAUTHENTICATIONFAILED, consts.CODECODETOKENAUTHENTICATIONFAILEDMESSAGE)
		return
	}

	// 上下文数据
	r.SetCtxVar(consts.CTXUSERID, user.UserId)
	r.SetCtxVar(consts.CTXUSERNAME, user.UserName)
	r.SetCtxVar(consts.CTXPHONE, user.Phone)

	// 权限认证
	service.Middleware.PermissionAuthentication(r, user)
}

// Description 自定义注销方法
// Author daixk
// Date 2023-09-07 19:11:06
func logoutFunc(r *ghttp.Request) bool {
	return true
}

// Description 自定义注销后方法
// Author daixk
// Date 2023-09-07 19:11:34
func logoutAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	rr.SuccessJsonWithMessageExitAll(r, "注销成功")
	return
}
