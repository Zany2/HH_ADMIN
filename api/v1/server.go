package api

import "github.com/gogf/gf/v2/frame/g"

type ServerListReq struct {
	g.Meta `path:"/monitor/server/manage" method:"get"`
}
