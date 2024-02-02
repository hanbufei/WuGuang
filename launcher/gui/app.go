package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	font2 "launcher/gui/res/font"
	"net/url"
)

type Gui struct {
	a fyne.App
	W fyne.Window
	MainPane *fyne.Container

	IsRunning  bool
	ApiUrl url.URL

	ApiStr binding.String
	RootPath binding.String
	Premissions binding.String //"ReadOnly", "Writable"
	Message binding.String
}

func NewGui() Gui {
	gui := Gui{
		a: app.NewWithID("com.github.hanbufei.wuguang"),
		IsRunning: false,
		ApiStr:binding.NewString(),
		RootPath: binding.NewString(),
		Premissions: binding.NewString(),
		Message: binding.NewString(),
		ApiUrl: url.URL{
			Scheme:"http",
			Host: "127.0.0.1:8567",
		},
	}
	gui.a.Settings().SetTheme(&font2.HeiTiTheme{})
	root := gui.a.Preferences().String("RootPath")
	if root == ""{
		gui.RootPath.Set("~")
	}else {
		gui.RootPath.Set(root)
	}
	gui.W = gui.a.NewWindow("雾光笔记")
	gui.ApiStr.Set("127.0.0.1:8567")
	gui.Premissions.Set("Writable")
	return gui
}

func (g *Gui)Init() {
	g.W.Resize(fyne.NewSize(400, 320))
	header  := NewHeader(g)
	center := NewCenter(g)
	footer := NewFooter(g)
	sider := NewSider(header,footer,g)
	g.MainPane = container.NewBorder(header.Box,footer.Box,nil,sider.Box,center.Box)
	g.W.SetContent(g.MainPane)
}

func (g *Gui)Start() {
	if desk, ok := g.a.(desktop.App); ok {
		g.W.SetCloseIntercept(func() {
			g.W.Hide()
		})
		m := fyne.NewMenu("雾光笔记",
			fyne.NewMenuItem("显示启动器", func() {
				g.W.Show()
			}),
			fyne.NewMenuItem("关闭并退出", func() {
				g.Exit()
			}),
		)
		desk.SetSystemTrayMenu(m)
	}
	g.W.SetOnClosed(g.Exit)
	g.W.ShowAndRun()
}

func (g *Gui)Exit() {
	err := StopService()
	if err != nil {
		g.Message.Set(err.Error())
	}
	rootpath ,err := g.RootPath.Get()
	if err != nil {
		rootpath = ""
	}
	g.a.Preferences().SetString("RootPath",rootpath)
	g.W.Close()
	g.a.Quit()
}
