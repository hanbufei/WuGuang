package lifeCyc

import (
	"github.com/gogf/gf/v2/os/gfile"
	"log"
	"os"
	"path/filepath"
	"server/config"
)

func BeforeStart() {
	//storage.GetRecord()
	getRunPath()
	//在程序运行路径下，建立下载文件的本地存放目录
	gfile.Mkdir(config.RunPath + gfile.Separator + "download")
	config.AddStaticFile()
}

// 获取程序运行路径
func getRunPath() {
	if config.RunPath != "" {
		return
	}
	ex, err := os.Executable()
	if err != nil {
		log.Panic(err)
	}
	path := filepath.Dir(ex)
	config.RunPath = path
}
