package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"pixl/pxcanvas/brush"
)

// implement the Scrollable interface by implement the Scrolled function:
func (pxCanvas *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.scale(int(ev.Scrolled.DY))
	pxCanvas.Refresh()
}

// some function implementations gonna be empty, because they have to be defined anyway for the interface to be implemented by us correctly:
// dispatches mouse events when appropriate
func (pxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	// is the mouse inside canvas?
	if x, y := pxCanvas.MouseToCanvasXY(ev); x != nil && y != nil {
		brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
	}
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)
	pxCanvas.Refresh()

	/* We setting the previousCoord last, because we don't want to update it until our TryPan operation has completed. Otherwise, our TryPan operation will be
	operating on the same coordinates. So if we set the previousCoord before TryPan here, the TryPan it's gonna have the current coordinate(which we assigned to previousCoord) along with
	the current coordinate. So nothing will work!*/
	pxCanvas.mouseState.previousCoord = &ev.PointEvent
}
func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {}
func (pxCanvas *PxCanvas) MouseOut()                      {}

func (pxCanvas *PxCanvas) MouseUp(ev *desktop.MouseEvent) {}
func (pxCanvas *PxCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
}
