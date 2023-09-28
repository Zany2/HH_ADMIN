package logic

import (
	"HH_ADMIN/internal/consts"
	"HH_ADMIN/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// UserParentDeptIdList 查找User父id列表
func UserParentDeptIdList(ctx context.Context, deptId int64) (deptParentIdList []int64, err error) {
	depts := make([]*entity.SysDept, 0)
	deptsRedisCache, err := g.Redis().Get(ctx, consts.CACHEDEPT)

	err = gconv.Struct(deptsRedisCache, &depts)
	if err != nil {
		g.Log().Line().Errorf(ctx, "UserParentDeptIdList err:%s", err.Error())
		return nil, err
	}

	deptsParentIds := findSonByParentId(depts, deptId)

	deptParentIdList = make([]int64, 0)
	for _, id := range deptsParentIds {
		deptParentIdList = append(deptParentIdList, id.DeptId)
	}
	deptParentIdList = append(deptParentIdList, deptId)
	return
}

func findSonByParentId(deptList []*entity.SysDept, deptId int64) []*entity.SysDept {
	children := make([]*entity.SysDept, 0, len(deptList))
	for _, v := range deptList {
		if v.ParentId == deptId {
			children = append(children, v)
			fChildren := findSonByParentId(deptList, v.DeptId)
			children = append(children, fChildren...)
		}
	}
	return children
}
