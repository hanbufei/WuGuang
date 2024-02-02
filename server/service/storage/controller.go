package storage

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"strings"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// 获取已经配置的文件系统根路径
func (c *Controller) GetRootPathReq(ctx context.Context, req *GetRootPathReq) (res *GetRootPathRes, err error) {
	res = new(GetRootPathRes)
	//统一以"/"作为分割符
	rootpath, err := GetRootPath()
	if err != nil {
		return nil, err
	}
	res.RootPath = strings.ReplaceAll(rootpath, "\\", "/")
	return
}

// 配置文件系统根路径
func (c *Controller) SetRootPathReq(ctx context.Context, req *SetRootPathReq) (res *SetRootPathRes, err error) {
	res = new(SetRootPathRes)
	if strings.Contains(req.RootPath, "..") {
		err = gerror.New("路径含非法字符")
		return
	}
	if strings.HasSuffix(req.RootPath, "/") {
		err = gerror.New("路径结尾不能是/或\\")
		return
	}
	if strings.HasSuffix(req.RootPath, "\\") {
		err = gerror.New("路径结尾不能是/或\\")
		return
	}
	//统一以"/"作为分割符
	err = SetRootPath(req.RootPath)
	if err != nil {
		return nil, err
	}
	return
}
