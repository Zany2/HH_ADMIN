package test

import (
	"fmt"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"testing"
)

// MenuItem 表示菜单项
type MenuItem struct {
	ID       int
	Name     string
	ParentID int
	Children []*MenuItem
}

// MenuItem 表示菜单项
type MenuIte struct {
	ID       int
	Name     string
	ParentID int
}

// CreateMenuItems 模拟从数据库获取菜单项数据
func CreateMenuItems() []*MenuItem {
	menuItems := []*MenuItem{
		{1, "Root", 0, nil},
		{2, "File", 1, nil},
		{3, "Edit", 1, nil},
		{4, "New", 2, nil},
		{5, "Open", 2, nil},
		{6, "Cut", 3, nil},
		{7, "Copy", 3, nil},
		{8, "Submenu", 2, nil},
		{9, "Submenu Item 1", 8, nil},
		{10, "Submenu Item 2", 8, nil},
	}
	return menuItems
}

type TestStr struct {
	DeptId   int64      `json:"deptId"     description:"部门id"`
	ParentId int64      `json:"parentId"   description:"父部门id"`
	DeptName string     `json:"deptName"   description:"部门名称"`
	Leader   string     `json:"leader"     description:"负责人"`
	Phone    string     `json:"phone"      description:"联系电话"`
	Children []*TestStr `json:"children"`
}

// BuildMenuTree 构建菜单项树
func BuildMenuTree(menuItems []*TestStr, parentID int64) []*TestStr {
	var result []*TestStr

	for _, item := range menuItems {
		if item.ParentId == parentID {
			item.Children = BuildMenuTree(menuItems, item.DeptId)

			//if len(children) > 0 {
			//	item.Children = children
			//}

			result = append(result, item)
		}
	}

	return result
}

var (
	config = gredis.Config{
		Address: "127.0.0.1:6379",
		Db:      1,
		Pass:    "",
	}
	group = "cache"
	ctx   = gctx.New()
)

func Test(t *testing.T) {

	ite := MenuIte{
		ID:       1,
		Name:     "sdfsdf",
		ParentID: 35532,
	}

	var iiite MenuItem
	err := gconv.Struct(ite, &iiite)
	if err != nil {
		panic(err)
	}
	fmt.Println(iiite)

	//gredis.SetConfig(&config, group)
	//depts := make([]entity.SysDept, 0)
	//deptsRedisCache, err := g.Redis(group).Get(ctx, consts.CACHEDEPTHM)
	//if err != nil {
	//	panic(err)
	//}
	//err = gconv.Struct(deptsRedisCache, &depts)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, dept := range depts {
	//	g.Dump(dept)
	//}

	//depts := make([]entity.SysDept, 0)
	//deptsRedisCache, err := g.Redis().Get(gctx.New(), consts.CACHEDEPTHM)
	//if err != nil {
	//	panic(err)
	//}
	//err = gconv.Struct(deptsRedisCache, &depts)
	//if err != nil {
	//	panic(err)
	//}

	//strs := []*TestStr{
	//	{1, 0, "AAA", "AAA", "13888888888", nil},
	//	{2, 1, "BBB", "BBB", "13888888889", nil},
	//	{3, 1, "CCC", "CCC", "13888888884", nil},
	//	{4, 2, "DDD", "DDD", "13888888882", nil},
	//	{5, 4, "EEE", "EEE", "13888888884", nil},
	//}
	//
	//tree := BuildMenuTree(strs, 0)
	//marshal, _ := json.Marshal(tree)
	//fmt.Println(string(marshal))

	//menuItems := CreateMenuItems()

	//for _, item := range menuItems {
	//	fmt.Println(item)
	//}
	//menuTree := BuildMenuTree(menuItems, 0)
	//
	//// 打印菜单树
	//printMenu(menuTree, 0)
	//marshal, _ := json.Marshal(menuTree)
	//fmt.Println(string(marshal))
}

type Tng struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Lng struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Child []*Lng `json:"child"`
}

func printMenu(menu []*MenuItem, level int) {
	for _, item := range menu {
		fmt.Printf("%s%s\n", getIndent(level), item.Name)
		if len(item.Children) > 0 {
			printMenu(item.Children, level+1)
		}
	}
}

func getIndent(level int) string {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	return indent
}
