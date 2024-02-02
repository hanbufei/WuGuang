package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"server/service/content"
	"server/service/menu"
	"server/service/storage"
	"server/service/version"
)

func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"127.0.0.1", "localhost"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	//version 路由，绑定url和controller
	group.Group("/version", func(group *ghttp.RouterGroup) {
		group.Bind(
			version.New(),
		)
	})
	//storage 路由
	group.Group("/storage", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.Bind(
			storage.New(),
		)
	})
	//menu 路由
	group.Group("/menu", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.Bind(
			menu.New(),
		)
	})
	//content 路由
	group.Group("/content", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.Bind(
			content.New(),
		)
	})
}
