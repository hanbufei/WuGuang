package config

import (
	"flag"
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"regexp"
)

var (
	Addr     string
	ReadOnly string
	RunPath  string
	RootPath string
)

// 从命令行获取接口地址
func Init() {
	flag.StringVar(&Addr, "addr", "127.0.0.1:8567", " 服务接口。请勿绑定互联网ip，本产品仅作为个人的本地笔记本，强烈不建议联通互联网，以免产生安全风险")
	flag.StringVar(&ReadOnly, "r", "false", "只读模式")
	flag.StringVar(&RunPath, "runpath", "", "程序运行目录")
	flag.StringVar(&RootPath, "rootpath", "~", "笔记根目录")
	flag.Parse()
}

// 将运行目录下的dist目录打包进可执行文件中
func AddStaticFile() {
	apiUrlJs := RunPath + gfile.Separator + "dist" + gfile.Separator + "p__index.async.js"
	content := gfile.GetContents(apiUrlJs)
	re := regexp.MustCompile(`apiUrl:"http://\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d{1,5}"`)
	content = re.ReplaceAllString(content, fmt.Sprintf(`apiUrl:"http://%s"`, Addr))
	gfile.PutContents(apiUrlJs, content)
	binContent, err := gres.Pack(RunPath + gfile.Separator + "dist")
	if err != nil {
		panic(err)
	}
	gres.Add(string(binContent))
}
