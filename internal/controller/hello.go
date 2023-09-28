package controller

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/internal/consts"
	"HH_ADMIN/internal/model/entity"
	"HH_ADMIN/utility/rr"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) HelloTest(ctx context.Context, req *api.Req) (res *rr.CommonRes, err error) {

	roleRedisCache, err := g.Redis().HMGet(ctx, consts.CACHEROLEHM, "1")
	if err != nil {
		panic(err)
	}
	var role entity.SysRole
	fmt.Println(roleRedisCache[0])
	err = gconv.Struct(roleRedisCache[0], &role)
	if err != nil {
		panic(err)
	}
	fmt.Println(role)

	//get, err := g.Redis().HMGet(ctx, consts.CACHEMENUHM, "1")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(get)
	//fmt.Println("----------------------------------------------------")
	//for _, v := range get {
	//	fmt.Println(v)
	//}
	//fmt.Println("----------------------------------------------------")
	//
	//var menu entity.SysMenu
	//err = gconv.Struct(get[0], &menu)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(menu)
	//fmt.Println("----------------------------------------------------")

	//userId := ctx.Value(consts.CTXUSERID)
	//phone := ctx.Value(consts.CTXPHONE)

	//g.Dump(userId)
	//g.Dump(phone)

	//g.Log().Line().Infof(ctx, "Infof哈哈哈 %s", "ddsfsdfsd")
	//g.Log().Line().Warningf(ctx, "Warningf哈哈哈 %s", "ddsfsdfsd")

	//depts := make([]entity.SysDept, 0)
	//deptsRedisCache, err := g.Redis().Get(ctx, consts.CACHEDEPT)
	//if err != nil {
	//	g.Log().Line().Errorf(ctx, "DeptAdd err:%s", err.Error())
	//	return rr.FailedWithMessage("新增失败"), err
	//}
	//err = gconv.Struct(deptsRedisCache, &depts)
	//if err != nil {
	//	g.Log().Line().Errorf(ctx, "DeptAdd err:%s", err.Error())
	//	return rr.FailedWithMessage("新增失败"), err
	//}
	//
	//var dept entity.SysDept
	//for i := 0; i < len(depts); i++ {
	//	if depts[i].DeptId == 4 {
	//		dept = depts[i]
	//	}
	//}
	//
	//g.Dump("列表为：")
	//fmt.Println(depts)
	//g.Dump("查询的单个数据为：")
	//fmt.Println(dept)
	//
	//list := logic.DeptParentIdList(depts, dept)
	//g.Dump("查询出来的父级列表为：")
	//fmt.Println(list)
	//
	//ids := make([]int64, 0)
	//ids = append(ids, 0)
	//ids = append(ids, dept.DeptId)
	//for _, sysDept := range list {
	//	ids = append(ids, sysDept.DeptId)
	//}
	//sort.Slice(ids, func(i, j int) bool {
	//	return ids[i] < ids[j]
	//})
	//
	//fmt.Println("ids数据为：")
	//fmt.Println(ids)

	return rr.Success(), err
}
