package controller

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/internal/service"
	"HH_ADMIN/utility/rr"
	"context"
)

var (
	User = cUser{}
)

type cUser struct{}

// Description 用户列表
// Author daixk
// Date 2023-09-07 09:28:16
func (c *cUser) UserList(ctx context.Context, req *api.UserListReq) (res *rr.CommonRes, err error) {
	return service.User().UserList(ctx, req)
}

// Description 用户新增
// Author daixk
// Date 2023-09-07 10:02:23
func (c *cUser) UserAdd(ctx context.Context, req *api.UserAddReq) (res *rr.CommonRes, err error) {
	return service.User().UserAdd(ctx, req)
}

// Description 用户修改
// Author daixk
// Date 2023-09-07 10:18:41
func (c *cUser) UserUpdate(ctx context.Context, req *api.UserUpdateReq) (res *rr.CommonRes, err error) {
	return service.User().UserUpdate(ctx, req)
}

// Description 用户删除与重置
// Author daixk
// Date 2023-09-07 10:40:39
func (c *cUser) UserDelete(ctx context.Context, req *api.UserDeleteReq) (res *rr.CommonRes, err error) {
	return service.User().UserDelete(ctx, req)
}
