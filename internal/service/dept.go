package service

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/internal/consts"
	"HH_ADMIN/internal/dao"
	"HH_ADMIN/internal/logic"
	"HH_ADMIN/internal/model"
	"HH_ADMIN/internal/model/entity"
	"HH_ADMIN/task"
	"HH_ADMIN/util"
	"HH_ADMIN/utility/rr"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sort"
	"strings"
)

type sDept struct{}

func Dept() *sDept {
	return &sDept{}
}

// DeptList 部门列表
func (c *sDept) DeptList(ctx context.Context, req *api.DeptListReq) (res *rr.CommonRes, err error) {
	m := g.Model(dao.SysDept.Table())
	m.Where(dao.SysDept.Columns().DelFlag, 1)
	if req.Status != 0 {
		m.Where(dao.SysUser.Columns().Status, req.Status)
	}
	if req.DeptName != "" {
		m.Where(dao.SysDept.Columns().DeptName+" like ?", "%"+req.DeptName+"%")
	}

	var total int
	depts := make([]entity.SysDept, 0)
	err = m.Page(req.Pn, req.PageSize).OrderAsc(dao.SysDept.Columns().OrderNum).ScanAndCount(&depts, &total, false)
	if err != nil {
		g.Log().Line().Errorf(ctx, "DeptList err:%s", err.Error())
		return rr.Failed(), err
	}
	if total <= 0 {
		return rr.SuccessWithData(map[string]interface{}{
			"total":     0,
			"dept_list": g.Array{},
		}), err
	}

	sysDeptSs := make([]*model.SysDeptList, 0)
	for _, dept := range depts {
		dep := new(model.SysDeptList)
		err := gconv.Struct(dept, &dep)
		if err != nil {
			g.Log().Line().Errorf(ctx, "DeptList err:%s", err.Error())
			return rr.Failed(), err
		}
		sysDeptSs = append(sysDeptSs, dep)
	}

	// 递归拼接菜单
	recursive := util.GetTreeRecursive(sysDeptSs, sysDeptSs[0].ParentId)
	return rr.SuccessWithData(map[string]interface{}{
		"total":     total,
		"dept_list": recursive,
	}), err
}

// DeptAdd 部门新增
func (c *sDept) DeptAdd(ctx context.Context, req *api.DeptAddReq) (res *rr.CommonRes, err error) {
	//userId := ctx.Value(consts.CTXUSERID)
	//phone := ctx.Value(consts.CTXPHONE)
	userName := ctx.Value(consts.CTXUSERNAME)

	ids := make([]string, 0)
	// todo 父级id字符串
	if req.ParentId != 0 {
		depts := make([]entity.SysDept, 0)
		deptsRedisCache, err := g.Redis().Get(ctx, consts.CACHEDEPT)
		if err != nil {
			g.Log().Line().Errorf(ctx, "DeptAdd err:%s", err.Error())
			return rr.FailedWithMessage("新增失败"), err
		}
		err = gconv.Struct(deptsRedisCache, &depts)
		if err != nil {
			g.Log().Line().Errorf(ctx, "DeptAdd err:%s", err.Error())
			return rr.FailedWithMessage("新增失败"), err
		}

		var dept entity.SysDept
		for i := 0; i < len(depts); i++ {
			if depts[i].DeptId == req.ParentId {
				dept = depts[i]
			}
		}

		list := logic.DeptParentIdList(depts, dept)
		ids = append(ids, "0")
		ids = append(ids, gconv.String(dept.DeptId))
		for _, sysDept := range list {
			ids = append(ids, gconv.String(sysDept.DeptId))
		}
		sort.Slice(ids, func(i, j int) bool {
			return ids[i] < ids[j]
		})
	}

	dept := entity.SysDept{
		ParentId:   req.ParentId,
		Ancestors:  strings.Join(ids, ","),
		DeptName:   req.DeptName,
		OrderNum:   req.OrderNum,
		Leader:     req.Leader,
		Phone:      req.Phone,
		Email:      req.Email,
		Status:     req.Status,
		DelFlag:    1,
		CreateBy:   gconv.String(userName),
		CreateTime: gtime.Now(),
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysDept.Ctx(ctx).TX(tx).Data(dept).OmitEmpty().Insert()
			util.ErrIsNil(ctx, err, "新增失败")

			// 刷新缓存
			err = task.CacheRefreshRPD(ctx)
			util.ErrIsNil(ctx, err, "新增失败")
		})
		return err
	}); err != nil {
		g.Log().Line().Errorf(ctx, "DeptAdd err:%s", err.Error())
		return rr.Failed(), err
	}

	return rr.SuccessWithMessage("新增成功"), nil
}

// DeptUpdate 部门修改
func (c *sDept) DeptUpdate(ctx context.Context, req *api.DeptUpdateReq) (res *rr.CommonRes, err error) {
	//userId := ctx.Value(consts.CTXUSERID)
	//phone := ctx.Value(consts.CTXPHONE)
	userName := ctx.Value(consts.CTXUSERNAME)

	dept := entity.SysDept{
		ParentId:   req.ParentId,
		Ancestors:  "",
		DeptName:   req.DeptName,
		OrderNum:   req.OrderNum,
		Leader:     req.Leader,
		Phone:      req.Phone,
		Email:      req.Email,
		Status:     req.Status,
		UpdateBy:   gconv.String(userName),
		UpdateTime: gtime.Now(),
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 修改原数据
			_, err = dao.SysDept.Ctx(ctx).TX(tx).OmitEmpty().Data(dept).Update()
			util.ErrIsNil(ctx, err, "修改失败")

			//// 修改全部的用户关联部门数据
			//if req.Status == 2 {
			//	_, err := dao.SysUser.Ctx(ctx).TX(tx).Data(g.Map{dao.SysUser.Columns().DeptId: 0}).OmitEmpty().Where(dao.SysUser.Columns().DeptId, req.DeptId).Update()
			//	util.ErrIsNil(ctx, err, "修改失败")
			//}

			// 缓存
			err = task.CacheRefreshRPD(ctx)
			util.ErrIsNil(ctx, err, "修改失败")
		})
		return err
	}); err != nil {
		g.Log().Line().Errorf(ctx, "DeptUpdate err:%s", err.Error())
		return rr.Failed(), err
	}

	return rr.SuccessWithMessage("修改成功"), nil
}

// DeptDelete 部门删除
func (c *sDept) DeptDelete(ctx context.Context, req *api.DeptDeleteReq) (res *rr.CommonRes, err error) {

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysDept.Ctx(ctx).Data(g.Map{dao.SysDept.Columns().DelFlag: 2}).Where(dao.SysDept.Columns().DeptId, req.DeptId).Update()
			util.ErrIsNil(ctx, err, "删除失败")

			// 修改全部的用户关联部门数据
			_, err = dao.SysUser.Ctx(ctx).TX(tx).Data(g.Map{dao.SysUser.Columns().DeptId: 0}).Where(dao.SysUser.Columns().DeptId, req.DeptId).Update()
			util.ErrIsNil(ctx, err, "删除失败")

			// 缓存
			err = task.CacheRefreshRPD(ctx)
			util.ErrIsNil(ctx, err, "修改失败")

		})
		return err
	}); err != nil {
		g.Log().Line().Errorf(ctx, "DeptUpdate err:%s", err.Error())
		return rr.Failed(), err
	}

	return rr.SuccessWithMessage("删除成功"), nil
}
