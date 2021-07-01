package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

//go:embed fonts/msyhbd.ttc
var FontBold []byte

//go:embed fonts/msyhl.ttc
var FontLight []byte

//go:embed fonts/msyh.ttc
var FontRegular []byte

type CnTheme struct {}

var _ fyne.Theme = (*CnTheme)(nil)

func (r CnTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	//fmt.Println(name)
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}
	if name == theme.ColorNameButton {
		return color.White
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (r CnTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (r CnTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Monospace {
		return fyne.NewStaticResource("regular.ttc",FontRegular)
	}
	if style.Bold {
		return fyne.NewStaticResource("bold.ttc",FontBold)
	}
	return fyne.NewStaticResource("regular.ttc",FontRegular)
}

func (r CnTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}