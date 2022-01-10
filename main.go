package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	x := float32(1000)
	y := float32(500)
	windowSize := fyne.NewSize(x, y)
	w.Resize(windowSize)

	urlLabel := widget.NewLabel("Request URL:")

	urlInput := widget.NewEntry()
	urlInputPos := fyne.NewPos(float32(150), float32(0))
	urlInput.Move(urlInputPos)
	urlInput.SetPlaceHolder("https://example.com/example-api")
	urlSize := fyne.NewSize(float32(400), float32(30))
	urlInput.Resize(urlSize)

	methodLabel := widget.NewLabel("HTTP method:")
	methodLabelPos := fyne.NewPos(float32(0), float32(45))
	methodLabel.Move(methodLabelPos)

	methodInput := widget.NewEntry()
	methodInputPos := fyne.NewPos(float32(150), float32(45))
	methodInput.Move(methodInputPos)
	methodInput.SetPlaceHolder("Method (ex.: GET, POST etc.)")
	methodSize := fyne.NewSize(float32(400), float32(30))
	methodInput.Resize(methodSize)

	sendBT := widget.NewButton("Send request", func() {
		log.Println(urlInput.Text, methodInput.Text)
	})
	sendBTS := fyne.NewSize(100, 50)
	sendBT.Resize(sendBTS)
	buttonX := float32(0)
	buttonY := y / 5
	sendBTPos := fyne.NewPos(buttonX, buttonY)
	sendBT.Move(sendBTPos)

	c := container.NewWithoutLayout(
		urlLabel,
		urlInput,
		methodLabel,
		methodInput,
		sendBT,
	)
	cSize := fyne.NewSize(1000, 500)
	c.Resize(cSize)
	w.SetContent(c)

	w.ShowAndRun()
}
