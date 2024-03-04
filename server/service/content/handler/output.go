package handler

import (
	"github.com/gogf/gf/v2/os/gfile"
	"server/service/content/common"
	"server/service/content/fileType"
	"server/service/storage"
)

func HandlerOut(path string) string {
	rootpath, err := storage.GetRootPath()
	if err != nil {
		return err.Error()
	}
	flag := common.GetExtName(gfile.ExtName(path))
	switch flag {
	case "txt":
		return fileType.ReadTxt(rootpath + path)
	case "md":
		return fileType.ReadMd(rootpath + path)
	case "html":
		return fileType.ReadHtml(rootpath + path)
	default:
		return "[error]选择的文件不是受支持的文件类型！\n" +
			"目前支持的类型如下：\n" +
			"txt\n" +
			"md\n" +
			"html"
	}
}
