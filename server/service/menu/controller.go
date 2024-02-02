package menu

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"server/service/storage"
	"strings"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// 获得初始化菜单数据
func (c *Controller) InitMenu(ctx context.Context, req *InitMenuReq) (res *InitMenuRes, err error) {
	res = new(InitMenuRes)
	rootpath, err := storage.GetRootPath()
	if err != nil {
		return nil, err
	}
	res.Label = getTitle(rootpath)
	result, _ := gfile.ScanDir(rootpath, "*", false)
	for _, item := range result {
		tmpdata := Menu{}
		key := strings.Replace(item, rootpath, "", 1)
		tmpdata.Key = strings.ReplaceAll(key, "\\", "/")
		tmpdata.Title = getTitle(key)
		//不显示隐藏文件
		if !(strings.HasPrefix(tmpdata.Title, ".")) {
			if gfile.IsFile(item) {
				tmpdata.IsLeaf = true
			} else {
				tmpdata.IsLeaf = false
			}
			res.MenuInitData = append(res.MenuInitData, tmpdata)
		}
	}
	return
}

// 获得指定路径菜单
func (c *Controller) ListMenu(ctx context.Context, req *ListMenuReq) (res *ListMenuRes, err error) {
	res = new(ListMenuRes)
	rootpath, err := storage.GetRootPath()
	if err != nil {
		return nil, err
	}
	if strings.Contains(req.Key, "..") {
		err = gerror.New("路径含非法字符")
		return
	}
	if gfile.IsFile(rootpath + req.Key) {
		err = gerror.New("选择的不是目录")
		return
	}
	if rootpath+req.Key == "/"{
		err = gerror.New("请设置笔记本路径")
		return
	}
	result, _ := gfile.ScanDir(rootpath+req.Key, "*", false)
	for _, item := range result {
		tmpdata := Menu{}
		key := strings.Replace(item, rootpath, "", 1)
		tmpdata.Key = strings.ReplaceAll(key, "\\", "/")
		tmpdata.Title = getTitle(key)
		//不显示隐藏文件
		if !(strings.HasPrefix(tmpdata.Title, ".")) {
			if gfile.IsFile(item) {
				tmpdata.IsLeaf = true
			} else {
				tmpdata.IsLeaf = false
			}
			res.MenuListData = append(res.MenuListData, tmpdata)
		}
	}
	return
}

func getTitle(key string) string {
	parts := strings.Split(key, "/")
	return parts[len(parts)-1]
}
