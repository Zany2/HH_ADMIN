package api

import (
	"HH_ADMIN/utility/rr"
	"github.com/gogf/gf/v2/frame/g"
)

type MenuListReq struct {
	g.Meta `path:"/auth/menu/list" method:"get"`
	rr.CommonReq
	MenuName string `json:"menu_name"`
	Perms    string `json:"perms"`
}

type MenuAddReq struct {
	g.Meta    `path:"/auth/menu/add" method:"post"`
	MenuName  string  `json:"menu_name"` // 菜单名称
	ParentId  int64   `json:"parent_id"` // 父菜单ID 默认0
	OrderNum  int     `json:"order_num"` // 显示顺序 默认0
	Path      string  `json:"path"`      // 路由地址
	RoleId    []int64 `json:"role_id"`   // 角色关联id
	Method    string  `json:"method"`    // 方法
	Component string  `json:"component"   description:"组件路径"`
	IsFrame   int     `json:"is_frame"    description:"是否为外链 1否 2是 默认1"`
	IsCache   int     `json:"is_cache"    description:"是否缓存 1缓存 2不缓存 默认1"`
	MenuType  int     `json:"menu_type"   description:"菜单类型（1目录 2菜单 3按钮）"`
	Visible   int     `json:"visible"     description:"菜单状态 1显示 2隐藏 默认1"`
	Status    int     `json:"status"      description:"菜单状态 1正常 2停用 默认1"`
	Perms     string  `json:"perms"       description:"权限标识"`
	Icon      string  `json:"icon"        description:"菜单图标 默认'#'"`
	Remark    string  `json:"remark"      description:"备注"`
}

type MenuUpdateReq struct {
	g.Meta    `path:"/auth/menu/update" method:"post"`
	MenuId    int64  `json:"menu_id"`   // 菜单ID
	MenuName  string `json:"menu_name"` // 菜单名称
	ParentId  int64  `json:"parent_id"` // 父菜单ID 默认0
	OrderNum  int    `json:"order_num"` // 显示顺序 默认0
	Path      string `json:"path"`      // 路由地址
	RoleId    []int  `json:"role_id"`   // 角色关联id
	Method    string `json:"method"`    // 方法
	Component string `json:"component"   description:"组件路径"`
	IsFrame   int    `json:"is_frame"    description:"是否为外链 1否 2是 默认1"`
	IsCache   int    `json:"is_cache"    description:"是否缓存 1缓存 2不缓存 默认1"`
	MenuType  int    `json:"menu_type"   description:"菜单类型（1目录 2菜单 3按钮）"`
	Visible   int    `json:"visible"     description:"菜单状态 1显示 2隐藏 默认1"`
	Status    int    `json:"status"      description:"菜单状态 1正常 2停用 默认1"`
	Perms     string `json:"perms"       description:"权限标识"`
	Icon      string `json:"icon"        description:"菜单图标 默认'#'"`
	Remark    string `json:"remark"      description:"备注"`
}

type MenuDeleteReq struct {
	g.Meta `path:"/auth/menu/delete" method:"post"`
	MenuId int64 `json:"menu_id"` // 菜单ID
}
