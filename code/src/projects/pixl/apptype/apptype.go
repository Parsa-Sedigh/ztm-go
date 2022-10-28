package apptype

import (
	"fyne.io/fyne/v2"
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
	SwatchSelected int // all of our swatches will be sorted in the slice and the SwatchSelected variable is just the index into that slice that the user is clicked on
	FilePath       string
}

func (state *State) SetFilePath(path string) {
	state.FilePath = path
}
