package fileType

import (
	"github.com/gogf/gf/v2/os/gfile"
)

func ReadXsbj(path string) string {
	content := gfile.GetContents(path)
	return content
}

// 保存
func SaveXsbj(path string, content string) error {
	return gfile.PutContents(path, content)
}
