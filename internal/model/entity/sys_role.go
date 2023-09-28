// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	RoleId            int64       `json:"role_id"             description:"角色ID"`
	RoleName          string      `json:"role_name"           description:"角色名称"`
	RoleKey           string      `json:"role_key"            description:"角色权限字符串"`
	RoleSort          int         `json:"role_sort"           description:"显示顺序"`
	DataScope         int         `json:"data_scope"          description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）默认2"`
	MenuCheckStrictly int         `json:"menu_check_strictly" description:"菜单树选择项是否关联显示"`
	DeptCheckStrictly int         `json:"dept_check_strictly" description:"部门树选择项是否关联显示"`
	Status            int         `json:"status"              description:"角色状态 1正常 2停用 默认1"`
	DelFlag           int         `json:"del_flag"            description:"删除标志 1正常 2停用 默认1"`
	CreateBy          string      `json:"create_by"           description:"创建者 默认0"`
	UpdateBy          string      `json:"update_by"           description:"更新者 默认0"`
	Remark            string      `json:"remark"              description:"备注"`
	CreateTime        *gtime.Time `json:"create_time"         description:"创建时间"`
	UpdateTime        *gtime.Time `json:"update_time"         description:"更新时间"`
}
