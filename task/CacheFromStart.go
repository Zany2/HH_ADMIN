package task

import (
	"HH_ADMIN/internal/consts"
	"HH_ADMIN/internal/dao"
	"HH_ADMIN/internal/model/entity"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func InitLua(ctx context.Context) {
	// 定义 Lua 脚本
	luaScript := `
		local keys = KEYS
		for i, key in ipairs(keys) do
			redis.call('DEL', key)
		end
		return keys
		`
	_, err := g.Redis().Set(ctx, consts.CACHEDELLUA, gconv.String(luaScript))
	if err != nil {
		panic(err)
	}
}

// Description 缓存部门、岗位、角色信息
// Author daixk
// Date 2023-09-11 09:19:44
func CacheRPD(ctx context.Context) (err error) {

	if err = cacheRole(ctx); err != nil {
		g.Log().Line().Errorf(ctx, "CacheRPD err:%s", err.Error())
		return
	}
	if err = cachePosts(ctx); err != nil {
		g.Log().Line().Errorf(ctx, "CacheRPD err:%s", err.Error())
		return
	}
	if err = cacheDept(ctx); err != nil {
		g.Log().Line().Errorf(ctx, "CacheRPD err:%s", err.Error())
		return
	}
	if err = cacheMenu(ctx); err != nil {
		g.Log().Line().Errorf(ctx, "CacheRPD err:%s", err.Error())
		return
	}

	return
}

// Description 刷新部门、岗位、角色信息缓存
// Author daixk
// Data 2023-09-20 22:46:55
func CacheRefreshRPD(ctx context.Context) (err error) {

	lua, err := g.Redis().Get(ctx, consts.CACHEDELLUA)
	if err != nil {
		return err
	}
	if lua.IsEmpty() {
		return gerror.New("lua 脚本为空")
	}

	_, err = g.Redis().Eval(ctx, lua.String(), gconv.Int64(len(consts.DelLuaSlice)), []string{}, consts.DelLuaSlice)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//result, err := g.Redis().Do(ctx, "EVAL", lua, len(consts.DelLuaSlice), consts.DelLuaSlice...)
	//Strings(conn.Do("EVAL", luaScript, len(keys), redis.Args{}.AddFlat(keys)...))
	//
	//g.Redis().Eval(ctx, lua.String(), len(consts.DelLuaSlice), consts.DelLuaSlice)

	//_, err = g.Redis().Del(ctx, consts.CACHEEROLE)
	//if err != nil {
	//	return err
	//}
	//_, err = g.Redis().Del(ctx, consts.CACHEROLEHM)
	//if err != nil {
	//	return err
	//}
	//_, err = g.Redis().Del(ctx, consts.CACHEPOST)
	//if err != nil {
	//	return err
	//}
	//_, err = g.Redis().Del(ctx, consts.CACHEPOSTHM)
	//if err != nil {
	//	return err
	//}
	//_, err = g.Redis().Del(ctx, consts.CACHEDEPT)
	//if err != nil {
	//	return err
	//}
	//_, err = g.Redis().Del(ctx, consts.CACHEDEPTHM)
	//if err != nil {
	//	return err
	//}
	//_, err = g.Redis().Del(ctx, consts.CACHEEMENU)
	//if err != nil {
	//	return err
	//}
	//_, err = g.Redis().Del(ctx, consts.CACHEMENUHM)
	//if err != nil {
	//	return err
	//}
	err = CacheRPD(ctx)
	return err
}

// CacheRole 缓存角色信息
func cacheRole(ctx context.Context) (err error) {
	// 缓存角色
	roles := make([]entity.SysRole, 0)
	err = dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().Status, 1).Where(dao.SysRole.Columns().DelFlag, 1).Scan(&roles)
	if err != nil {
		g.Log().Line().Errorf(ctx, "角色缓存查询失败,err:%s", err.Error())
		return err
	}
	rolesJson := make(map[string]interface{}, 0)
	for _, role := range roles {
		marshal, err := json.Marshal(role)
		if err != nil {
			g.Log().Line().Errorf(ctx, "角色缓存序列化失败,err:%s", err.Error())
			return err
		}
		rolesJson[gconv.String(role.RoleId)] = string(marshal)
	}
	_, err = g.Redis().Set(ctx, consts.CACHEEROLE, roles)
	if err != nil {
		return err
	}
	err = g.Redis().HMSet(ctx, consts.CACHEROLEHM, rolesJson)
	if err != nil {
		return err
	}
	g.Log().Line().Info(ctx, "角色信息缓存成功...")
	return
}

// CachePosts 缓存岗位信息
func cachePosts(ctx context.Context) (err error) {
	// 缓存岗位
	posts := make([]entity.SysPost, 0)
	err = dao.SysPost.Ctx(ctx).Where(dao.SysPost.Columns().Status, 1).Where(dao.SysPost.Columns().DelFlag, 1).Scan(&posts)
	if err != nil {
		g.Log().Line().Errorf(ctx, "岗位缓存查询失败,err:%s", err.Error())
		return err
	}
	postsJson := make(map[string]interface{}, 0)
	for _, post := range posts {
		marshal, err := json.Marshal(post)
		if err != nil {
			g.Log().Line().Errorf(ctx, "岗位缓存序列化失败,err:%s", err.Error())
			return err
		}
		postsJson[gconv.String(post.PostId)] = string(marshal)
	}
	_, err = g.Redis().Set(ctx, consts.CACHEPOST, posts)
	if err != nil {
		return err
	}
	err = g.Redis().HMSet(ctx, consts.CACHEPOSTHM, postsJson)
	if err != nil {
		return err
	}
	g.Log().Line().Info(ctx, "岗位信息缓存成功...")
	return
}

// CacheDept 缓存部门信息
func cacheDept(ctx context.Context) (err error) {
	// 缓存部门
	depts := make([]entity.SysDept, 0)
	err = dao.SysDept.Ctx(ctx).Where(dao.SysDept.Columns().Status, 1).Where(dao.SysDept.Columns().DelFlag, 1).Scan(&depts)
	if err != nil {
		g.Log().Line().Errorf(ctx, "部门缓存查询失败,err:%s", err.Error())
		return err
	}
	deptsJson := make(map[string]interface{}, 0)
	for _, dept := range depts {
		marshal, err := json.Marshal(dept)
		if err != nil {
			g.Log().Line().Errorf(ctx, "部门缓存序列化失败,err:%s", err.Error())
			return err
		}
		deptsJson[gconv.String(dept.DeptId)] = string(marshal)
	}
	_, err = g.Redis().Set(ctx, consts.CACHEDEPT, depts)
	if err != nil {
		return err
	}
	err = g.Redis().HMSet(ctx, consts.CACHEDEPTHM, deptsJson)
	if err != nil {
		return err
	}
	g.Log().Line().Info(ctx, "部门信息缓存成功...")
	return
}

// cacheMenu 缓存菜单信息
func cacheMenu(ctx context.Context) (err error) {
	// 缓存菜单
	menus := make([]entity.SysMenu, 0)
	err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Status, 1).Where(dao.SysMenu.Columns().Visible, 1).Scan(&menus)
	if err != nil {
		g.Log().Line().Errorf(ctx, "菜单缓存查询失败,err:%s", err.Error())
		return err
	}
	menusJson := make(map[string]interface{}, 0)
	for _, menu := range menus {
		marshal, err := json.Marshal(menu)
		if err != nil {
			g.Log().Line().Errorf(ctx, "部门缓存序列化失败,err:%s", err.Error())
			return err
		}
		menusJson[gconv.String(menu.MenuId)] = string(marshal)
	}
	_, err = g.Redis().Set(ctx, consts.CACHEEMENU, menus)
	if err != nil {
		return err
	}
	err = g.Redis().HMSet(ctx, consts.CACHEMENUHM, menusJson)
	if err != nil {
		return err
	}
	g.Log().Line().Info(ctx, "菜单信息缓存成功...")
	return

}
