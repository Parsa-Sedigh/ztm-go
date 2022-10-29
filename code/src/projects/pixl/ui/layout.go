package ui

import "fyne.io/fyne/v2/container"

/* this function will setup all of the different UI elements.
We'll preodically come back to this Setup function to add new UI elements as we create them. */
func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)

	/* create a new border layout(in bottom, we're gonna place the swatches and right, the color picker and nothing on the other sides).
	The center content which will be our pixel canvas which we create later, will go after the colorPicker in function args.*/
	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker, app.PixlCanvas)

	//app.PixlWindow.SetContent(swatchesContainer)
	app.PixlWindow.SetContent(appLayout)
}
