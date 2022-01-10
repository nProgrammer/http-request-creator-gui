package main

import (
	"request-creator-gui/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	x := 1000
	y := 500
	windowSize := fyne.NewSize(float32(x), float32(y))
	w.Resize(windowSize)

	content := views.ShowRequestUI(x, y)

	w.SetContent(content)

	w.ShowAndRun()
}
