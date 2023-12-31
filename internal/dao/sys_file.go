// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"HH_ADMIN/internal/dao/internal"
)

// internalSysFileDao is internal type for wrapping internal DAO implements.
type internalSysFileDao = *internal.SysFileDao

// sysFileDao is the data access object for table sys_file.
// You can define custom methods on it to extend its functionality as you wish.
type sysFileDao struct {
	internalSysFileDao
}

var (
	// SysFile is globally public accessible object for table sys_file operations.
	SysFile = sysFileDao{
		internal.NewSysFileDao(),
	}
)

// Fill with you ideas below.
