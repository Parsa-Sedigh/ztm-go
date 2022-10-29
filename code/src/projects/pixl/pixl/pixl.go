package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"image/color"
	"pixl/apptype"
	"pixl/pxcanvas"
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

	pixlCanvasConfig := apptype.PxCanvasConfig{
		// the onscreen size of entire drawing area
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       10,
		PxCols:       10,
		/* The pixel canvas represents the scale factor. So for every 1 pixel in the image, it's gonna take up 30 pixels on screen. */
		PxSize: 30,
	}

	pixelCanvas := pxcanvas.NewPxCanvas(&state, pixlCanvasConfig)

	appInit := ui.AppInit{
		PixlCanvas: pixelCanvas,
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlWindow.ShowAndRun()
}
