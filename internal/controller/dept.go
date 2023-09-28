package controller

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/internal/service"
	"HH_ADMIN/utility/rr"
	"context"
)

var (
	Dept = cDept{}
)

type cDept struct{}

// DeptList 部门列表
func (s *cDept) DeptList(ctx context.Context, req *api.DeptListReq) (res *rr.CommonRes, err error) {
	return service.Dept().DeptList(ctx, req)
}

// DeptAdd 部门新增
func (s *cDept) DeptAdd(ctx context.Context, req *api.DeptAddReq) (res *rr.CommonRes, err error) {
	return service.Dept().DeptAdd(ctx, req)
}

// DeptUpdate 部门修改
func (s *cDept) DeptUpdate(ctx context.Context, req *api.DeptUpdateReq) (res *rr.CommonRes, err error) {
	return service.Dept().DeptUpdate(ctx, req)
}

// DeptDelete 部门删除
func (s *cDept) DeptDelete(ctx context.Context, req *api.DeptDeleteReq) (res *rr.CommonRes, err error) {
	return service.Dept().DeptDelete(ctx, req)
}
