package content

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"server/config"
	"server/service/content/common"
	"server/service/content/handler"
	"server/service/storage"
	"strings"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// 获取文件内容
func (c *Controller) GetContentReq(ctx context.Context, req *GetContentReq) (res *GetContentRes, err error) {
	res = new(GetContentRes)
	if strings.Contains(req.Key, "..") {
		err = gerror.New("路径含非法字符")
		return
	}
	rootpath, err := storage.GetRootPath()
	if err != nil {
		return nil, err
	}
	key := strings.ReplaceAll(req.Key, "/", gfile.Separator)
	if !gfile.IsFile(rootpath + key) {
		err = gerror.New("选择的不是笔记文件")
		return
	}
	res.Content = handler.HandlerOut(key)
	return
}

// 新增文件
func (c *Controller) AddReq(ctx context.Context, req *AddReq) (res *AddRes, err error) {
	if config.ReadOnly == "ReadOnly" {
		err = gerror.New("当前为只读模式")
		return
	}
	if strings.Contains(req.Key, "..") {
		err = gerror.New("路径含非法字符")
		return
	}
	rootpath, err := storage.GetRootPath()
	if err != nil {
		return nil, err
	}
	key := strings.ReplaceAll(req.Key, "/", gfile.Separator)
	//新增目录
	if req.Type == "d" {
		if strings.Contains(key, ".") {
			err = gerror.New("非法目录名")
			return
		}
		err = gfile.Mkdir(rootpath + key)
		return
	}
	// 新增文件
	flag := common.GetExtName(gfile.ExtName(key))
	switch flag {
	case "html":
		_, err = gfile.Create(rootpath + key)
	default:
		err = gerror.New("仅允许新增 html 后缀的笔记")
	}
	return
}

// 删除文件
func (c *Controller) DeleteReq(ctx context.Context, req *DeleteReq) (res *DeleteRes, err error) {
	if config.ReadOnly == "ReadOnly" {
		err = gerror.New("当前为只读模式")
		return
	}
	if strings.Contains(req.Key, "..") {
		err = gerror.New("路径含非法字符")
		return
	}
	if req.Key == "/" {
		err = gerror.New("不可直接删除全部笔记")
		return
	}
	rootpath, err := storage.GetRootPath()
	if err != nil {
		return nil, err
	}
	key := strings.ReplaceAll(req.Key, "/", gfile.Separator)
	//删除目录
	if !gfile.IsFile(rootpath + key) {
		err = gfile.Remove(rootpath + key)
		return
	}
	// 删除文件
	flag := common.GetExtName(gfile.ExtName(key))
	switch flag {
	case "html":
		err = gfile.Remove(rootpath + key)
	default:
		err = gerror.New("仅允许删除 html 后缀的笔记")
	}
	return
}

// 保存内容到文件，如果文件不存在则新建
func (c *Controller) SaveContentReq(ctx context.Context, req *SaveContentReq) (res *SaveContentRes, err error) {
	if config.ReadOnly == "ReadOnly" {
		err = gerror.New("当前为只读模式")
		return
	}
	if strings.Contains(req.Key, "..") {
		err = gerror.New("路径含非法字符")
		return
	}
	rootpath, err := storage.GetRootPath()
	if err != nil {
		return nil, err
	}
	key := strings.ReplaceAll(req.Key, "/", gfile.Separator)
	if !gfile.IsFile(rootpath + key) {
		err = gerror.New("选择的不是笔记文件")
		return
	}
	err = handler.HandlerInput(key, req.Content)
	return
}

// 下载图片
func (c *Controller) DownloadImgReq(ctx context.Context, req *DownloadImgReq) (res *DownloadImgRes, err error) {
	if config.ReadOnly == "ReadOnly" {
		err = gerror.New("当前为只读模式")
		return
	}
	rootpath, err := storage.GetRootPath()
	if err != nil {
		return nil, err
	}
	imgPath := rootpath + gfile.Separator + "download" + gfile.Separator + gbase64.EncodeString(req.ImgUrl)
	resp, err := g.Client().Get(ctx, req.ImgUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	err = gfile.PutBytes(imgPath, resp.ReadAll())
	return
}
