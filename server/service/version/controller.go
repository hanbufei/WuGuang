package version

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Hello(ctx context.Context, req *Req) (res *Res, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("雾光笔记：基于本地文件系统的个人知识库")
	return
}
