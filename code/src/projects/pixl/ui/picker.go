package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/lusingander/colorpicker"
	"image/color"
)

func SetupColorPicker(app *AppInit) *fyne.Container {
	picker := colorpicker.New(200, colorpicker.StyleHue)

	/* This function will be called when the color changes on the colorpicker. So when the user changes the colorpicker location or the oapcity,
	we're gonna call this function. */
	picker.SetOnChanged(func(c color.Color) {
		app.State.BrushColor = c
		app.Swatches[app.State.SwatchSelected].SetColor(c)
	})

	return container.NewVBox(picker)
}
