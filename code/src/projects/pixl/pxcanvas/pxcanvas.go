package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
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
	renderer   *PxCanvasRenderer
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

/* We name this function TryPan, because our mouse functionality that we'll be creating soon will attempt to pan the image but it may not always succeed, because
the canvas may not want to do panning. For examplem you may have to hold hotkeys or setup some different configuration that will allow the
canvas to control whether or not panning occurs.

When we do panning, we have to have a previous coordinate. So when we first try panning sth, nothing will happen and then once we move the mouse at least 1 pixel,
the previous coordinate will be populated and so panning will work.

In the fyne toolkit, the teriary mouse button represents the scroll wheel.*/
func (pxCanvas *PxCanvas) TryPan(previousCoord *fyne.PointEvent, ev *desktop.MouseEvent) {
	/* Make sure that the mouse button is the middle mouse buttons(we pan by holding scroll wheel on the mouse). */
	if previousCoord != nil && ev.Button == desktop.MouseButtonTertiary {
		pxCanvas.Pan(*previousCoord, ev.PointEvent)
	}
}

// implement Brushable interface
func (pxCanvas *PxCanvas) SetColor(c color.Color, x, y int) {
	/* When the user loads a file, it may be in NRGBA or RGBA format. But regardless of the format, we're gonna `set` the pixel color. */

	/* Our PixelData is stored in an image.Image and that's an interface and it doesn't allow us to actually `set` the image data. So here, we're accessing the underlying
	type and we're checking to see if it's an NRGBA type? and if it is, then we have access to the Set function and we can set the x and y coordinate to a specific color(c variable)*/
	if nrgba, ok := pxCanvas.PixelData.(*image.NRGBA); ok {
		nrgba.Set(x, y, c)
	}

	if rgba, ok := pxCanvas.PixelData.(*image.RGBA); ok {
		rgba.Set(x, y, c)
	}

	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int) {
	// calculate the bounds of our canvas:
	bounds := pxCanvas.Bounds()

	// is the mouse pointer within the bounds?
	/* The reason we're checking right away to see if the mouse is within the bounds of the canvas, is the whole point of this function is to get the X, Y coordinates in
	the image which will be contained within the canvas and if the mouse is outside of that range, then there's no point in continuing.*/
	if !InBounds(ev.Position, bounds) {
		return nil, nil
	}

	// copy some information that we need:
	pxSize := float32(pxCanvas.PxSize)
	xOffset := pxCanvas.CanvasOffset.X
	yOffset := pxCanvas.CanvasOffset.Y

	/* Take the current position of the mouse(ev.Position.X) and take away(minus) one of the offsets and divide the result by the size of the pixels.*/
	x := int((ev.Position.X - xOffset) / pxSize)
	y := int((ev.Position.Y - yOffset) / pxSize)

	return &x, &y
}

func (pxCanvas *PxCanvas) LoadImage(img image.Image) {
	dimensions := img.Bounds()

	// set the columns and rows of our pxCanvas to the dimensions of the image:
	pxCanvas.PxCanvasConfig.PxCols = dimensions.Dx()
	pxCanvas.PxCanvasConfig.PxRows = dimensions.Dy()

	pxCanvas.PixelData = img
	pxCanvas.reloadImage = true
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) NewDrawing(cols, rows int) {
	pxCanvas.appState.SetFilePath("")
	pxCanvas.PxCols = cols
	pxCanvas.PxRows = rows
	pixelData := NewBlankImage(cols, rows, color.NRGBA{128, 128, 128, 255})
	pxCanvas.LoadImage(pixelData)
}
