package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"server/service/content"
	"server/service/menu"
	"server/service/storage"
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
