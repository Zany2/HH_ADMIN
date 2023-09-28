// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"HH_ADMIN/internal/dao/internal"
)

// internalSysRoleMenuDao is internal type for wrapping internal DAO implements.
type internalSysRoleMenuDao = *internal.SysRoleMenuDao

// sysRoleMenuDao is the data access object for table sys_role_menu.
// You can define custom methods on it to extend its functionality as you wish.
type sysRoleMenuDao struct {
	internalSysRoleMenuDao
}

var (
	// SysRoleMenu is globally public accessible object for table sys_role_menu operations.
	SysRoleMenu = sysRoleMenuDao{
		internal.NewSysRoleMenuDao(),
	}
)

// Fill with you ideas below.
