package apptype

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"image/color"
)

type BrushType = int

type PxCanvasConfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PxSize         int // represents the scale factor of the pixels
}

type State struct {
	BrushColor     color.Color
	BrushType      int
	SwatchSelected int // all of our swatches will be sorted in the slice and the SwatchSelected variable is just the index into that slice that the user is clicked on
	FilePath       string
}

func (state *State) SetFilePath(path string) {
	state.FilePath = path
}

type Brushable interface {
	SetColor(c color.Color, x, y int)

	// this function is gonna convert our mouse event, which will be our pointer location into our pixel x y coordinates within our image.
	MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int)
}
