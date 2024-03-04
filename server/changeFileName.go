package main

import (
	"github.com/gogf/gf/v2/os/gfile"
	"strings"
)

// 批量修改文件后缀：把 wg 改为 html
func main() {
	rootpath := "/Users/hanfei/wgbj"
	result, _ := gfile.ScanDir(rootpath, "*", true)
	for _, item := range result {
		newItem := strings.ReplaceAll(item, ".wg", ".html")
		gfile.Move(item, newItem)
	}
}
