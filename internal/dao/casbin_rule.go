// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"HH_ADMIN/internal/dao/internal"
)

// internalCasbinRuleDao is internal type for wrapping internal DAO implements.
type internalCasbinRuleDao = *internal.CasbinRuleDao

// casbinRuleDao is the data access object for table casbin_rule.
// You can define custom methods on it to extend its functionality as you wish.
type casbinRuleDao struct {
	internalCasbinRuleDao
}

var (
	// CasbinRule is globally public accessible object for table casbin_rule operations.
	CasbinRule = casbinRuleDao{
		internal.NewCasbinRuleDao(),
	}
)

// Fill with you ideas below.
