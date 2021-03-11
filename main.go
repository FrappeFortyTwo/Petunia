package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var (
	tap   = false
	board *canvas.Rectangle
	key   *canvas.Rectangle
	c     *fyne.Container
)

func keyTyped(e *fyne.KeyEvent) {
	switch e.Name {
	case fyne.KeyUp:
		tap = true
		break
	default:
		tap = false
	}
}

func main() {

	// create app instance
	a := app.New()
	// create app window & pass window title
	w := a.NewWindow("Petunia")
	// resize window
	w.Resize(fyne.NewSize(730, 230))
	// fix size
	w.SetFixedSize(true)

	// add content ~ key
	w.SetContent(makeKeyboard())
	// render changes
	w.Canvas().SetOnTypedKey(keyTyped)

	// separately monitor key presses
	go runGame()
	// show & run the app
	w.ShowAndRun()

}

func runGame() {
	if tap == true {
		key.FillColor = color.RGBA{74, 71, 163, 1}
		println("tap occured")
	} else {
		key.FillColor = color.White
	}

	makeKeyboard().Refresh()

}

func makeKeyboard() *fyne.Container {

	// keyboard layout
	board = canvas.NewRectangle(color.RGBA{36, 36, 36, 1})
	key = canvas.NewRectangle(color.White)

	// container to hold canvas items
	c = container.NewWithoutLayout(key, board)

	board.Resize(fyne.NewSize(700, 200))
	board.Move(fyne.NewPos(10, 10))

	key.Resize(fyne.NewSize(20, 20))
	key.Move(fyne.NewPos(10, 10))

	return c
}
