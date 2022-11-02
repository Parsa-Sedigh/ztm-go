package brush

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"image/color"
	"pixl/apptype"
)

const (
	Pixel = iota
)

/* The TryBrush function is similar to our TryPan function. Our mouse package will be calling the TryBrush() with the apprpriate info and it
it will up to this function to decide whether or not to paint a pixel.*/
func TryBrush(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	switch {
	case appState.BrushType == Pixel:
		return TryPaintPixel(appState, canvas, ev)
	default:
		return false //indicating the paint didn't succeed
	}
}

// the return value indicates whether or not it succeeded
func TryPaintPixel(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	x, y := canvas.MouseToCanvasXY(ev)

	// if we have x and y coordinates and user is pressing down his left mouse button:
	if x != nil && y != nil && ev.Button == desktop.MouseButtonPrimary {
		canvas.SetColor(appState.BrushColor, *x, *y)

		// to indicate operation was succeeded
		return true
	}

	return false
}

/* x and y params refer to the pixel coordinates in the actual image itself, not our virtual coordinates on screen.

The return type will be sent directly into our renderer and so we'll be able to make lines in this function and then the renderer will render those lines as a cursor.*/
func Cursor(config apptype.PxCanvasConfig, brush apptype.BrushType, ev *desktop.MouseEvent, x, y int) []fyne.CanvasObject {
	var objects []fyne.CanvasObject
	switch {
	case brush == Pixel:
		pxSize := float32(config.PxSize)

		/* The xOrigin and yOrigin represent the upper left corner of one of our virtual pixels onscreen and we need to calculate the upper left corners,
		that way we can create borders around our pixel and we're gonna be using the same logic that we did with the canvas borders.*/
		xOrigin := (float32(x) * pxSize) + config.CanvasOffset.X
		yOrigin := (float32(y) * pxSize) + config.CanvasOffset.Y

		cursorColor := color.NRGBA{80, 80, 80, 255}

		left := canvas.NewLine(cursorColor)
		left.StrokeWidth = 3
		left.Position1 = fyne.NewPos(xOrigin, yOrigin)
		left.Position2 = fyne.NewPos(xOrigin, yOrigin+pxSize)

		top := canvas.NewLine(cursorColor)
		top.StrokeWidth = 3
		top.Position1 = fyne.NewPos(xOrigin, yOrigin)
		top.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin)

		right := canvas.NewLine(cursorColor)
		right.StrokeWidth = 3
		right.Position1 = fyne.NewPos(xOrigin+pxSize, yOrigin)
		right.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin+pxSize)

		bottom := canvas.NewLine(cursorColor)
		bottom.StrokeWidth = 3
		bottom.Position1 = fyne.NewPos(xOrigin, yOrigin+pxSize)
		bottom.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin+pxSize)

		objects = append(objects, left, top, right, bottom)
	}

	return objects
}
