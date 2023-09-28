package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileReq struct {
	g.Meta `path:"/upload" method:"post"`
	Files  ghttp.UploadFiles `json:"files" type:"file"`
}
