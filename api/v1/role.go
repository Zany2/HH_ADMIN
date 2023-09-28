package api

import (
	"HH_ADMIN/internal/model"
	"HH_ADMIN/internal/model/entity"
	"HH_ADMIN/utility/rr"
	"github.com/gogf/gf/v2/frame/g"
)

type RoleListReq struct {
	g.Meta `path:"/auth/role/list" method:"get"`
	rr.CommonReq
	RoleName string `json:"role_name"`
	Status   int    `json:"status" v:"in:1,2 #状态必须在0,1,2之间"`
}
type RoleListRes struct {
	Role     entity.SysRole    `json:"role"`
	MenuList []*model.SysMenuS `json:"menu_list"`
}

type RoleAddReq struct {
	g.Meta   `path:"/auth/role/add" method:"post"`
	RoleName string `json:"role_name"` // 角色名称
	RoleKey  string `json:"role_key"`  // 角色权限字符串
	MenuIds  []int  `json:"menu_ids"`  // menu列表
	RoleSort int    `json:"role_sort"` // 显示顺序
	Status   int    `json:"status"`    // 角色状态 1正常 2停用 默认1
	Remark   string `json:"remark"`    // 备注
}

type RoleUpdateReq struct {
	g.Meta   `path:"/auth/role/update" method:"post"`
	RoleId   int64  `json:"role_id"`   // 角色id
	RoleName string `json:"role_name"` // 角色名称
	RoleKey  string `json:"role_key"`  // 角色权限字符串
	MenuIds  []int  `json:"menu_ids"`  // menu列表
	RoleSort int    `json:"role_sort"` // 显示顺序
	Status   int    `json:"status"`    // 角色状态 1正常 2停用 默认1
	Remark   string `json:"remark"`    // 备注
}

type RoleDeleteReq struct {
	g.Meta `path:"/auth/role/delete" method:"post"`
	RoleId int64 `json:"role_id"` // 角色id
}

//type RoleQueryReq struct {
//	g.Meta `path:"/auth/role/query" method:"get"`
//	RoleId int64 `json:"role_id"`
//}
//type RoleQueryRes struct {
//	A string `json:"a"`
//	B int    `json:"b"`
//}
