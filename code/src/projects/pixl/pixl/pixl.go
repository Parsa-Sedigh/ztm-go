package main

import (
	"fyne.io/fyne/v2/app"
	"image/color"
	"pixl/apptype"
	"pixl/swatch"
	"pixl/ui"
)

func main() {
	// create a new app:
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")

	// create an application state:
	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlWindow.ShowAndRun()
}
