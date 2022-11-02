package ui

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	image2 "image"
	"image/png"
	"os"
	"pixl/util"
	"strconv"
)

func BuildNewMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		sizeValidator := func(s string) error {
			width, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("must be a positive integer")
			}
			if width <= 0 {
				return errors.New("must be > 0")
			}

			/* If we return nil from this function, that indicates that our validation succeeded. */
			return nil
		}

		widthEntry := widget.NewEntry()
		widthEntry.Validator = sizeValidator

		heightEntry := widget.NewEntry()
		heightEntry.Validator = sizeValidator

		widthFormEntry := widget.NewFormItem("Width", widthEntry)
		heightFormEntry := widget.NewFormItem("Height", heightEntry)

		formItems := []*widget.FormItem{widthFormEntry, heightFormEntry}

		// the ok param will be true if user clicks on confirm button and ...
		dialog.ShowForm("New Image", "Create", "Cancel", formItems, func(ok bool) {
			if ok {
				pixelWidth := 0
				pixelHeight := 0
				if widthEntry.Validate() != nil {
					dialog.ShowError(errors.New("invalid width"), app.PixlWindow)
				} else {
					pixelWidth, _ = strconv.Atoi(widthEntry.Text)
				}

				if heightEntry.Validate() != nil {
					dialog.ShowError(errors.New("invalid height"), app.PixlWindow)
				} else {
					pixelHeight, _ = strconv.Atoi(heightEntry.Text)
				}

				app.PixlCanvas.NewDrawing(pixelWidth, pixelHeight)
			}
		}, app.PixlWindow)
	})
}

func BuildMenus(app *AppInit) *fyne.Menu {
	return fyne.NewMenu("File", BuildNewMenu(app), BuildOpenMenu(app), BuildSaveMenu(app), BuildSaveAsMenu(app))
}

func SetupMenus(app *AppInit) {
	menus := BuildMenus(app)
	mainMenu := fyne.NewMainMenu(menus)
	app.PixlWindow.SetMainMenu(mainMenu)
}

// callback function for save menu and saveAs menu
func SaveFileDialog(app *AppInit) {
	// uri is gonna be the filepath and with URIWriteCloser, we have the ability to write to it and close it
	dialog.ShowFileSave(func(uri fyne.URIWriteCloser, e error) {
		// this means that the user clicked on cancel or the provided an invalid filepath, specifically the uri, so we don't need to open or write the file
		if uri == nil {
			return
		} else {
			// encode the image data:
			err := png.Encode(uri, app.PixlCanvas.PixelData)

			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}

			/* We set the filepath here, that way after the user picks the save location and they try to save it again, they don't need to go through the popup and
			choose location over again. It'll automatically save to the correct location. */
			app.State.SetFilePath(uri.URI().Path())
		}
	}, app.PixlWindow)
}

func BuildSaveAsMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save As...", func() {
		SaveFileDialog(app)
	})
}

func BuildSaveMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		if app.State.FilePath == "" {
			SaveFileDialog(app)
		} else {
			/* We're using a function literal, that way we can display an error if sth goes wrong. */
			tryClose := func(fh *os.File) {
				err := fh.Close()
				if err != nil {
					dialog.ShowError(err, app.PixlWindow)
				}
			}

			fh, err := os.Create(app.State.FilePath)
			defer tryClose(fh)

			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}

			err = png.Encode(fh, app.PixlCanvas.PixelData)
			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}
		}
	})
}

func BuildOpenMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Open...", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, e error) {
			/* If they click on cancel or have an invalid path: */
			if uri == nil {
				return
			} else {
				// try to open the image
				image, _, err := image2.Decode(uri)
				if err != nil {
					dialog.ShowError(err, app.PixlWindow)
					return
				}

				app.PixlCanvas.LoadImage(image)
				app.State.SetFilePath(uri.URI().Path())
				imgColors := util.GetImageColors(image)

				i := 0
				for c := range imgColors {
					if i == len(app.Swatches) {
						break
					}

					app.Swatches[i].SetColor(c)
					i++
				}
			}
		}, app.PixlWindow)
	})
}
