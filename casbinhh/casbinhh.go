package casbinhh

import (
	"HH_ADMIN/internal/consts"
	"HH_ADMIN/internal/dao"
	"HH_ADMIN/internal/model"
	"HH_ADMIN/internal/model/entity"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/frame/g"
	adapter "github.com/hailaz/gf-casbin-adapter/v2"
	"strings"
)

// Description 初始化casbin
// Author daixk
// Data 2023-09-06 23:04:39
func InitCasbin(ctx context.Context) {
	a := adapter.NewAdapter(adapter.Options{GDB: g.DB()})
	enforcer, err := casbin.NewEnforcer(g.Cfg().MustGet(ctx, "casbin.model_path").String(), a)
	if err != nil {
		panic(err)
	}

	// 初始化加载用户角色到权限中
	// 角色 list
	sysRoles := make([]entity.SysRole, 0)
	err = dao.SysRole.Ctx(ctx).Scan(&sysRoles)
	if err != nil {
		panic(err)
	}

	// 角色 菜单 list
	menus := make([]entity.SysRoleMenu, 0)
	err = dao.SysRoleMenu.Ctx(ctx).Scan(&menus)
	if err != nil {
		panic(err)
	}

	// 菜单id list
	ints := make([]int64, 0)
	for _, menu := range menus {
		ints = append(ints, menu.MenuId)
	}

	// 菜单 list
	sysMenus := make([]entity.SysMenu, 0)
	err = dao.SysMenu.Ctx(ctx).Scan(&sysMenus)
	if err != nil {
		panic(err)
	}

	methods := make([]model.RoleMenuAccessMethod, 0)
	for _, menu := range menus {
		for _, role := range sysRoles {
			if menu.RoleId == role.RoleId {
				for _, sysMenu := range sysMenus {
					if sysMenu.MenuId == menu.MenuId {
						split := strings.Split(sysMenu.Perms, ":")
						methods = append(methods, model.RoleMenuAccessMethod{
							RoleName:  role.RoleKey,
							MenuPerms: split[0],
							Method:    split[1],
						})
					}
				}
			}
		}
	}

	// 设置权限
	for _, method := range methods {
		_, err = enforcer.AddPermissionForUser(method.RoleName, method.MenuPerms, method.Method)
		if err != nil {
			panic(err)
		}
	}

	consts.Casbin = enforcer
}

// Description 根据userName添加casbin角色
// Author daixk
// Date 2023-09-21 08:54:22
func CasbinAddRoleForUser(userName string, sysRoles []entity.SysRole) (err error) {
	// 添加 casbin 角色
	for _, role := range sysRoles {
		_, err = consts.Casbin.AddRoleForUser(userName, role.RoleKey)
		if err != nil {
			return err
		}
	}
	return
}

// Description 根据userName删除casbin角色
// Author daixk
// Date 2023-09-21 09:14:54
func CasbinDeleteRoleForUser(userName string, sysRoles []entity.SysRole) (err error) {
	for _, sysRole := range sysRoles {
		_, err = consts.Casbin.DeleteRoleForUser(userName, sysRole.RoleKey)
		if err != nil {
			return err
		}
	}
	return
}
