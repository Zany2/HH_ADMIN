package service

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/internal/consts"
	"HH_ADMIN/internal/dao"
	"HH_ADMIN/internal/model/entity"
	"HH_ADMIN/task"
	"HH_ADMIN/util"
	"HH_ADMIN/utility/rr"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sPost struct{}

func Post() *sPost {
	return &sPost{}
}

// PostList 岗位列表
func (s *sPost) PostList(ctx context.Context, req *api.PostListReq) (res *rr.CommonRes, err error) {
	m := g.Model(dao.SysPost.Table())
	m.Where(dao.SysPost.Columns().DelFlag, 1)
	if req.PostName != "" {
		m.Where(dao.SysPost.Columns().PostName+" like ?", "%"+req.PostName+"%")
	}
	if req.PostCode != "" {
		m.Where(dao.SysPost.Columns().PostCode+" like ?", "%"+req.PostCode+"%")
	}
	if req.Status != 0 {
		m.Where(dao.SysPost.Columns().Status, req.Status)
	}

	var total int
	posts := make([]entity.SysPost, 0)
	err = m.Page(req.Pn, req.PageSize).ScanAndCount(&posts, &total, false)
	if err != nil {
		g.Log().Line().Errorf(ctx, "PostList err:%s", err.Error())
		return rr.Failed(), err
	}

	return rr.SuccessWithData(map[string]interface{}{
		"total":     total,
		"post_list": posts,
	}), err
}

// PostAdd 岗位新增
func (s *sPost) PostAdd(ctx context.Context, req *api.PostAddReq) (res *rr.CommonRes, err error) {
	//userId := ctx.Value(consts.CTXUSERID)
	//phone := ctx.Value(consts.CTXPHONE)
	userName := ctx.Value(consts.CTXUSERNAME)

	post := entity.SysPost{
		PostCode:   req.PostCode,
		PostName:   req.PostName,
		PostSort:   req.PostSort,
		Status:     1,
		DelFlag:    1,
		CreateBy:   gconv.String(userName),
		CreateTime: gtime.Now(),
		Remark:     req.Remark,
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysPost.Ctx(ctx).TX(tx).Data(post).Insert()
			util.ErrIsNil(ctx, err, "新增失败")

			// 缓存
			err = task.CacheRefreshRPD(ctx)
			util.ErrIsNil(ctx, err, "新增失败")

		})
		return err
	}); err != nil {
		g.Log().Line().Errorf(ctx, "PostAdd err:%s", err.Error())
		return rr.FailedWithMessage("新增失败"), err
	}

	return rr.SuccessWithMessage("新增成功"), err
}

// PostUpdate 岗位修改
func (s *sPost) PostUpdate(ctx context.Context, req *api.PostUpdateReq) (res *rr.CommonRes, err error) {
	//userId := ctx.Value(consts.CTXUSERID)
	//phone := ctx.Value(consts.CTXPHONE)
	userName := ctx.Value(consts.CTXUSERNAME)

	post := entity.SysPost{
		PostCode:   req.PostCode,
		PostName:   req.PostName,
		PostSort:   req.PostSort,
		Status:     req.Status,
		UpdateBy:   gconv.String(userName),
		UpdateTime: gtime.Now(),
		Remark:     req.Remark,
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysPost.Ctx(ctx).TX(tx).OmitEmpty().Data(post).Where(dao.SysPost.Columns().PostId, req.PostId).Update()
			util.ErrIsNil(ctx, err, "修改失败")

			err = task.CacheRefreshRPD(ctx)
			util.ErrIsNil(ctx, err, "修改失败")

		})
		return err
	}); err != nil {
		g.Log().Line().Errorf(ctx, "PostAdd err:%s", err.Error())
		return rr.FailedWithMessage("修改失败"), err
	}

	return rr.SuccessWithMessage("修改成功"), nil
}

// PostDelete 岗位删除
func (s *sPost) PostDelete(ctx context.Context, req *api.PostDeleteReq) (res *rr.CommonRes, err error) {

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysPost.Ctx(ctx).TX(tx).Data(g.Map{dao.SysPost.Columns().DelFlag: 2}).Where(dao.SysPost.Columns().PostId, req.PostId).Update()
			util.ErrIsNil(ctx, err, "删除失败")

			_, err = dao.SysUserPost.Ctx(ctx).TX(tx).Where(dao.SysUserPost.Columns().PostId, req.PostId).Delete()
			util.ErrIsNil(ctx, err, "删除失败")

			err = task.CacheRefreshRPD(ctx)
			util.ErrIsNil(ctx, err, "删除失败")

		})
		return err
	}); err != nil {
		g.Log().Line().Errorf(ctx, "PostDelete err:%s", err.Error())
		return rr.FailedWithMessage("删除失败"), err
	}

	return rr.SuccessWithMessage("删除成功"), nil
}
