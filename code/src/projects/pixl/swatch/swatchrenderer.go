package swatch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

type SwatchRenderer struct {
	square  canvas.Rectangle
	objects []fyne.CanvasObject // objects that we need to draw

	/* this parent points to the swatch that created this renderer and we need this in order to determine whether or not sth selected
	in our refresh function.*/
	parent *Swatch
}

// implement the interface for a WidgetRenderer
func (renderer *SwatchRenderer) MinSize() fyne.Size {
	return renderer.square.MinSize()
}

/* this Layout functio ndetermines where in the layout that the swatch will be placed.
fyne.Size is the amount of space that we have available to draw our swatch in this case.*/
func (renderer *SwatchRenderer) Layout(size fyne.Size) {
	// since our swatch only consists of a single object within our CanvasObject slice, we can just access it directly:
	/* The resize() function is gonna resize our existing square to the maximum size that we have available for us. This will allow the layout to
	change the size of all those swatches at the same time. */
	renderer.objects[0].Resize(size)
}

func (renderer *SwatchRenderer) Refresh() {
	renderer.Layout(fyne.NewSize(20, 20))
	renderer.square.FillColor = renderer.parent.Color

	/* Whenever the Selected value is true, we're gonna set the StrokeWidth and StrokeColor. That will indicate the user that they have this current swatch selected. */
	if renderer.parent.Selected {
		renderer.square.StrokeWidth = 3
		renderer.square.StrokeColor = color.NRGBA{255, 255, 255, 255}

		/* We have to then reassign  the square into the objects slice. We have to do that because the objects are just CanvasObject interface type and those are
		types that don't have access to these specific square fields like StrokeWidth and ... .*/
		renderer.objects[0] = &renderer.square
	} else {
		renderer.square.StrokeWidth = 0
		renderer.objects[0] = &renderer.square
	}

	canvas.Refresh(renderer.parent)

}

func (renderer *SwatchRenderer) Objects() []fyne.CanvasObject {
	return renderer.objects
}

func (renderer *SwatchRenderer) Destroy() {}
