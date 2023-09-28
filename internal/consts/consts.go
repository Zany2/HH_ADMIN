package consts

import (
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/os/gtime"
)

var Casbin *casbin.Enforcer // casbin
var StartTime *gtime.Time   // 启动时间
var DelLuaSlice []interface{}

func init() {
	DelLuaSlice = append(DelLuaSlice, CACHEDEPT, CACHEPOST, CACHEEROLE, CACHEEMENU, CACHEDEPTHM, CACHEPOSTHM, CACHEROLEHM, CACHEMENUHM)
}

// Redis 缓存Key
const (
	GTOKENLOGINPREFIX = "gTokenLogin:" // gtoken登录redis前缀

	CACHEDELLUA = "HH_ADMIN:DelLua" // 删除lua脚本

	CACHEDEPT  = "HH_ADMIN:CacheDept" // 部门信息RedisKey
	CACHEPOST  = "HH_ADMIN:CachePost" // 岗位信息RedisKey
	CACHEEROLE = "HH_ADMIN:CacheRole" // 角色信息RedisKey
	CACHEEMENU = "HH_ADMIN:CacheMenu" // 菜单信息RedisKey

	CACHEDEPTHM = "HH_ADMIN:CacheDeptHM" // 部门信息RedisKeyHM
	CACHEPOSTHM = "HH_ADMIN:CachePostHM" // 岗位信息RedisKeyHM
	CACHEROLEHM = "HH_ADMIN:CacheRoleHM" // 角色信息RedisKeyHM
	CACHEMENUHM = "HH_ADMIN:CacheMenuHM" // 菜单信息RedisKeyHM

)

// 全局 上下文Key
const (
	CTXUSERID   = "CtxUserId"   // 上下文USERID
	CTXPHONE    = "CtxPhone"    // 上下文CTXPHONE
	CTXUSERNAME = "CtxUserName" // 上下文CTXUSERNAME
	CTXDEPID    = "CtxDepid"    // 上下文CTXDEPID
)

// 全局 code 状态码
const (
	CODEOK                            = 20000 // 请求成功（非必须）
	CODECODETOKENAUTHENTICATIONFAILED = 40000 // token认证过期
	CODENOPERMISSIONS                 = 40001 // 无权限
	CODESERVERERROR                   = 50000 // 服务异常
	CODEUSERNAMEORPASSWORDERROR       = 50001 // 账号或密码错误
)

// 全局 message 返回值
const (
	CODEOKMESSAGE                            = "请求成功"
	CODECODETOKENAUTHENTICATIONFAILEDMESSAGE = "身份认证过期，请重新登录"
	CODENOPERMISSIONSMESSAGE                 = "暂无此权限"
	CODESERVERBUSYMESSAGE                    = "服务异常，请稍后再试或联系管理员"
	CODEUSERNAMEORPASSWORDERRORMESSAGE       = "账号或密码错误"
	CODEUSERNOTEXISTMESSAGE                  = "登录失败，不存在该用户"
)

func init() {
	// 启动时间
	StartTime = gtime.Now()
}
