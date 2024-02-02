package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	res2 "launcher/gui/res"
)

type Sider struct {
	Box *fyne.Container

	RunButton *widget.Button
	ExitButton *widget.Button

	header *Header
	footer *Footer
	gui *Gui
}

func NewSider(header *Header,footer *Footer,g *Gui) *Sider {
	box := &Sider{
		header: header,
		footer: footer,
		gui: g,
	}
	box.RunButton = widget.NewButtonWithIcon("启动", res2.StartPng, func() {
		box.UpdateToRunning()
	})

	box.ExitButton = widget.NewButton("关闭程序", func() {
		g.Exit()
	})

	box.Box = container.NewVBox(box.RunButton,box.ExitButton)
	return box
}

func (s *Sider) UpdateToRunning()  {
	go StartSerice(s.gui)
	s.header.UpdateToRunning()
	s.footer.UpdateToRunning()

	s.RunButton.SetText("停止")
	s.RunButton.SetIcon(res2.StopPng)
	s.RunButton.OnTapped = func() {
		s.UpdateToStop()
	}
}

func (s *Sider) UpdateToStop()  {
	err := StopService()
	if err != nil {
		go  s.footer.AlertError(err)
	}
	s.header.UpdateToStop()
	s.footer.UpdateToStop()
	s.RunButton.SetText("启动")
	s.RunButton.SetIcon(res2.StartPng)
	s.RunButton.OnTapped = func() {
		s.UpdateToRunning()
	}
}



