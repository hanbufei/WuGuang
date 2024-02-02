package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	res2 "launcher/gui/res"
)

type Footer struct {
	Box *fyne.Container

	icon *widget.Icon
	gui *Gui
}

func NewFooter(g *Gui) *Footer {
	box := &Footer{
		icon:widget.NewIcon(res2.DownPng),
		gui: g,
	}
	box.icon.Resize(fyne.NewSize(20,20))
	box.Hello()
	box.Box = container.NewVBox(widget.NewSeparator(),container.NewHBox(box.icon,widget.NewLabelWithData(box.gui.Message)))
	return box
}

func (h *Footer) Hello()  {
	h.gui.Message.Set("那些你记录下的，像刻在雾里的光，每一缕都相伴相随^_^")
}

func (h *Footer) AlertError(err error)  {
	h.gui.Message.Set(err.Error())
}

func (f *Footer) UpdateToRunning()  {
	f.icon.SetResource(res2.UpPng)
	f.Hello()
}

func (f *Footer) UpdateToStop()  {
	f.icon.SetResource(res2.DownPng)
}
