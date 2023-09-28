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

type sMenu struct{}

func Menu() *sMenu {
	return &sMenu{}
}

// MenuList 菜单列表
func (s *sMenu) MenuList(ctx context.Context, req *api.MenuListReq) (res *rr.CommonRes, err error) {
	m := g.Model(dao.SysMenu.Table())
	m.Where(dao.SysMenu.Columns().Status, 1)

	if req.MenuName != "" {
		m.Where(dao.SysMenu.Columns().MenuName+" like ?", "%"+req.MenuName+"%")
	}
	if req.Perms != "" {
		m.Where(dao.SysMenu.Columns().Perms+" like ?", "%"+req.Perms+"%")
	}

	count, err := m.Count()
	if err != nil {
		g.Log().Line().Errorf(ctx, "MenuList err:%s", err.Error())
		return rr.Failed(), err
	}
	if count <= 0 {
		return rr.SuccessWithData(map[string]interface{}{
			"total":     0,
			"menu_list": g.Array{},
		}), err
	}

	menus := make([]entity.SysMenu, 0)
	err = m.Page(req.Pn, req.PageSize).OrderAsc(dao.SysMenu.Columns().OrderNum).Scan(&menus)
	if err != nil {
		g.Log().Line().Errorf(ctx, "MenuList err:%s", err.Error())
		return rr.Failed(), err
	}

	menuS := make([]*model.SysMenuS, 0)
	for _, menu := range menus {
		me := new(model.SysMenuS)
		err := gconv.Struct(menu, &me)
		if err != nil {
			return nil, err
		}
		menuS = append(menuS, me)
	}

	menu := util.GetTreeRecursiveMenu(menuS, menus[0].ParentId)
	return rr.SuccessWithData(map[string]interface{}{
		"total":     count,
		"menu_list": menu,
	}), err
}

// MenuAdd 菜单添加
func (s *sMenu) MenuAdd(ctx context.Context, req *api.MenuAddReq) (res *rr.CommonRes, err error) {
	//userId := ctx.Value(consts.CTXUSERID)
	//phone := ctx.Value(consts.CTXPHONE)
	userName := ctx.Value(consts.CTXUSERNAME)

	// 检查
	count, err := dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Perms, req.Perms).Count()
	if err != nil {
		g.Log().Line().Errorf(ctx, "MenuAdd err:%s", err.Error())
		return rr.FailedWithMessage("添加失败"), err
	}
	if err != nil && err != sql.ErrNoRows {
		g.Log().Line().Errorf(ctx, "MenuAdd err:%s", err.Error())
		return rr.FailedWithMessage("添加失败"), err
	}
	if count > 0 {
		return rr.FailedWithMessage("改菜单路径已存在"), err
	}

	roles := make([]entity.SysRole, 0)
	var roleOne entity.SysRole
	if len(req.RoleId) > 0 {
		for _, roleId := range req.RoleId {
			get, err := g.Redis().HMGet(ctx, consts.CACHEROLEHM, gconv.String(roleId))
			if err != nil {
				g.Log().Line().Errorf(ctx, "MenuAdd err:%s", err.Error())
				return rr.FailedWithMessage("添加失败"), err
			}
			err = gconv.Scan(get[0], &roleOne)
			if err != nil {
				g.Log().Line().Errorf(ctx, "MenuAdd err:%s", err.Error())
				return rr.FailedWithMessage("添加失败"), err
			}
			roles = append(roles, roleOne)
		}
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 添加菜单
			id, err := dao.SysMenu.Ctx(ctx).TX(tx).InsertAndGetId(entity.SysMenu{
				MenuName:   req.MenuName,
				ParentId:   req.ParentId,
				OrderNum:   req.OrderNum,
				Path:       req.Path,
				Component:  req.Component,
				Query:      "",
				IsFrame:    req.IsFrame,
				IsCache:    req.IsCache,
				MenuType:   req.MenuType,
				Visible:    req.Visible,
				Status:     req.Status,
				Perms:      req.Perms,
				Icon:       req.Icon,
				CreateBy:   gconv.String(userName),
				Remark:     req.Remark,
				CreateTime: gtime.Now(),
			})
			util.ErrIsNil(ctx, err, "添加失败")

			menus := make([]entity.SysRoleMenu, 0)
			if len(req.RoleId) > 0 {
				for _, role := range req.RoleId {
					menus = append(menus, entity.SysRoleMenu{
						RoleId: role,
						MenuId: id,
					})
				}
				_, err := dao.SysRoleMenu.Ctx(ctx).TX(tx).Data(menus).Insert()
				util.ErrIsNil(ctx, err, "添加失败")
			}

			// 添加casbin角色的权限
			for _, role := range roles {
				_, err := consts.Casbin.AddPermissionForUser(role.RoleKey, req.Path, req.Method)
				util.ErrIsNil(ctx, err, "添加失败")
			}

			// 缓存
			err = task.CacheRefreshRPD(ctx)
			util.ErrIsNil(ctx, err, "添加失败")
		})
		return err
	}); err != nil {
		g.Log().Line().Errorf(ctx, "MenuAdd err:%s", err.Error())
		return rr.FailedWithMessage("添加失败"), err
	}
	return rr.SuccessWithMessage("添加成功"), nil
}

// MenuUpdate 菜单修改
func (s *sMenu) MenuUpdate(ctx context.Context, req *api.MenuUpdateReq) (res *rr.CommonRes, err error) {
	//userId := ctx.Value(consts.CTXUSERID)
	//phone := ctx.Value(consts.CTXPHONE)
	userName := ctx.Value(consts.CTXUSERNAME)

	menuKey := new(entity.SysMenu)
	err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Perms, req.Perms).Limit(1).Scan(&menuKey)
	if err != nil && err != sql.ErrNoRows {
		g.Log().Line().Errorf(ctx, "MenuUpdate err:%s", err.Error())
		return rr.FailedWithMessage("修改失败"), err
	}
	if err != sql.ErrNoRows && menuKey.MenuId != req.MenuId {
		return rr.FailedWithMessage("菜单Perms为" + req.Perms + "的菜单已存在"), err
	}

	// 原menu
	menu := new(entity.SysMenu)
	err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().MenuId, req.MenuId).Scan(&menu)
	if err != nil {
		g.Log().Line().Errorf(ctx, "MenuUpdate err:%s", err.Error())
		return rr.FailedWithMessage("修改失败"), err
	}

	// 原role关联menu
	menusOldList := make([]entity.SysRoleMenu, 0)
	err = dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().MenuId, req.MenuId).Scan(&menusOldList)
	if err != nil {
		g.Log().Line().Errorf(ctx, "MenuUpdate err:%s", err.Error())
		return rr.FailedWithMessage("修改失败"), err
	}
	roles := make([]int64, 0)
	for _, roleMenu := range menusOldList {
		roles = append(roles, roleMenu.RoleId)
	}

	// 原来role列表
	sysRoles := make([]entity.SysRole, 0)
	for _, role := range roles {
		var sysRole entity.SysRole
		get, err := g.Redis().HMGet(ctx, consts.CACHEROLEHM, gconv.String(role))
		if err != nil {
			g.Log().Line().Errorf(ctx, "MenuUpdate err:%s", err.Error())
			return rr.FailedWithMessage("修改失败"), err
		}
		err = gconv.Scan(get[0], &sysRole)
		if err != nil {
			g.Log().Line().Errorf(ctx, "MenuUpdate err:%s", err.Error())
			return rr.FailedWithMessage("修改失败"), err
		}
		sysRoles = append(sysRoles, sysRole)
	}

	// 新role列表
	rolesNew := make([]entity.SysRole, 0)
	if len(req.RoleId) > 0 {
		err := dao.SysRole.Ctx(ctx).WhereIn(dao.SysRole.Columns().RoleId, req.RoleId).Scan(&rolesNew)
		if err != nil {
			g.Log().Line().Errorf(ctx, "MenuUpdate err:%s", err.Error())
			return rr.FailedWithMessage("修改失败"), err
		}
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			sysMenu := entity.SysMenu{
				MenuName:   req.MenuName,
				ParentId:   req.ParentId,
				OrderNum:   req.OrderNum,
				Path:       req.Path,
				Component:  req.Component,
				Query:      "",
				IsFrame:    req.IsFrame,
				IsCache:    req.IsCache,
				MenuType:   req.MenuType,
				Visible:    req.Visible,
				Status:     req.Status,
				Perms:      req.Path + ":" + req.Method,
				Icon:       req.Icon,
				UpdateBy:   gconv.String(userName),
				Remark:     req.Remark,
				UpdateTime: gtime.Now(),
			}
			_, err := dao.SysMenu.Ctx(ctx).TX(tx).OmitEmpty().Data(sysMenu).Where(dao.SysMenu.Columns().MenuId, req.MenuId).Update()
			util.ErrIsNil(ctx, err, "更新失败")

			_, err = dao.SysRoleMenu.Ctx(ctx).TX(tx).Where(dao.SysRoleMenu.Columns().MenuId, req.MenuId).Delete()
			util.ErrIsNil(ctx, err, "更新失败")

			split := strings.Split(menu.Perms, ":")
			for _, role := range sysRoles {
				g.Dump(split[0])
				g.Dump(split[1])
				_, err := consts.Casbin.DeletePermissionForUser(role.RoleKey, split[0], split[1])
				util.ErrIsNil(ctx, err, "更新失败")
			}

			menus := make([]entity.SysRoleMenu, 0)
			if len(req.RoleId) > 0 {
				for _, role := range req.RoleId {
					menus = append(menus, entity.SysRoleMenu{
						RoleId: gconv.Int64(role),
						MenuId: req.MenuId,
					})
				}
				_, err := dao.SysRoleMenu.Ctx(ctx).TX(tx).Data(menus).Insert()
				util.ErrIsNil(ctx, err, "更新失败")
			}

			//_, err = dao.CasbinRule.Ctx(ctx).TX(tx).Where(dao.CasbinRule.Columns().V1, split[0]).Where(dao.CasbinRule.Columns().V2, split[1]).Delete()
			//util.ErrIsNil(ctx, err, "删除casbin权限失败")

			// 添加casbin角色的权限
			for _, role := range rolesNew {
				_, err := consts.Casbin.AddPermissionForUser(role.RoleKey, req.Path, req.Method)
				util.ErrIsNil(ctx, err, "更新失败")
			}

			// 缓存
			err = task.CacheRefreshRPD(ctx)
			util.ErrIsNil(ctx, err, "更新失败")
		})
		return err
	}); err != nil {
		g.Log().Line().Errorf(ctx, "MenuAdd err:%s", err.Error())
		return rr.FailedWithMessage("更新失败"), err
	}
	return rr.SuccessWithMessage("更新成功"), err
}

// MenuDelete 菜单删除
func (s *sMenu) MenuDelete(ctx context.Context, req *api.MenuDeleteReq) (res *rr.CommonRes, err error) {

	// 原menu
	menu := new(entity.SysMenu)
	err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().MenuId, req.MenuId).Scan(&menu)
	if err != nil {
		g.Log().Line().Errorf(ctx, "MenuDelete err:%s", err.Error())
		return rr.FailedWithMessage("删除失败"), err
	}

	// 原role关联menu
	menus := make([]entity.SysRoleMenu, 0)
	err = dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().MenuId, req.MenuId).Scan(&menus)
	if err != nil {
		g.Log().Line().Errorf(ctx, "MenuDelete err:%s", err.Error())
		return rr.FailedWithMessage("删除失败"), err
	}

	roleIds := make([]int64, 0)
	for _, roleMenu := range menus {
		roleIds = append(roleIds, roleMenu.RoleId)
	}

	roles := make([]entity.SysRole, 0)
	var roleOne entity.SysRole
	for _, id := range roleIds {
		get, err := g.Redis().HMGet(ctx, consts.CACHEROLEHM, gconv.String(id))
		if err != nil {
			g.Log().Line().Errorf(ctx, "MenuDelete err:%s", err.Error())
			return rr.FailedWithMessage("删除失败"), err
		}
		err = gconv.Scan(get, &roleOne)
		if err != nil {
			g.Log().Line().Errorf(ctx, "MenuDelete err:%s", err.Error())
			return rr.FailedWithMessage("删除失败"), err
		}
		roles = append(roles, roleOne)
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 删除菜单
			_, err := dao.SysMenu.Ctx(ctx).TX(tx).Data(g.Map{dao.SysMenu.Columns().Status: 2}).Where(dao.SysMenu.Columns().MenuId, req.MenuId).Update()
			util.ErrIsNil(ctx, err, "删除菜单失败")

			// 删除角色菜单关联
			_, err = dao.SysRoleMenu.Ctx(ctx).TX(tx).Where(dao.SysRoleMenu.Columns().MenuId, req.MenuId).Delete()
			util.ErrIsNil(ctx, err, "删除角色菜单关联失败")

			// 删除casbin权限
			for _, role := range roles {
				split := strings.Split(menu.Perms, ":")
				_, err := consts.Casbin.DeletePermissionForUser(role.RoleKey, split[0], split[1])
				util.ErrIsNil(ctx, err, "删除casbin权限失败")
			}

			// 缓存
			err = task.CacheRefreshRPD(ctx)
			util.ErrIsNil(ctx, err, "删除失败")
		})
		return err
	}); err != nil {
		g.Log().Line().Errorf(ctx, "MenuDelete err:%s", err.Error())
		return rr.FailedWithMessage("删除失败"), err
	}
	return rr.SuccessWithMessage("删除成功"), err
}
