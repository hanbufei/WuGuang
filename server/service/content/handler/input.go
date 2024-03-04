package handler

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"server/service/content/common"
	"server/service/content/fileType"
	"server/service/storage"
)

func HandlerInput(path string, content string) error {
	rootpath, err := storage.GetRootPath()
	if err != nil {
		return err
	}
	flag := common.GetExtName(gfile.ExtName(path))
	switch flag {
	case "html":
		return fileType.SaveHtml(rootpath+path, content)
	default:
		return gerror.New("仅允许保存 html 后缀的笔记")
	}
}
