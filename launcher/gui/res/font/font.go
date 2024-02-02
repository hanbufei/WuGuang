package font

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type HeiTiTheme struct{}

var _ fyne.Theme = (*HeiTiTheme)(nil)

// HTfont 对应的是 ttf.go 中的变量名
func (m HeiTiTheme) Font(fyne.TextStyle) fyne.Resource {
	return HTfont
}

func (*HeiTiTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (*HeiTiTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*HeiTiTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}