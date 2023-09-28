package logic

import (
	"HH_ADMIN/internal/model/entity"
)

// DeptParentIdList 根据deptId查找父id列表
func DeptParentIdList(deptList []entity.SysDept, dept entity.SysDept) []entity.SysDept {
	var parents []entity.SysDept

	for _, n := range deptList {
		if n.DeptId == dept.ParentId {
			// 找到父节点
			parents = append(parents, n)
			parents = append(parents, DeptParentIdList(deptList, n)...)
		}
	}

	return parents
}
