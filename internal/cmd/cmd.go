package cmd

import (
	"HH_ADMIN/casbinhh"
	"HH_ADMIN/internal/controller"
	"HH_ADMIN/internal/service"
	"HH_ADMIN/task"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "HH_ADMIN",
		Usage: "HH_ADMIN",
		Brief: "HH_ADMIN",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 缓存lua脚本
			task.InitLua(ctx)
			// 缓存部门、岗位、角色信息
			err = task.CacheRPD(ctx)
			if err != nil {
				panic(err)
			}
			// gtoken启动
			gToken := StartBackendGToken(ctx)
			// casbin启动
			casbinhh.InitCasbin(ctx)
			// 静态目录
			s.AddStaticPath("/static", g.Cfg().MustGet(ctx, "file_upload.location").String())

			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				// 跨域
				group.Middleware(service.Middleware.CORSMiddleware)
				// 统一返回
				group.Middleware(service.Middleware.HandlerResponseMiddleware)
				// gtoken
				err = gToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}

				// hello测试
				group.Bind(
					controller.Hello,
				)

				// system
				group.Group("/system", func(group *ghttp.RouterGroup) {
					group.Bind(
						controller.User,
						controller.File,
						controller.Role,
						controller.Dept,
						controller.Post,
						controller.Menu,
						controller.Login,
						controller.Server,
					)
				})
			})

			s.Run()
			return nil
		},
	}
)
