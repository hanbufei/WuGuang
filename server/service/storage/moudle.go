package storage

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"server/config"
	"strings"
)

// 所有的路径，都以   /或 \  开头
// 所有的路径，都不以 /或 \  结尾
var (
//XmlFile  = gfile.SelfDir() + gfile.Separator + "xiangsui.xml"
//rootPath string
)

func GetRootPath() (string, error) {
	tmp := formatPath(config.RootPath)
	//err := checkRootPath(tmp)
	//if err != nil {
	//	return "", err
	//}
	return tmp, nil
}

func SetRootPath(path string) error {
	tmp := formatPath(path)
	config.RootPath = tmp
	config.TmpRoot = strings.ReplaceAll(config.RootPath, "\\", "/")
	return nil
}

//// 程序启动前，获取本地文件中的笔记本路径
//func GetRecord() {
//	txt := gfile.GetContents(XmlFile)
//	j := gjson.New(txt)
//	rootPath = j.Get("record.RootPath").String()
//}

//// 保存到本地xml文件
//func SetRecord() {
//	txt := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
//	<record>
//		<RootPath>%s</RootPath>
//	</record>`, rootPath)
//	gfile.PutContents(XmlFile, txt)
//}

// 删除路径结尾的 / 和 \
func formatPath(path string) string {
	tmp := strings.TrimRight(path, "/")
	tmp = strings.TrimRight(tmp, "\\")
	return tmp
}

func checkRootPath(path string) (err error) {
	if path == "/" {
		err = gerror.New("linux的根路径不可设置为笔记本路径")
		return
	}
	if strings.HasPrefix(path, "c:") || strings.HasPrefix(path, "C:") {
		tmp := strings.Split(path, "\\")
		if len(tmp) < 3 {
			err = gerror.New("c盘下的一级目录不可设置为笔记本路径")
			return
		}
	}
	return
}
