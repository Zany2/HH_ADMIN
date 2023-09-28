package api

import (
	"HH_ADMIN/utility/rr"
	"github.com/gogf/gf/v2/frame/g"
)

type DeptListReq struct {
	g.Meta `path:"/auth/dept/list" method:"get"`
	rr.CommonReq
	DeptName string `json:"dept_name"`
	Status   int    `json:"status" v:"in:0,1,2 #状态必须在0,1,2之间"`
}

type DeptAddReq struct {
	g.Meta   `path:"/auth/dept/add" method:"post"`
	ParentId int64  `json:"parent_id"` // 父部门id
	DeptName string `json:"dept_name"` // 部门名称
	OrderNum int    `json:"order_num"` // 显示顺序
	Leader   string `json:"leader"`    // 负责人
	Phone    string `json:"phone"`     // 联系电话
	Email    string `json:"email"`     // 邮箱
	Status   int    `json:"status"`    // 状态 1,2
}

type DeptUpdateReq struct {
	g.Meta   `path:"/auth/dept/update" method:"post"`
	DeptId   int64  `json:"dept_id"`   // 部门id
	ParentId int64  `json:"parent_id"` // 父部门id
	DeptName string `json:"dept_name"` // 部门名称
	OrderNum int    `json:"order_num"` // 显示顺序
	Leader   string `json:"leader"`    // 负责人
	Phone    string `json:"phone"`     // 联系电话
	Email    string `json:"email"`     // 邮箱
	Status   int    `json:"status"`    // 部门状态（1正常 2停用）
}

type DeptDeleteReq struct {
	g.Meta `path:"/auth/dept/delete" method:"post"`
	DeptId int64 `json:"dept_id"` // 部门id
}
