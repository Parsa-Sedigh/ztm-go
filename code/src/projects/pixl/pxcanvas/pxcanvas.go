package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/color"
	"pixl/apptype"
)

type PxCanvasMouseState struct {
	/* previous coordinate of the mouse. We need to store the previous coordinate of the mouse and that's used for panning the image around, so
	we can calculate where the mouse was, versus where it is right now?*/
	previousCoord *fyne.PointEvent
}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer   *Px
	PixelData  image.Image // our actual pixel image data
	mouseState PxCanvasMouseState

	/* PxCanvas is working with the application state(*apptype.State) and that's because it needs access to the brush information in order to pain pixels. */
	appState *apptype.State

	/* reloadImage is used to reload whatever image happens to be stored in PixelData and it's used with new and open menus. */
	reloadImage bool
}

/*  Look at PxCanvas slide.

We're calcualting top left and bottom right corners of the canvas, in order to create a rectangle. Top left corener is x0 and y0 coordinates.*/
func (pxCanvas *PxCanvas) Bounds() image.Rectangle {
	x0 := int(pxCanvas.CanvasOffset.X)
	y0 := int(pxCanvas.CanvasOffset.Y)
	x1 := int(pxCanvas.PxCols*pxCanvas.PxSize + int(pxCanvas.CanvasOffset.X))
	y1 := int(pxCanvas.PxRows*pxCanvas.PxSize + int(pxCanvas.CanvasOffset.Y))

	return image.Rect(x0, y0, x1, y1)
}

/* a function that determines whether or not the mouse is within the bounding box.

pos means position of the mouse. */
func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= float32(bounds.Min.X) && pos.X < float32(bounds.Max.X) && pos.Y >= float32(bounds.Min.Y) && pos.Y < float32(bounds.Max.Y) {
		return true
	}

	return false
}

/* A function that creates a new blank image. That'll be when you open up the application, we just need some pixel data to work with.
We need a size(cols and rows) and an initial oclor.*/
func NewBlankImage(cols, rows int, c color.Color) image.Image {
	// we need to pass the size of the image to image.NewNRGBA()
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))

	// go through each pixel of the image and set the color. Start with rows and in each row, we're gonna fill out all the columns:
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			// set x and y coordinate to a specific color
			img.Set(x, y, c)
		}
	}

	return img
}

// a function to create a new pixel canvas:
func NewPxCanvas(state *apptype.State, config apptype.PxCanvasConfig) *PxCanvas {
	pxCanvas := &PxCanvas{
		PxCanvasConfig: config,
		appState:       state,
	}

	pxCanvas.PixelData = NewBlankImage(pxCanvas.PxCols, pxCanvas.PxRows, color.NRGBA{128, 128, 128, 255})
	pxCanvas.ExtendBaseWidget(pxCanvas)

	return pxCanvas
}

func (pxCanvas *PxCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &PxCanvasRenderer{
		pxCanvas:     pxCanvas,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
	}

	pxCanvas.renderer = renderer

	return renderer
}
