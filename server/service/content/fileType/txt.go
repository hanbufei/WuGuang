package fileType

import (
	"github.com/gogf/gf/v2/os/gfile"
	"strings"
)

// 返回到前端时，将换行替换为p标签
func ReadTxt(path string) string {
	content := gfile.GetContents(path)
	output := ""
	for _, item := range strings.Split(content, "\n") {
		if item == "" {
			item = "<br>"
		}
		output = output + "<p>" + item + "</p>"
	}
	return output
}

//保存时，将</p><p>替换为换行，并删除首尾自动添加的p标签
//func SaveTxt(path string,content string) error {
//	content = strings.ReplaceAll(content,"</p><p>","\n")
//	content = strings.TrimPrefix(content,"<p>")
//	content = strings.TrimSuffix(content,"</p>")
//	content = strings.ReplaceAll(content,"<br>","\n")
//	fmt.Println(content)
//	return gfile.PutContents(path,content)
//}
