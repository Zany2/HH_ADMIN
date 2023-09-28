package model

import "github.com/gogf/gf/v2/os/gtime"

type SysMenuS struct {
	MenuId     int64       `json:"menuId"     description:"菜单ID"`
	MenuName   string      `json:"menuName"   description:"菜单名称"`
	ParentId   int64       `json:"parentId"   description:"父菜单ID 默认0"`
	OrderNum   int         `json:"orderNum"   description:"显示顺序 默认0"`
	Path       string      `json:"path"       description:"路由地址"`
	Component  string      `json:"component"  description:"组件路径"`
	Query      string      `json:"query"      description:"路由参数"`
	IsFrame    int         `json:"isFrame"    description:"是否为外链 1否 2是 默认1"`
	IsCache    int         `json:"isCache"    description:"是否缓存 1缓存 2不缓存 默认1"`
	MenuType   int         `json:"menuType"   description:"菜单类型（1目录 2菜单 3按钮）"`
	Visible    int         `json:"visible"    description:"菜单状态 1显示 2隐藏 默认1"`
	Status     int         `json:"status"     description:"菜单状态 1正常 2停用 默认1"`
	Perms      string      `json:"perms"      description:"权限标识"`
	Icon       string      `json:"icon"       description:"菜单图标 默认'#'"`
	CreateBy   int64       `json:"createBy"   description:"创建者 默认1"`
	UpdateBy   int64       `json:"updateBy"   description:"更新者 默认1"`
	Remark     string      `json:"remark"     description:"备注"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"更新时间"`
	Children   []*SysMenuS `json:"children"`
}
