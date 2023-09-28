package controller

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/internal/service"
	"HH_ADMIN/utility/rr"
	"context"
)

var (
	File = cFile{}
)

type cFile struct{}

func (c *cFile) File(ctx context.Context, req *api.FileReq) (res *rr.CommonRes, err error) {
	upload, err := service.File().FileUpload(ctx, req)
	return upload, err
}
