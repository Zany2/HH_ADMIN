package controller

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/internal/service"
	"HH_ADMIN/utility/rr"
	"context"
)

var (
	Role = cRole{}
)

type cRole struct{}

// Description 角色列表
// Author daixk
// Date 2023-09-07 11:30:29
func (c *cRole) RoleList(ctx context.Context, req *api.RoleListReq) (res *rr.CommonRes, err error) {
	return service.Role().RoleList(ctx, req)
}

// Description 角色新增
// Author daixk
// Date 2023-09-18 10:53:17
func (c *cRole) RoleAdd(ctx context.Context, req *api.RoleAddReq) (res *rr.CommonRes, err error) {
	return service.Role().RoleAdd(ctx, req)
}

// Description 角色修改
// Author daixk
// Date 2023-09-18 10:53:17
func (c *cRole) RoleUpdate(ctx context.Context, req *api.RoleUpdateReq) (res *rr.CommonRes, err error) {
	return service.Role().RoleUpdate(ctx, req)
}

// Description 角色删除
// Author daixk
// Date 2023-09-18 16:39:55
func (c *cRole) RoleDelete(ctx context.Context, req *api.RoleDeleteReq) (res *rr.CommonRes, err error) {
	return service.Role().RoleDelete(ctx, req)
}

// Description 角色查询
// Author daixk
// Date 2023-09-07 13:41:54
//func (c *cUser) RoleQuery(ctx context.Context, req *api.RoleQueryReq) (res *api.RoleQueryRes, err error) {
//	return service.Role().RoleQuery(ctx, req)
//}
