package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"server/config"
	"server/lifeCyc"
	"server/router"
)

func main() {
	config.Init() //获取配置
	lifeCyc.BeforeStart()
	var ctx = gctx.New()
	cmd.Run(ctx)
}

var (
	cmd = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "WuGuang Notepad",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Printf("欢迎使用雾光笔记，浏览器打开如下网址即可：\nhttp://%s\n", config.Addr)
			SetDefaultHandler() // 替代默认的log
			s := g.Server()
			s.SetIndexFolder(true)
			s.SetServerRoot("dist")
			//设置图片外链的本地存放路径在当前用户目录下,并添加到本地静态服务
			s.AddStaticPath("/download", config.RunPath+gfile.Separator+"download")
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse) //默认的错误处理
				router.R.BindController(ctx, group)
			})
			s.SetAddr(config.Addr)
			s.Run()
			return nil
		},
	}
)

// 替代默认的日志handler，禁止控制台输出，全部输出到文件
func SetDefaultHandler() {
	glog.SetDefaultHandler(func(ctx context.Context, in *glog.HandlerInput) {
		m := map[string]interface{}{
			"stdout": false,
			//"path":   g.Config().MustGet(ctx, "logger.path", "log/").String(), // 此处必须重新设置，才可以实现写入文件
			//"path":config.RunPath,
		}
		in.Logger.SetConfigWithMap(m)
		in.Next(ctx)
	})
}
