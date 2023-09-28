package api

import (
	"HH_ADMIN/utility/rr"
	"github.com/gogf/gf/v2/frame/g"
)

type PostListReq struct {
	g.Meta `path:"/auth/post/list" method:"get"`
	rr.CommonReq
	PostName string `json:"post_name"`
	PostCode string `json:"post_code"`
	Status   int    `json:"status" v:"in:0,1,2 #状态必须在0,1,2之间"`
}

type PostAddReq struct {
	g.Meta   `path:"/auth/post/add" method:"post"`
	PostCode string `json:"post_code"` // 岗位编码
	PostName string `json:"post_name"` // 岗位名称
	PostSort int    `json:"post_sort"` // 显示顺序
	Status   int    `json:"status"`    // 状态（1正常 2停用）
	Remark   string `json:"remark"`    // 备注
}

type PostUpdateReq struct {
	g.Meta   `path:"/auth/post/update" method:"post"`
	PostId   int64  `json:"post_id"`   // 岗位ID
	PostCode string `json:"post_code"` // 岗位编码
	PostName string `json:"post_name"` // 岗位名称
	PostSort int    `json:"post_sort"` // 显示顺序
	Status   int    `json:"status"`    // 状态（1正常 2停用）
	Remark   string `json:"remark"`    // 备注
}

type PostDeleteReq struct {
	g.Meta `path:"/auth/post/delete" method:"post"`
	PostId int64 `json:"post_id"` // 岗位ID
}
