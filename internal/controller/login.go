package controller

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/internal/service"
	"HH_ADMIN/utility/rr"
	"context"
)

var (
	Login = cLogin{}
)

type cLogin struct{}

// LoginList 登录列表
func (s *cLogin) LoginList(ctx context.Context, req *api.LoginListReq) (res *rr.CommonRes, err error) {
	return service.Login().LoginList(ctx, req)
}
