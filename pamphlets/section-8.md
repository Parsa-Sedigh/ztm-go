# 08 - Final Milestone Project Pixl

## 88-001 gcc Installation
brew cleanup is optional which is a cleanup to free up your space.

On linux or WSL, by installing build-essential, it will install gcc inner system.

### How to install gcc on windows without using WSL:
First we need to install msys2 software distribution platform. By going to msys2.org . After installing, you will be greeted with msys terminal. Run:
`pacman -Syu` and this is gonna update the system. Then in the start menu type `msys` and click on MSYS2 MSYS application. Then run `pacman -Su`,
after that, `pacman -S --neded base-devel mingw-w64-x86_64-toolchain`. It's gonna ask your for a selection, just hit enter and again enter and enter to proceed
when it asks a [Y/n] question.

Then in start menu type msys and click on `MSYS2 MinGW x64`. Then we need to add go to the terminal. We can do so with:
`export PATH=$PATH:/c/Program\ Files/Go/bin:~/go/bin`.

We confirm that everything is installed properly by running:
`go`
and:
`gcc` which should say: gcc.exe: fatal error: no input files.

Next, change directory to the location where you downloaded the source files of the project to run the project.

## 89-002 Intro & Project Setup
Run command below to create a go module:
```shell
 go mod init pixl
```

Create a folder called apptype then pixl directory which is gonna have our main application that runs the project. Then pxcanvas directory which is gonna containing
our source files for our actual pixel canvas area where we can draw things.

The first part of the project we're gonna work is the component that allows us to draw pixels into an image.

## 90-003 Canvas Overview & State
In apptype directory create apptype.go

The GUI toolkit that we'll be using for this project is named `fyne`.

In order to communicate with the gui toolkit and specify the size of our drawing area, we have to use the specific size type that's provided by the library.
So we used `fyne.Size` and `fyne.Position` in `PxCanvasConfig struct`.

The `PxRows` and `PxCols` represent the number of pixels in our pixel art.

The PxSize is just how big on screen I want the pixels. So if we set it to 1, it'll be an actual screen pixel. But if we set PxSize to for example 10,
that way each pixel in the art is 10 pixels on screen.

By running `go mod tidy`, we can download the packages.

If you're still getting errors regarding installing 3rd party packages after the command above, you can hit: `ctrl+shift+p` and then type `reload` to reload the window
in vscode which is gonna take care of language server.

## 91-004 Creating a Swatch
The first widget we're gonna create is a swatch which is just a square and we're gonna click it and it'll change the color.

In the fyne toolkit, in order to create sth that you can click on, you have to implement the widget interface.

The Widget interface has two components: 
- `CanvasObject` embedded type which keeps track of the size and position and also provides functionality to retreive and set those.
- `CreateRenderer` function which returns a `WidgetRenderer`. This function will be called by the toolkit itself and WidgetRenderer that we have to return,
will define how the widget actually looks.

In order to create a `WidgetRenderer`, we have to also implment the `WidgetRenderer` interface:
`BackgroundColor` in `WidgetRenderer` interface is deprecated. We do have to implement Destroy but we can just leave it blank. So the toolkit will take care of it for us.

The mouse.go will contain behaviors such as mouse clicking and mouseover.



## 92-005 Swatch Logout & First Run
In ui directory, create types.go . Then create swatches.go and create `BuildSwatches` func there.

Create layout.go in ui folder.

Run the project by:
```shell
go run -v ./pixl #-v for verbose to list out everything that's being compiled.
```
When you first run the compilation, it may take a little while to run, because there are a lot of deps to compile and using `-v` option will list them out. So you know
you're making progress.

Now you might get some errors and that's because we added a bunch of deps from fyne and some of those have additional deps. So to fetch those, run:
```shell
go mod tidy
```
Then try to run the program again.

yeah, it's weird, because our app is only 20px wide.

You can close the app if the app is highlighted by pressing: `alt + f4` or ctrl+c on the terminal.

## 93-006 Color Picker & App Layout
For creating color picker, create `picker.go` in `ui` directory.

VBox places elements on top of each other.

## 94-007 Pixel Canvas Structure
Now we're ready to start working on our pixel drawing canvas. Create pxcanvas.go in pxcanvas directory.

RGBA vs NRGBA: NRGBA allows us to set the alpha value independently of the other colors, whereas the RGBA will have the alpha and the RGB components multiplied
together. Using the NRGBA allows us to easily change the alpha without having to worry about the colors getting changed as well.

## 95-008 Pixel Canvas Renderer
Now that we have basics of our `pxCanvas`, we need a renderer to display some stuff on the screen. So create pxcanvasrenderer.go in pxcanvas folder.
## 96-009 Pixel Canvas Layout
In types.go , add `PixlCanvas` to AppInit struct. Then in pixl.go , create a pixel canvas by first creating a configuration. Then change layout.go to 
change the layout to load our pixel canvas in layout.

### Layout slide:
The border layout has right side(in this case, Picker), bottom side(swatches), left side and top side(nothing in this case). So we're ignoring top and left and we want to
put our pixel canvas in the center. To place sth in the center of border layout, it should be the last arg to `NewBorder()` function.
The args in `NewBorder` would be: top, bottom, left, right and then everything else afterwards will be included in the center.

## 97-010 Panning & Zooming
`ops.go` in `pxcanvas` directory, represents the mouse operations.

`PointEvent` is just the mouse location, the `MouseEvent` is actual buttons being pressed. 

In pxcanvas directory, create mouse.go where we put our mouse events and there, we implement the interfaces needed to handle mouse movements.

Now we can move the canvas by holding down the scroll wheel on mouse and move the mouse around.

Also you can scroll down and up to zoom in or zoom out the canvas. 

## 98-011 Painting Pixels
Now it's time to start paining. So we need to implement a brush. In apptype.go , create an interface called Brushable. Now implement that interface on PxCanvas.
So go to pxcanvas>pxcanvas.go .

With MouseToCanvasXY function, we can map our mouse position to the pixel coordinates. Now we can create a brush by creating a new folder called `brush` in `pxcanvas` folder and
create `brush.go` there.

Currently, if we click, we can paint a pixel, but if we click and hold the left click, we only paint one at a time. We'll have to implement that functionality next.
For this, in mouse.go , implement `MouseMove` function.

Now we can click and drag it to paint!

## 99-012 Cursor Display
In brush.go , create a function named `Cursor`.

Now that we have a functional canvas, we need to be able to work with files, including saving and loading.

## 100-013 Creating New Images
Now that we can create files, we need to be able to save the files which we're gonna do in the next lecture.
## 101-014 Saving Images
## 102-015 Loading Images
Create the functionality of opening existing image files.