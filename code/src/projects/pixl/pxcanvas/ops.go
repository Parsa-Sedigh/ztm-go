package pxcanvas

import "fyne.io/fyne/v2"

func (pxCanvas *PxCanvas) scale(direction int) {
	switch {
	case direction > 0:
		/* Increase the size of the pixels by 1 unit. Meaning, the onscreen size of the pixels will increase by 1.*/
		pxCanvas.PxSize += 1
	case direction < 0:
		if pxCanvas.PxSize > 2 {
			pxCanvas.PxSize -= 1
		}

	default:
		pxCanvas.PxSize = 10
	}
}

/* The PointEvent type, stores the position of the mouse pointer. */
func (pxCanvas *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDiff := currentCoord.Position.X - previousCoord.Position.X
	yDiff := currentCoord.Position.Y - previousCoord.Position.Y

	pxCanvas.CanvasOffset.X += xDiff
	pxCanvas.CanvasOffset.Y += yDiff

	pxCanvas.Refresh()
}
