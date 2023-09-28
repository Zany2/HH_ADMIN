package service

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/internal/consts"
	"HH_ADMIN/internal/dao"
	"HH_ADMIN/internal/model"
	"HH_ADMIN/internal/model/entity"
	"HH_ADMIN/task"
	"HH_ADMIN/util"
	"HH_ADMIN/utility/rr"
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

type sRole struct{}

func Role() *sRole {
	return &sRole{}
}

// RoleList 角色列表
func (s *sRole) RoleList(ctx context.Context, req *api.RoleListReq) (res *rr.CommonRes, err error) {
	m := g.Model(dao.SysRole.Table())
	m.Where(dao.SysRole.Columns().DelFlag, 1)
	if req.RoleName != "" {
		m.Where(dao.SysRole.Columns().RoleName+" like ?", "%"+req.RoleName+"%")
	}
	if req.Status != 0 {
		m.Where(dao.SysRole.Columns().Status, req.Status)
	}

	var total int
	roles := make([]entity.SysRole, 0)
	err = m.Page(req.Pn, req.PageSize).OrderAsc(dao.SysRole.Columns().RoleSort).ScanAndCount(&roles, &total, false)
	if err != nil {
		g.Log().Line().Errorf(ctx, "RoleList err:%s", err.Error())
		return rr.Failed(), err
	}
	if total <= 0 {
		return rr.SuccessWithData(map[string]interface{}{
			"total":     0,
			"role_list": g.Array{},
		}), err
	}

	// 角色返回列表
	roleListRes := make([]api.RoleListRes, 0)

	for _, role := range roles {
		menusAll := make([]*entity.SysMenu, 0)

		if role.RoleKey == "admin" {
			// 超级管理员
			get, err := g.Redis().Get(ctx, consts.CACHEEMENU)
			if err != nil {
				g.Log().Line().Errorf(ctx, "RoleList err:%s", err.Error())
				return rr.Failed(), err
			}
			err = gconv.Struct(get, &menusAll)
			if err != nil {
				g.Log().Line().Errorf(ctx, "RoleList err:%s", err.Error())
				return rr.Failed(), err
			}

		} else {
			all, err := dao.SysRoleMenu.Ctx(ctx).Fields(dao.SysRoleMenu.Columns().MenuId).Where(dao.SysRoleMenu.Columns().RoleId, role.RoleId).All()
			if err != nil {
				g.Log().Line().Errorf(ctx, "RoleList err:%s", err.Error())
				return rr.Failed(), err
			}

			array := all.Array(dao.SysRoleMenu.Columns().MenuId)
			for _, value := range array {
				menu := new(entity.SysMenu)
				menuRedisCache, err := g.Redis().HMGet(ctx, consts.CACHEMENUHM, gconv.String(value))
				if err != nil {
					g.Log().Line().Errorf(ctx, "RoleList err:%s", err.Error())
					return rr.Failed(), err
				}
				err = gconv.Struct(menuRedisCache[0], &menu)
				if err != nil {
					g.Log().Line().Errorf(ctx, "RoleList err:%s", err.Error())
					return rr.Failed(), err
				}
				menusAll = append(menusAll, menu)
			}
		}

		menuS := make([]*model.SysMenuS, 0)

		for _, menuAllOne := range menusAll {
			menuSOne := new(model.SysMenuS)
			err := gconv.Struct(menuAllOne, &menuSOne)
			if err != nil {
				g.Log().Line().Errorf(ctx, "RoleList err:%s", err.Error())
				return rr.Failed(), err
			}
			menuS = append(menuS, menuSOne)
		}

		menuTree := util.GetTreeRecursiveMenu(menuS, menusAll[0].ParentId)

		roleListRes = append(roleListRes, api.RoleListRes{
			Role:     role,
			MenuList: menuTree,
		})
	}

	return rr.SuccessWithData(map[string]interface{}{
		"total":     total,
		"role_list": roleListRes,
	}), err
}

// RoleAdd 角色添加
func (s *sRole) RoleAdd(ctx context.Context, req *api.RoleAddReq) (res *rr.CommonRes, err error) {
	//userId := ctx.Value(consts.CTXUSERID)
	//phone := ctx.Value(consts.CTXPHONE)
	userName := ctx.Value(consts.CTXUSERNAME)

	count, err := dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().RoleKey, req.RoleKey).Count()
	if err != nil {
		g.Log().Line().Errorf(ctx, "RoleAdd err:%s", err.Error())
		return rr.FailedWithMessage("添加失败"), err
	}
	if count > 0 {
		return rr.FailedWithMessage("角色权限为" + req.RoleKey + "的角色已存在"), err
	}

	menus := make([]entity.SysMenu, 0)
	var menu entity.SysMenu
	if len(req.MenuIds) > 0 {
		for _, id := range req.MenuIds {
			get, err := g.Redis().HMGet(ctx, consts.CACHEMENUHM, gconv.String(id))
			if err != nil {
				g.Log().Line().Errorf(ctx, "RoleAdd err:%s", err.Error())
				return rr.FailedWithMessage("添加失败"), err
			}
			err = gconv.Struct(get[0], &menu)
			if err != nil {
				g.Log().Line().Errorf(ctx, "RoleAdd err:%s", err.Error())
				return rr.FailedWithMessage("添加失败"), err
			}
			menus = append(menus, menu)
		}
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 添加角色
			id, err := dao.SysRole.Ctx(ctx).TX(tx).InsertAndGetId(entity.SysRole{
				RoleName:          req.RoleName,
				RoleKey:           req.RoleKey,
				RoleSort:          req.RoleSort,
				DataScope:         2,
				MenuCheckStrictly: 1,
				DeptCheckStrictly: 1,
				Status:            1,
				DelFlag:           1,
				CreateBy:          gconv.String(userName),
				Remark:            req.Remark,
				CreateTime:        gtime.Now(),
			})
			util.ErrIsNil(ctx, err, "添加角色失败")

			// 添加 角色和菜单关联
			if len(req.MenuIds) > 0 {
				menus := make([]entity.SysRoleMenu, 0)
				for _, menuId := range req.MenuIds {
					menus = append(menus, entity.SysRoleMenu{
						RoleId: id,
						MenuId: gconv.Int64(menuId),
					})
				}
				_, err := dao.SysRoleMenu.Ctx(ctx).TX(tx).Data(menus).Insert()
				util.ErrIsNil(ctx, err, "添加角色和菜单关联失败")
			}

			// casbin添加角色权限
			if len(req.MenuIds) > 0 {
				for _, menu := range menus {
					split := strings.Split(menu.Perms, ":")
					_, err = consts.Casbin.AddPermissionForUser(req.RoleKey, split[0], split[1])
					util.ErrIsNil(ctx, err, "设置用户角色失败")
				}
			}

			// 刷新缓存
			err = task.CacheRefreshRPD(ctx)
			util.ErrIsNil(ctx, err, "添加失败")

		})
		return err
	}); err != nil {
		g.Log().Line().Errorf(ctx, "RoleAdd err:%s", err.Error())
		return rr.Failed(), err
	}
	return rr.SuccessWithMessage("新增成功"), err
}

// RoleUpdate 角色修改
func (s *sRole) RoleUpdate(ctx context.Context, req *api.RoleUpdateReq) (res *rr.CommonRes, err error) {
	//userId := ctx.Value(consts.CTXUSERID)
	//phone := ctx.Value(consts.CTXPHONE)
	userName := ctx.Value(consts.CTXUSERNAME)

	roleKey := new(entity.SysRole)
	err = dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().RoleKey, req.RoleKey).Limit(1).Scan(&roleKey)
	if err != nil && err != sql.ErrNoRows {
		g.Log().Line().Errorf(ctx, "RoleUpdate err:%s", err.Error())
		return rr.FailedWithMessage("修改失败"), err
	}
	if err != sql.ErrNoRows && roleKey.RoleId != req.RoleId {
		return rr.FailedWithMessage("角色权限为" + req.RoleKey + "的角色已存在"), err
	}

	// 原role
	var roleOld entity.SysRole
	oldRedis, err := g.Redis().HMGet(ctx, consts.CACHEROLEHM, gconv.String(req.RoleId))
	err = gconv.Struct(oldRedis[0], &roleOld)
	if err != nil {
		g.Log().Line().Errorf(ctx, "RoleUpdate err:%s", err.Error())
		return rr.FailedWithMessage("修改失败"), err
	}

	// 原role与menu关联
	SysRoleMenuOld := make([]entity.SysRoleMenu, 0)
	err = dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().RoleId, req.RoleId).Scan(&SysRoleMenuOld)
	if err != nil {
		g.Log().Line().Errorf(ctx, "RoleUpdate err:%s", err.Error())
		return rr.FailedWithMessage("修改失败"), err
	}

	// 原来menu
	menusOldList := make([]entity.SysMenu, 0)
	for _, SysRoleMenuOldOne := range SysRoleMenuOld {
		var menuOld entity.SysMenu
		get, err := g.Redis().HMGet(ctx, consts.CACHEMENUHM, gconv.String(SysRoleMenuOldOne.MenuId))
		if err != nil {
			g.Log().Line().Errorf(ctx, "RoleUpdate err:%s", err.Error())
			return rr.FailedWithMessage("修改失败"), err
		}
		err = gconv.Struct(get[0], &menuOld)
		if err != nil {
			g.Log().Line().Errorf(ctx, "RoleUpdate err:%s", err.Error())
			return rr.FailedWithMessage("修改失败"), err
		}
		menusOldList = append(menusOldList, menuOld)
	}

	// 新menu
	menus := make([]entity.SysMenu, 0)
	if len(req.MenuIds) > 0 {
		var menu entity.SysMenu
		for _, id := range req.MenuIds {
			get, err := g.Redis().HMGet(ctx, consts.CACHEMENUHM, gconv.String(id))
			if err != nil {
				g.Log().Line().Errorf(ctx, "RoleUpdate err:%s", err.Error())
				return rr.FailedWithMessage("修改失败"), err
			}
			err = gconv.Struct(get[0], &menu)
			if err != nil {
				g.Log().Line().Errorf(ctx, "RoleUpdate err:%s", err.Error())
				return rr.FailedWithMessage("修改失败"), err
			}
			menus = append(menus, menu)
		}
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			role := entity.SysRole{
				RoleName:   req.RoleName,
				RoleKey:    req.RoleKey,
				RoleSort:   req.RoleSort,
				DataScope:  2,
				Status:     req.Status,
				UpdateBy:   gconv.String(userName),
				Remark:     req.Remark,
				UpdateTime: gtime.Now(),
			}
			_, err := dao.SysRole.Ctx(ctx).TX(tx).OmitEmpty().Data(role).Where(dao.SysRole.Columns().RoleId, req.RoleId).Update()
			util.ErrIsNil(ctx, err, "修改失败")

			// 角色菜单关联删除
			_, err = dao.SysRoleMenu.Ctx(ctx).TX(tx).Where(dao.SysRoleMenu.Columns().RoleId, req.RoleId).Delete()
			util.ErrIsNil(ctx, err, "修改失败")

			// 删除casbin权限
			if len(menusOldList) > 0 {
				for _, menu := range menusOldList {
					split := strings.Split(menu.Perms, ":")
					_, err := consts.Casbin.DeletePermissionForUser(roleOld.RoleKey, split[0], split[1])
					util.ErrIsNil(ctx, err, "修改失败")
				}
			}

			if len(req.MenuIds) > 0 {
				menusNew := make([]entity.SysRoleMenu, 0)
				for _, menuId := range req.MenuIds {
					menusNew = append(menusNew, entity.SysRoleMenu{
						RoleId: req.RoleId,
						MenuId: gconv.Int64(menuId),
					})
				}
				_, err := dao.SysRoleMenu.Ctx(ctx).TX(tx).Data(menusNew).Insert()
				util.ErrIsNil(ctx, err, "修改失败")

				// 添加新的
				for _, menu := range menus {
					split := strings.Split(menu.Perms, ":")
					_, err := consts.Casbin.AddPermissionForUser(req.RoleKey, split[0], split[1])
					util.ErrIsNil(ctx, err, "修改失败")
				}

			}

			// 缓存
			err = task.CacheRefreshRPD(ctx)
			util.ErrIsNil(ctx, err, "修改失败")
		})
		return err
	}); err != nil {
		g.Log().Line().Errorf(ctx, "RoleUpdate err:%s", err.Error())
		return rr.FailedWithMessage("修改失败"), err
	}
	return rr.SuccessWithMessage("修改成功"), err
}

// RoleDelete 角色删除
func (s *sRole) RoleDelete(ctx context.Context, req *api.RoleDeleteReq) (res *rr.CommonRes, err error) {

	menus := make([]entity.SysRoleMenu, 0)
	err = dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().RoleId, req.RoleId).Scan(&menus)
	if err != nil {
		g.Log().Line().Errorf(ctx, "RoleDelete err:%s", err.Error())
		return rr.FailedWithMessage("删除失败"), err
	}

	// 原来menu
	menusOldList := make([]entity.SysMenu, 0)
	for _, id := range menus {
		var menuOld entity.SysMenu
		get, err := g.Redis().HMGet(ctx, consts.CACHEMENUHM, gconv.String(id))
		if err != nil {
			g.Log().Line().Errorf(ctx, "RoleDelete err:%s", err.Error())
			return rr.FailedWithMessage("删除失败"), err
		}
		err = gconv.Struct(get[0], &menuOld)
		if err != nil {
			g.Log().Line().Errorf(ctx, "RoleDelete err:%s", err.Error())
			return rr.FailedWithMessage("删除失败"), err
		}
		menusOldList = append(menusOldList, menuOld)
	}

	// 原来的role
	var roleOld entity.SysRole
	roleRedis, err := g.Redis().HMGet(ctx, consts.CACHEROLEHM, gconv.String(req.RoleId))
	if err != nil {
		g.Log().Line().Errorf(ctx, "RoleDelete err:%s", err.Error())
		return rr.FailedWithMessage("删除失败"), err
	}
	err = gconv.Struct(roleRedis[0], &roleOld)
	if err != nil {
		g.Log().Line().Errorf(ctx, "RoleDelete err:%s", err.Error())
		return rr.FailedWithMessage("删除失败"), err
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err := dao.SysRole.Ctx(ctx).TX(tx).Data(g.Map{dao.SysRole.Columns().DelFlag: 2}).Where(dao.SysRole.Columns().RoleId, req.RoleId).Update()
			util.ErrIsNil(ctx, err, "删除失败")

			_, err = dao.SysUserRole.Ctx(ctx).TX(tx).Where(dao.SysUserRole.Columns().RoleId, req.RoleId).Delete()
			util.ErrIsNil(ctx, err, "删除失败")

			_, err = dao.SysRoleMenu.Ctx(ctx).TX(tx).Where(dao.SysRoleMenu.Columns().RoleId, req.RoleId).Delete()
			util.ErrIsNil(ctx, err, "删除失败")

			// 删除casbin角色
			_, err = consts.Casbin.DeleteRole(roleOld.RoleKey)
			util.ErrIsNil(ctx, err, "删除失败")

			//// 删除casbin权限关联
			//for _, menu := range menusOldList {
			//	split := strings.Split(menu.Perms, ":")
			//	_, err = consts.Casbin.DeletePermissionForUser(roleOld.RoleKey, split[0], split[1])
			//	util.ErrIsNil(ctx, err, "删除失败")
			//}

			// 缓存
			err = task.CacheRefreshRPD(ctx)
			util.ErrIsNil(ctx, err, "删除失败")
		})
		return err
	}); err != nil {
		g.Log().Line().Errorf(ctx, "RoleDelete err:%s", err.Error())
		return rr.FailedWithMessage("修改失败"), err
	}
	return rr.SuccessWithMessage("删除成功"), err
}
