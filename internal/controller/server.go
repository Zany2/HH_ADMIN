package controller

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/internal/service"
	"HH_ADMIN/utility/rr"
	"context"
)

var (
	Server = cServer{}
)

type cServer struct{}

// Description 服务监控查询
// Author daixk
// Date 2023-09-07 11:30:29
func (c *cServer) ServerList(ctx context.Context, req *api.ServerListReq) (res *rr.CommonRes, err error) {
	return service.Server().ServerList(ctx, req)
}
