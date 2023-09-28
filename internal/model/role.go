package model

import "HH_ADMIN/internal/model/entity"

type RoleQueryResult struct {
	Role entity.SysRole   `json:"role"`
	Menu []entity.SysUser `json:"menu"`
}
