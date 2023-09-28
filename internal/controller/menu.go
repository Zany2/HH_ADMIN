package controller

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/internal/service"
	"HH_ADMIN/utility/rr"
	"context"
)

var (
	Menu = cMenu{}
)

type cMenu struct{}

// Description 菜单列表
// Author daixk
// Date 2023-09-07 11:30:29
func (c *cMenu) MenuList(ctx context.Context, req *api.MenuListReq) (res *rr.CommonRes, err error) {
	return service.Menu().MenuList(ctx, req)
}

// Description 菜单添加
// Author daixk
// Data 2023-09-18 21:45:03
func (c *cMenu) MenuAdd(ctx context.Context, req *api.MenuAddReq) (res *rr.CommonRes, err error) {
	return service.Menu().MenuAdd(ctx, req)
}

// Description 菜单修改
// Author daixk
// Data 2023-09-18 22:44:18
func (c *cMenu) MenuUpdate(ctx context.Context, req *api.MenuUpdateReq) (res *rr.CommonRes, err error) {
	return service.Menu().MenuUpdate(ctx, req)
}

// Description 菜单删除
// Author daixk
// Date 2023-09-19 09:23:01
func (c *cMenu) MenuDelete(ctx context.Context, req *api.MenuDeleteReq) (res *rr.CommonRes, err error) {
	return service.Menu().MenuDelete(ctx, req)
}
