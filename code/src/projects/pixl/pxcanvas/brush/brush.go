package brush

import (
	"fyne.io/fyne/v2/driver/desktop"
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
