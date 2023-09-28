package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
)

type RoleMenu struct {
	gmeta.Meta `orm:"table:sys_role_menu"`
	RoleId     int64  `json:"role_id"`
	MenuId     int64  `json:"menu_id"`
	MenuLists  []Menu `orm:"with:menu_id=menu_id"`
}

type Menu struct {
	gmeta.Meta `orm:"table:sys_role_menu"`
	MenuId     int64  `json:"menu_id"`
	Perms      string `json:"perms"`
}
