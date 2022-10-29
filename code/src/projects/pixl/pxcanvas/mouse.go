package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

// implement the Scrollable interface by implement the Scrolled function:
func (pxCanvas *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.scale(int(ev.Scrolled.DY))
	pxCanvas.Refresh()
}

// some function implementations gonna be empty, because they have to be defined anyway for the interface to be implemented by us correctly:
func (pxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)
	pxCanvas.Refresh()

	/* We setting the previousCoord last, because we don't want to update it until our TryPan operation has completed. Otherwise, our TryPan operation will be
	operating on the same coordinates. So if we set the previousCoord before TryPan here, the TryPan it's gonna have the current coordinate(which we assigned to previousCoord) along with
	the current coordinate. So nothing will work!*/
	pxCanvas.mouseState.previousCoord = &ev.PointEvent
}
func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {}
func (pxCanvas *PxCanvas) MouseOut()                      {}
