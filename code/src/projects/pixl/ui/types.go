package ui

import (
	"fyne.io/fyne/v2"
	"pixl/apptype"
	"pixl/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
