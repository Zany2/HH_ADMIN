package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

type Req struct {
	g.Meta `path:"/hello" method:"get"`
}
type Res struct {
	Id int64 `json:"id"`
}
