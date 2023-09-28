package service

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/internal/dao"
	"HH_ADMIN/internal/model/entity"
	"HH_ADMIN/utility/rr"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type sLogin struct{}

func Login() *sLogin {
	return &sLogin{}
}

// LoginList 登录列表
func (c *sLogin) LoginList(ctx context.Context, req *api.LoginListReq) (res *rr.CommonRes, err error) {
	m := g.Model(dao.SysLoginLog.Table())
	if req.Ip != "" {
		m.Where(dao.SysLoginLog.Columns().Ipaddr+" like ?", "%"+req.Ip+"%")
	}
	if req.UserName != "" {
		m.Where(dao.SysLoginLog.Columns().LoginName+" like ?", "%"+req.UserName+"%")
	}
	if req.Status != 0 {
		m.Where(dao.SysLoginLog.Columns().Status, req.Status)
	}
	if req.BeforeTime != "" {
		m.WhereGTE(dao.SysLoginLog.Columns().LoginTime, req.BeforeTime)
	}
	if req.EndTime != "" {
		m.WhereLTE(dao.SysLoginLog.Columns().LoginTime, req.BeforeTime)
	}

	var total int
	logs := make([]entity.SysLoginLog, 0)
	err = m.Page(req.Pn, req.PageSize).OrderDesc(dao.SysLoginLog.Columns().LoginTime).ScanAndCount(&logs, &total, false)
	if err != nil {
		g.Log().Line().Errorf(ctx, "LoginList err:%s", err.Error())
		return rr.Failed(), err
	}

	if total <= 0 {
		return rr.SuccessWithData(map[string]interface{}{
			"count":      0,
			"login_list": g.Array{},
		}), err
	}

	return rr.SuccessWithData(map[string]interface{}{
		"count":      total,
		"login_list": logs,
	}), err
}
