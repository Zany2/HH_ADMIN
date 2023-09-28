package api

import (
	"HH_ADMIN/internal/model/entity"
	"HH_ADMIN/utility/rr"
	"github.com/gogf/gf/v2/frame/g"
)

type UserListReq struct {
	g.Meta `path:"/auth/user/list" method:"get"`
	rr.CommonReq
	DeptId  int64  `json:"dept_id"`                           // 部门id
	KeyWord string `json:"key_word"`                          // 关键字
	Phone   string `json:"phone"`                             // 手机号
	Stime   string `json:"stime" v:"datetime #时间格式错误"`        // 开始时间
	Etime   string `json:"etime" v:"datetime #时间格式错误"`        // 结束时间
	Status  int    `json:"status" v:"in:0,1,2 #状态必须在0,1,2之间"` // 状态 0全部 1正常 2停用 默认1
}
type UserListRes struct {
	SysUser entity.SysUser   `json:"sys_user"`
	SysDept entity.SysDept   `json:"sys_dept"`
	SysRole []entity.SysRole `json:"sys_role"`
	SysPost []entity.SysPost `json:"sys_post"`
}

type UserAddReq struct {
	g.Meta   `path:"/auth/user/add" method:"post"`
	UserName string  `json:"user_name" v:"required #用户名不能为空"` // 用户名
	Password string  `json:"password" v:"required #密码不能为空"`   // 密码
	Phone    string  `json:"phone" v:"phone #手机号码格式不正确"`      // 手机号码
	NickName string  `json:"nick_name"`                       // 用户昵称
	RoleId   []int64 `json:"role_id"`                         // 关联角色
	DeptId   int64   `json:"dept_id" v:"required #部门不能为空"`    // 关联部门
	Email    string  `json:"email"`                           // 用户邮箱
	Sex      int     `json:"sex" v:"in:1,2,3 #性别必须在1,2,3之间"`  // 性别 用户性别 1男 2女 3未知 默认1
	Status   int     `json:"status" v:"in:1,2 #状态必须在1,2之间"`   // 状态 1启用 2禁用
	PostId   []int64 `json:"post_id"`                         // 岗位
	Remark   string  `json:"remark"`                          // 描述
}

type UserUpdateReq struct {
	g.Meta   `path:"/auth/user/update" method:"post"`
	UserId   int64   `json:"user_id"`                      // 用户ID
	NickName string  `json:"nick_name"`                    // 用户昵称
	Phone    string  `json:"phone" v:"phone #手机号码格式不正确"`   // 手机号码
	RoleId   []int64 `json:"role_id"`                      // 关联角色
	DeptId   int64   `json:"dept_id" v:"required #部门不能为空"` // 关联部门
	Email    string  `json:"email"`                        // 用户邮箱
	Sex      int     `json:"sex"`                          // 用户性别 1男 2女 3未知 默认1
	Status   int     `json:"status"`                       // 状态 1启用 2禁用
	PostId   []int64 `json:"post_id"`                      // 岗位
	Remark   string  `json:"remark"`                       // 描述
}

type UserDeleteReq struct {
	g.Meta   `path:"/auth/user/delete" method:"post"`
	UserId   int64  `json:"user_id" v:"required #user_id不能为空"`
	EditType int    `json:"edit_type" v:"in:1,2 #类型必须在1,2之间"` // 1删除 2重置
	Password string `json:"password"`                         // 密码
}
