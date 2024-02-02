package fileType

import (
	"github.com/gogf/gf/v2/os/gfile"
)

func ReadMd(path string) string {
	content := gfile.GetContents(path)
	return content
}
