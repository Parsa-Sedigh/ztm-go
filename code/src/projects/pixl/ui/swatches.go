package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"image/color"
	"pixl/swatch"
)

/* In fyne, the containers are used to contain layouts. For example we're gonna place swatches into a layout and returning that as a container. */
func BuildSwatches(app *AppInit) *fyne.Container {
	/* canvasSwatches variable acts as a buffer of CanvasObjects. We're just setting the initial value to 0 with a capacity of 64 and we can resize it later
	as needed. */
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)

	// a loop that builds our swatches:
	for i := 0; i < cap(app.Swatches); i++ {
		initialColor := color.NRGBA{255, 255, 255, 255}
		// create a new swatch:
		s := swatch.NewSwatch(app.State, initialColor, i, func(s *swatch.Swatch) {

			// deselect all the swatches and after that
			for j := 0; j < len(app.Swatches); j++ {
				app.Swatches[j].Selected = false
				canvasSwatches[j].Refresh()
			}

			// select this particular swatch
			app.State.SwatchSelected = s.SwatchIndex
			app.State.BrushColor = s.Color
		})

		/* Some initialization: We're setting the 0 index swatch to be the one that's selected. So whenever the user opens the program, the very first swatch will be
		the selected swatch. */
		if i == 0 {
			s.Selected = true
			app.State.SwatchSelected = 0

			/* If we fail to call Refresh() , then our state would still be valid, but the UI would not reflect the internal state. */
			s.Refresh()
		}

		app.Swatches = append(app.Swatches, s)
		canvasSwatches = append(canvasSwatches, s)
	}

	/* Whenever we're working with layouts such as GridWrap, they only operate on canvas objects. That's why we created that canvasSwatches variable that
	contains canvas objects. Because we're only able to utilize canvas objects within layouts such as this NewGridWrap() and we have our
	swatches as Swatch types saved in app.Swatches that way we will be able to change the color and change whether or not they're selected. Because once they're
	actual canvas objects, you lose the ability to do that.*/
	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatches...)
}
