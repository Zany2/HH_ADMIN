package util

import "HH_ADMIN/internal/model"

// GetTreeRecursive 递归拼接部门菜单
func GetTreeRecursive(list []*model.SysDeptList, parentId int64) []*model.SysDeptList {
	res := make([]*model.SysDeptList, 0)

	for _, v := range list {
		if v.ParentId == parentId {
			v.Children = GetTreeRecursive(list, v.DeptId)
			res = append(res, v)
		}
	}

	return res
}

// GetTreeRecursiveMenu 递归拼接菜单菜单
func GetTreeRecursiveMenu(list []*model.SysMenuS, parentId int64) []*model.SysMenuS {
	res := make([]*model.SysMenuS, 0)

	for _, v := range list {
		if v.ParentId == parentId {
			v.Children = GetTreeRecursiveMenu(list, v.MenuId)
			res = append(res, v)
		}
	}

	return res
}
