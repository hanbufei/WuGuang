package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type Header struct {
	Box *fyne.Container

	label fyne.CanvasObject
	nillabel fyne.CanvasObject
	labelstr binding.String
	link fyne.CanvasObject
	gui *Gui
}

func NewHeader(g *Gui) *Header {
	n := &Header{
		gui:g,
		labelstr: binding.NewString(),
		nillabel: widget.NewLabel(""),
	}
	n.labelstr.Set("欢迎使用雾光笔记 (o^^o)")
	n.label = widget.NewLabelWithData(n.labelstr)
	n.Box = container.NewVBox(n.label,n.nillabel)
	return n
}

func (h *Header) UpdateToRunning()  {
	text,_ := h.gui.ApiStr.Get()
	h.gui.ApiUrl.Host = text
	h.labelstr.Set("雾光笔记已启动，直接访问下面的地址即可：")
	h.link = widget.NewHyperlink("http://"+text,&h.gui.ApiUrl)
	h.Box.Remove(h.nillabel)
	h.Box.Add(h.link)
}

func (h *Header) UpdateToStop()  {
	h.labelstr.Set("欢迎使用雾光笔记 (o^^o)")
	h.Box.Remove(h.link)
	h.Box.Add(h.nillabel)
}