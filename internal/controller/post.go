package controller

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/internal/service"
	"HH_ADMIN/utility/rr"
	"context"
)

var (
	Post = cPost{}
)

type cPost struct{}

// Description 岗位列表
// Author daixk
// Data 2023-09-17 15:15:07
func (c *cPost) PostList(ctx context.Context, req *api.PostListReq) (res *rr.CommonRes, err error) {
	return service.Post().PostList(ctx, req)
}

// Description 岗位新增
// Author daixk
// Data 2023-09-17 15:14:13
func (c *cPost) PostAdd(ctx context.Context, req *api.PostAddReq) (res *rr.CommonRes, err error) {
	return service.Post().PostAdd(ctx, req)
}

// Description 岗位修改
// Author daixk
// Data 2023-09-17 15:30:35
func (c *cPost) PostUpdate(ctx context.Context, req *api.PostUpdateReq) (res *rr.CommonRes, err error) {
	return service.Post().PostUpdate(ctx, req)
}

// Description 岗位删除
// Author daixk
// Data 2023-09-17 19:47:34
func (c *cPost) PostDelete(ctx context.Context, req *api.PostDeleteReq) (res *rr.CommonRes, err error) {
	return service.Post().PostDelete(ctx, req)
}
