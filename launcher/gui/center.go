package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Center struct {
	Box *fyne.Container

	ApiLabel *widget.Label
	ApiEntry *widget.Entry
	PermissionsLabel *widget.Label
	PermissionsRadioGroup *widget.RadioGroup
	LocalFolderEntry *widget.Entry
	LocalFolderButton *widget.Button
	LocalFolderOpen *dialog.FileDialog

	gui *Gui
}

func NewCenter(g *Gui) *Center {
	n := &Center{
		ApiLabel: widget.NewLabel("服务地址："),
		ApiEntry: widget.NewEntryWithData(g.ApiStr),
		PermissionsLabel: widget.NewLabel("权限："),
		PermissionsRadioGroup: widget.NewRadioGroup([]string{"ReadOnly", "Writable"},func(value string) {
			g.Premissions.Set(value)
		}),
		LocalFolderEntry: widget.NewEntryWithData(g.RootPath),
		LocalFolderOpen:dialog.NewFolderOpen(func(f fyne.ListableURI,e error){
			g.RootPath.Set(f.Path())
		},g.W),
		gui: g,
	}
	n.LocalFolderButton = widget.NewButton("选择笔记本", func() {
		n.LocalFolderOpen.Show()
	})
	n.PermissionsRadioGroup.SetSelected("Writable")
	n.Box = container.New(layout.NewFormLayout(),
		n.ApiLabel, n.ApiEntry,
		n.PermissionsLabel, n.PermissionsRadioGroup,
		n.LocalFolderButton,n.LocalFolderEntry)
	return n
}
