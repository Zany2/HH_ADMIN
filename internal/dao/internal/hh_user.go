// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HhUserDao is the data access object for table hh_user.
type HhUserDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns HhUserColumns // columns contains all the column names of Table for convenient usage.
}

// HhUserColumns defines and stores column names for table hh_user.
type HhUserColumns struct {
	UserId     string // 用户ID
	UserName   string // 用户账号
	Password   string // 密码
	NickName   string // 用户昵称
	UserType   string // 用户类型 1超级管理员 2普通用户 默认2
	Email      string // 用户邮箱
	Phone      string // 手机号码
	Sex        string // 用户性别 1男 2女 3未知 默认1
	Avatar     string // 头像地址
	Status     string // 帐号状态 1正常 2停用 默认1
	DelFlag    string // 删除标志 1正常 2删除 默认1
	LoginIp    string // 最后登录IP
	CreateBy   string // 创建者 默认1
	UpdateBy   string // 更新者 默认1
	Remark     string // 备注
	LoginDate  string // 最后登录时间
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// hhUserColumns holds the columns for table hh_user.
var hhUserColumns = HhUserColumns{
	UserId:     "user_id",
	UserName:   "user_name",
	Password:   "password",
	NickName:   "nick_name",
	UserType:   "user_type",
	Email:      "email",
	Phone:      "phone",
	Sex:        "sex",
	Avatar:     "avatar",
	Status:     "status",
	DelFlag:    "del_flag",
	LoginIp:    "login_ip",
	CreateBy:   "create_by",
	UpdateBy:   "update_by",
	Remark:     "remark",
	LoginDate:  "login_date",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewHhUserDao creates and returns a new DAO object for table data access.
func NewHhUserDao() *HhUserDao {
	return &HhUserDao{
		group:   "default",
		table:   "hh_user",
		columns: hhUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HhUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HhUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HhUserDao) Columns() HhUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HhUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HhUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HhUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}