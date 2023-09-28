package service

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/casbinhh"
	"HH_ADMIN/internal/consts"
	"HH_ADMIN/internal/dao"
	"HH_ADMIN/internal/logic"
	"HH_ADMIN/internal/model/entity"
	"HH_ADMIN/task"
	"HH_ADMIN/util"
	"HH_ADMIN/utility/rr"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sUser struct{}

func User() *sUser {
	return &sUser{}
}

// UserList 用户列表
func (s *sUser) UserList(ctx context.Context, req *api.UserListReq) (res *rr.CommonRes, err error) {
	m := g.Model(dao.SysUser.Table())
	m.Where(dao.SysUser.Columns().DelFlag, 1)
	if req.Status != 0 {
		m.Where(dao.SysUser.Columns().Status, req.Status)
	}
	if req.Phone != "" {
		m.Where(dao.SysUser.Columns().Phone+" like ?", "%"+req.Phone+"%")
	}
	if req.KeyWord != "" {
		m.Where(dao.SysUser.Columns().UserName+" like ?", "%"+req.KeyWord+"%").WhereOr(dao.SysUser.Columns().NickName+" like ?", "%"+req.KeyWord+"%")
	}
	if req.Stime != "" {
		m.WhereGTE(dao.SysUser.Columns().CreateTime, req.Stime)
	}
	if req.Etime != "" {
		m.WhereLTE(dao.SysUser.Columns().CreateTime, req.Etime)
	}
	if req.DeptId != 0 {
		list, err := logic.UserParentDeptIdList(ctx, req.DeptId)
		if err != nil {
			g.Log().Line().Errorf(ctx, "UserList err:%s", err.Error())
			return rr.Failed(), err
		}
		if len(list) > 0 {
			m.WhereIn(dao.SysUser.Columns().DeptId, list)
		}
	}

	var total int
	users := make([]entity.SysUser, 0)

	err = m.Page(req.Pn, req.PageSize).ScanAndCount(&users, &total, false)
	if err != nil {
		g.Log().Line().Errorf(ctx, "UserList err:%s", err.Error())
		return rr.Failed(), err
	}
	if total <= 0 {
		return rr.SuccessWithData(map[string]interface{}{
			"count":     0,
			"user_list": g.Array{},
		}), err
	}

	listRes := make([]api.UserListRes, 0)
	for _, user := range users {
		// 密码
		user.Password = ""

		// 角色
		allSysUserRole, err := dao.SysUserRole.Ctx(ctx).Fields(dao.SysUserRole.Columns().RoleId).Where(dao.SysUserRole.Columns().UserId, user.UserId).All()
		if err != nil {
			g.Log().Line().Errorf(ctx, "UserList err:%s", err.Error())
			return rr.Failed(), err
		}
		roleIdList := allSysUserRole.Array(dao.SysUserRole.Columns().RoleId)

		roles := make([]entity.SysRole, 0)
		var role entity.SysRole
		for _, i := range roleIdList {
			roleRedisCache, err := g.Redis().HGet(ctx, consts.CACHEROLEHM, gconv.String(i))
			if err != nil {
				g.Log().Line().Errorf(ctx, "UserList err:%s", err.Error())
				return rr.Failed(), err
			}
			if !roleRedisCache.IsEmpty() {
				err = gconv.Struct(roleRedisCache, &role)
				if err != nil {
					g.Log().Line().Errorf(ctx, "UserList err:%s", err.Error())
					return rr.Failed(), err
				}
				roles = append(roles, role)
			}
		}

		// 岗位
		allSysUserPost, err := dao.SysUserPost.Ctx(ctx).Fields(dao.SysUserPost.Columns().PostId).Where(dao.SysUserPost.Columns().UserId, user.UserId).All()
		if err != nil {
			g.Log().Line().Errorf(ctx, "UserList err:%s", err.Error())
			return rr.Failed(), err
		}
		postIdList := allSysUserPost.Array(dao.SysUserPost.Columns().PostId)

		posts := make([]entity.SysPost, 0)
		var post entity.SysPost
		for _, i := range postIdList {
			postRedisCache, err := g.Redis().HGet(ctx, consts.CACHEPOSTHM, gconv.String(i))
			if err != nil {
				g.Log().Line().Errorf(ctx, "UserList err:%s", err.Error())
				return rr.Failed(), err
			}

			fmt.Println(postRedisCache.IsEmpty())

			if !postRedisCache.IsEmpty() {
				err = gconv.Struct(postRedisCache, &post)
				if err != nil {
					g.Log().Line().Errorf(ctx, "UserList err:%s", err.Error())
					return rr.Failed(), err
				}
				posts = append(posts, post)
			}
		}

		// 部门
		var dept entity.SysDept
		deptRedisCache, err := g.Redis().HGet(ctx, consts.CACHEDEPTHM, gconv.String(user.DeptId))
		if err != nil {
			g.Log().Line().Errorf(ctx, "UserList err:%s", err.Error())
			return rr.Failed(), err
		}
		if !deptRedisCache.IsEmpty() || !deptRedisCache.IsNil() {
			err = gconv.Struct(deptRedisCache, &dept)
			if err != nil {
				g.Log().Line().Errorf(ctx, "UserList err:%s", err.Error())
				return rr.Failed(), err
			}
		}

		listRes = append(listRes, api.UserListRes{
			SysUser: user,
			SysDept: dept,
			SysRole: roles,
			SysPost: posts,
		})
	}

	return rr.SuccessWithData(map[string]interface{}{
		"total":     total,
		"user_list": listRes,
	}), err
}

// UserAdd 用户新增
func (s *sUser) UserAdd(ctx context.Context, req *api.UserAddReq) (res *rr.CommonRes, err error) {
	//userId := ctx.Value(consts.CTXUSERID)
	//phone := ctx.Value(consts.CTXPHONE)
	userName := ctx.Value(consts.CTXUSERNAME)

	count, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserName, req.UserName).Count()
	if err != nil {
		g.Log().Line().Errorf(ctx, "UserAdd err:%s", err.Error())
		return rr.FailedWithMessage("新增失败"), nil
	}
	if count > 0 {
		return rr.FailedWithMessage("新增失败 该用户名已存在"), nil
	}

	encrypt, err := gmd5.EncryptString(req.Password)
	if err != nil {
		g.Log().Line().Errorf(ctx, "UserAdd err:%s", err.Error())
		return rr.FailedWithMessage("新增失败"), nil
	}

	// 角色列表
	sysRoles := make([]entity.SysRole, 0)
	var roleOne entity.SysRole
	if len(req.RoleId) > 0 {
		for _, i := range req.RoleId {
			roleRedisCache, err := g.Redis().HMGet(ctx, consts.CACHEROLEHM, gconv.String(i))
			if err != nil {
				g.Log().Line().Errorf(ctx, "UserAdd err:%s", err.Error())
				return rr.FailedWithMessage("新增失败"), nil
			}
			fmt.Println(roleRedisCache)
			err = gconv.Struct(roleRedisCache[0], &roleOne)
			if err != nil {
				g.Log().Line().Errorf(ctx, "UserAdd err:%s", err.Error())
				return rr.FailedWithMessage("新增失败"), nil
			}
			sysRoles = append(sysRoles, roleOne)
		}
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 用户
			resultId, err := dao.SysUser.Ctx(ctx).TX(tx).InsertAndGetId(entity.SysUser{
				UserName:   req.UserName,
				Password:   encrypt,
				DeptId:     req.DeptId,
				NickName:   req.NickName,
				Email:      req.Email,
				Phone:      req.Phone,
				Sex:        req.Sex,
				Status:     1,
				DelFlag:    1,
				UserType:   2,
				CreateBy:   gconv.String(userName),
				Remark:     req.Remark,
				CreateTime: gtime.Now(),
			})
			util.ErrIsNil(ctx, err, "新增失败")

			// 角色
			if len(req.RoleId) > 0 {
				roles := make([]entity.SysUserRole, 0)
				for _, i := range req.RoleId {
					roles = append(roles, entity.SysUserRole{
						UserId: resultId,
						RoleId: i,
					})
				}
				_, err = dao.SysUserRole.Ctx(ctx).TX(tx).Data(roles).Insert()
				util.ErrIsNil(ctx, err, "用户角色关联添加失败")

				// 添加 casbin 角色
				err = casbinhh.CasbinAddRoleForUser(req.UserName, sysRoles)
				util.ErrIsNil(ctx, err, "添加casbin角色失败")
			}

			// 岗位
			if len(req.PostId) > 0 {
				posts := make([]entity.SysUserPost, 0)
				for _, i := range req.PostId {
					posts = append(posts, entity.SysUserPost{
						UserId: resultId,
						PostId: i,
					})
				}
				_, err = dao.SysUserPost.Ctx(ctx).TX(tx).Data(posts).Insert()
				util.ErrIsNil(ctx, err, "设置用户岗位失败")
			}

			// 更新缓存部门、岗位、角色信息
			err = task.CacheRefreshRPD(ctx)
			util.ErrIsNil(ctx, err, "更新缓存失败")
		})
		return err
	}); err != nil {
		g.Log().Line().Errorf(ctx, "UserAdd err:%s", err.Error())
		return rr.FailedWithMessage("新增失败"), err
	}

	return rr.SuccessWithMessage("新增成功"), err
}

// UserUpdate 用户修改
func (s *sUser) UserUpdate(ctx context.Context, req *api.UserUpdateReq) (res *rr.CommonRes, err error) {
	//userId := ctx.Value(consts.CTXUSERID)
	//phone := ctx.Value(consts.CTXPHONE)
	userName := ctx.Value(consts.CTXUSERNAME)

	// 原用户
	user := new(entity.SysUser)
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserId, req.UserId).Scan(&user)
	if err != nil {
		g.Log().Line().Errorf(ctx, "UserUpdate err:%s", err.Error())
		return rr.FailedWithMessage("修改失败"), err
	}

	// 原角色id列表
	all, err := dao.SysUserRole.Ctx(ctx).Fields(dao.SysUserRole.Columns().RoleId).Where(dao.SysUserRole.Columns().UserId, req.UserId).All()
	if err != nil {
		g.Log().Line().Errorf(ctx, "UserUpdate err:%s", err.Error())
		return rr.FailedWithMessage("修改失败"), err
	}
	oldRoleList := all.Array(dao.SysUserRole.Columns().RoleId)

	// 原角色列表
	oldRolesList := make([]entity.SysRole, 0)
	var oldRole entity.SysRole
	for _, i := range oldRoleList {
		roleRedisCache, err := g.Redis().HGet(ctx, consts.CACHEROLEHM, gconv.String(i))
		if err != nil {
			g.Log().Line().Errorf(ctx, "UserUpdate err:%s", err.Error())
			return rr.FailedWithMessage("修改失败"), err
		}
		err = gconv.Struct(roleRedisCache, &oldRole)
		if err != nil {
			g.Log().Line().Errorf(ctx, "UserUpdate err:%s", err.Error())
			return rr.FailedWithMessage("修改失败"), err
		}
		oldRolesList = append(oldRolesList, oldRole)
	}

	// 新角色列表
	newRolesList := make([]entity.SysRole, 0)
	var newRole entity.SysRole
	for _, i := range req.RoleId {
		roleRedisCache, err := g.Redis().HGet(ctx, consts.CACHEROLEHM, gconv.String(i))
		if err != nil {
			g.Log().Line().Errorf(ctx, "UserUpdate err:%s", err.Error())
			return rr.FailedWithMessage("修改失败"), err
		}
		err = gconv.Struct(roleRedisCache, &newRole)
		if err != nil {
			g.Log().Line().Errorf(ctx, "UserUpdate err:%s", err.Error())
			return rr.FailedWithMessage("修改失败"), err
		}
		newRolesList = append(newRolesList, newRole)
	}

	// 新用户
	userNew := entity.SysUser{
		DeptId:     req.DeptId,
		NickName:   req.NickName,
		Phone:      req.Phone,
		Email:      req.Email,
		Sex:        req.Sex,
		Status:     req.Status,
		UpdateBy:   gconv.String(userName),
		Remark:     req.Remark,
		UpdateTime: gtime.Now(),
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 修改用户
			_, err := dao.SysUser.Ctx(ctx).TX(tx).Data(&userNew).OmitEmpty().Where(dao.SysUser.Columns().UserId, req.UserId).Update()
			util.ErrIsNil(ctx, err, "修改失败")

			// 修改角色
			// 先删除角色关联表
			_, err = dao.SysUserRole.Ctx(ctx).TX(tx).Where(dao.SysUserRole.Columns().UserId, req.UserId).Delete()
			util.ErrIsNil(ctx, err, "用户角色关联删除失败")

			// 删除casbin角色
			err = casbinhh.CasbinDeleteRoleForUser(user.UserName, oldRolesList)
			util.ErrIsNil(ctx, err, "casbin角色删除失败")
			//for _, sysRole := range oldRolesList {
			//	_, err := consts.Casbin.DeleteRoleForUser(user.UserName, sysRole.RoleKey)
			//	util.ErrIsNil(ctx, err, "casbin角色删除失败")
			//}

			// 设置用户角色关联
			if len(req.RoleId) > 0 {
				roles := make([]entity.SysUserRole, 0)
				for _, i := range req.RoleId {
					roles = append(roles, entity.SysUserRole{
						UserId: req.UserId,
						RoleId: i,
					})
				}
				_, err = dao.SysUserRole.Ctx(ctx).TX(tx).Data(roles).Insert()
				util.ErrIsNil(ctx, err, "用户角色关联添加失败")

				// 添加casbin角色
				err = casbinhh.CasbinAddRoleForUser(user.UserName, newRolesList)
				util.ErrIsNil(ctx, err, "添加casbin角色失败")
				//for _, role := range newRolesList {
				//	_, err = consts.Casbin.AddRoleForUser(user.Phone, role.RoleKey)
				//	util.ErrIsNil(ctx, err, "设置用户角色失败")
				//}
			}

			// 修改岗位
			// 先删除角色岗位关联表
			_, err = dao.SysUserPost.Ctx(ctx).TX(tx).Where(dao.SysUserPost.Columns().UserId, req.UserId).Delete()
			util.ErrIsNil(ctx, err, "用户岗位关联删除失败")

			// 设置用户岗位关联
			if len(req.PostId) > 0 {
				posts := make([]entity.SysUserPost, 0)
				for _, i := range req.PostId {
					posts = append(posts, entity.SysUserPost{
						UserId: req.UserId,
						PostId: i,
					})
				}
				_, err = dao.SysUserPost.Ctx(ctx).TX(tx).Data(posts).Insert()
				util.ErrIsNil(ctx, err, "设置用户岗位关联失败")
			}

			// 更新缓存部门、岗位、角色信息
			err = task.CacheRefreshRPD(ctx)
			util.ErrIsNil(ctx, err, "更新缓存失败")

		})
		return err
	}); err != nil {
		g.Log().Line().Errorf(ctx, "UserUpdate err:%s", err.Error())
		return rr.FailedWithMessage("修改失败"), nil
	}
	return rr.SuccessWithMessage("修改成功"), nil
}

// UserDelete 用户删除
func (s *sUser) UserDelete(ctx context.Context, req *api.UserDeleteReq) (res *rr.CommonRes, err error) {
	switch req.EditType {
	case 1: // 1删除 2重置
		_, err := dao.SysUser.Ctx(ctx).Data(g.Map{dao.SysUser.Columns().DelFlag: 2}).OmitEmpty().Where(dao.SysUser.Columns().UserId, req.UserId).Update()
		if err != nil {
			g.Log().Line().Errorf(ctx, "UserDelete err:%s", err.Error())
			return rr.FailedWithMessage("删除失败"), err
		}
		return rr.SuccessWithMessage("删除成功"), err
	case 2: // 1删除 2重置
		if req.Password == "" {
			return rr.FailedWithMessage("密码不能为空"), nil
		}

		encrypt, err := gmd5.EncryptString(req.Password)
		if err != nil {
			g.Log().Line().Errorf(ctx, "UserDelete err:%s", err.Error())
			return rr.FailedWithMessage("重置失败"), err
		}

		_, err = dao.SysUser.Ctx(ctx).Data(g.Map{dao.SysUser.Columns().Password: encrypt}).OmitEmpty().Where(dao.SysUser.Columns().UserId, req.UserId).Update()
		if err != nil {
			g.Log().Line().Errorf(ctx, "UserDelete err:%s", err.Error())
			return rr.FailedWithMessage("重置失败"), nil
		}

		return rr.SuccessWithMessage("重置成功"), err
	default:
		return rr.FailedWithMessage("类型必须在1,2之间"), err
	}
}
