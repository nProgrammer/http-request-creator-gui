package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func ShowRequestUI(x int, y int) *fyne.Container {
	var args []string

	urlLabel, urlInput := createURLElement()
	methodLabel, methodInput := createMethodElement()
	bodyLabel, bodyInput := createBodyElement()
	headerLabel, headerInput := createHeaderElement()
	sendBT := createSendingElement(args, urlInput, methodInput, bodyInput, headerInput, y)

	c := container.NewWithoutLayout(
		urlLabel,
		urlInput,
		methodLabel,
		methodInput,
		bodyLabel,
		bodyInput,
		headerLabel,
		headerInput,
		sendBT,
	)
	cSize := fyne.NewSize(1000, 500)
	c.Resize(cSize)

	return c
}
