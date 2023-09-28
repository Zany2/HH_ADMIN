package api

import (
	"HH_ADMIN/utility/rr"
	"github.com/gogf/gf/v2/frame/g"
)

type LoginListReq struct {
	g.Meta `path:"/monitor/login/manage" method:"get"`
	rr.CommonReq
	Ip         string `json:"ip"`          // 登录Ip
	UserName   string `json:"user_name"`   // 登录账号
	Status     int    `json:"status"`      // 状态 1成功 2失败
	BeforeTime string `json:"before_time"` // 开始时间
	EndTime    string `json:"end_time"`    // 结束时间
}
