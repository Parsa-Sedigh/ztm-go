package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	// a reference to the parent pxCanvas
	pxCanvas    *PxCanvas
	canvasImage *canvas.Image

	// we're gonna use a SLICE of lines to draw a border around our canvas
	canvasBorder []canvas.Line
	canvasCursor []fyne.CanvasObject
}

// WidgetRenderer interface implementation
func (renderer *PxCanvasRenderer) MinSize() fyne.Size {
	// DrawingArea represents the entire widget's size. We want that entire size, that way we can display our canvas within it on screen
	return renderer.pxCanvas.DrawingArea
}

// WidgetRenderer interface implementation
func (renderer *PxCanvasRenderer) Objects() []fyne.CanvasObject {
	// set the initial amount of objects to 0 and with a capacity of 5
	objects := make([]fyne.CanvasObject, 0, 5)

	for i := 0; i < len(renderer.canvasBorder); i++ {
		/* The capacity of objects variable is 5 and 4 of those objects will be our borderlines and the fifth object will be our actual image itself*/
		objects = append(objects, &renderer.canvasBorder[i])
	}

	// append the image
	objects = append(objects, renderer.canvasImage)
	objects = append(objects, renderer.canvasCursor...)

	return objects
}

func (renderer *PxCanvasRenderer) Destroy() {}

/* We're gonna separate the Layout function into 3 separate functions:
Layout function whcih is required for implementation, a function for laying out the border and for laying out the image.

It's important to call LayoutCanvas() first, because that's going to be resizing the image and the size of the image needs to be accurate in order to
layout the border around the image itself.*/
func (renderer *PxCanvasRenderer) Layout(size fyne.Size) {
	renderer.LayoutCanvas(size)
	renderer.LayoutBorder(size)
}

func (renderer *PxCanvasRenderer) Refresh() {
	if renderer.pxCanvas.reloadImage {
		renderer.canvasImage = canvas.NewImageFromImage(renderer.pxCanvas.PixelData)

		/* What does ScaleMode does, is it changes how the images scale so we can have different filtering to make it smooth or ... . We want to have
		actual pixels scaling, so we use ImageScalePixels and that would give us pixel perfect scaling and that's perfect since we're deign pixel art. */
		renderer.canvasImage.ScaleMode = canvas.ImageScalePixels

		/* When we use ImageFillContain as our FillMode, that allows the image to be contained within the size that we specify. Since we're resizing the image
		in LayoutCanvas func, when we use ImageFillContain, it's not going to go outside of that size because the default behavior would be to try to
		fill the layout.*/
		renderer.canvasImage.FillMode = canvas.ImageFillContain

		// once the image has been reloaded, we can change the flag to false, that way on the next refresh, we don't try to reload it again
		renderer.pxCanvas.reloadImage = false
	}

	// renderer.pxCanvas.Size() gonna be our total drawing area. That way we have a plenty of room to move our pixelCanvas around.
	renderer.Layout(renderer.pxCanvas.Size())

	// make sure everything on the screen is updated
	canvas.Refresh(renderer.canvasImage)
}

func (renderer *PxCanvasRenderer) LayoutCanvas(size fyne.Size) {
	// these are the actual pixel width and pixel height of the pixelData within our image and not the actual on the screen representation
	imgPxWidth := renderer.pxCanvas.PxCols
	imgPxHeight := renderer.pxCanvas.PxRows
	pxSize := renderer.pxCanvas.PxSize

	/* the canvasImage by default is gonna be at 0,0 coordinates. But we want it to be at virtual position(see pxCanvas slide). That way, we can pan
	the image around. So this line is gonna move our pixelCanvas from 0,0 coordinates to wherever it should be based on the panning.*/
	renderer.canvasImage.Move(fyne.NewPos(renderer.pxCanvas.CanvasOffset.X, renderer.pxCanvas.CanvasOffset.Y))

	/* Now we need to take care of size of our image. We need to multiply the pixeels by the actual pixel size that we want to SEE onscreen, that way we're able
	to more easily edit the pixel art.

	For example `imgPxHeight*pxSize` is the size in the y direction.*/
	renderer.canvasImage.Resize(fyne.NewSize(float32(imgPxWidth*pxSize), float32(imgPxHeight*pxSize)))
}
func (renderer *PxCanvasRenderer) LayoutBorder(size fyne.Size) {
	offset := renderer.pxCanvas.CanvasOffset

	/* Remember that in LayoutCanvas function, we've resized our image, so these are the actual onscreen size. So when we draw our borders, they will surround
	the canvas itself.*/
	imgHeight := renderer.canvasImage.Size().Height
	imgWidth := renderer.canvasImage.Size().Width

	left := &renderer.canvasBorder[0]

	/* The canvas.Line type has two positions, Position1 and Position2. They represents the start and end positions of the line. */

	/* The left border has the same x offset, but different in y. */
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imgHeight)

	top := &renderer.canvasBorder[1]
	top.Position1 = fyne.NewPos(offset.X, offset.Y)
	top.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y)

	right := &renderer.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X+imgWidth, offset.Y)
	right.Position2 = fyne.NewPos(offset.X, offset.Y+imgHeight)

	// for bottom border, we start from bottom left to bottom right
	bottom := &renderer.canvasBorder[1]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y+imgHeight)
	bottom.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)
}

func (renderer *PxCanvasRenderer) SetCursor(objects []fyne.CanvasObject) {
	renderer.canvasCursor = objects
}
