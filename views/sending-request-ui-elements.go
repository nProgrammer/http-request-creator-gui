package views

import (
	"log"
	"request-creator-gui/controllers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func createURLElement() (*widget.Label, *widget.Entry) {
	urlLabel := widget.NewLabel("Request URL:")

	urlInput := widget.NewEntry()
	urlInputPos := fyne.NewPos(float32(150), float32(0))
	urlInput.Move(urlInputPos)
	urlInput.SetPlaceHolder("https://example.com/example-api")
	urlSize := fyne.NewSize(float32(400), float32(30))
	urlInput.Resize(urlSize)

	return urlLabel, urlInput
}

func createMethodElement() (*widget.Label, *widget.Entry) {
	methodLabel := widget.NewLabel("HTTP method:")
	methodLabelPos := fyne.NewPos(float32(0), float32(45))
	methodLabel.Move(methodLabelPos)

	methodInput := widget.NewEntry()
	methodInputPos := fyne.NewPos(float32(150), float32(45))
	methodInput.Move(methodInputPos)
	methodInput.SetPlaceHolder("Method (ex.: GET, POST etc.)")
	methodSize := fyne.NewSize(float32(400), float32(30))
	methodInput.Resize(methodSize)

	return methodLabel, methodInput
}

func createBodyElement() (*widget.Label, *widget.Entry) {
	bodyLabel := widget.NewLabel("Request body:")
	bodyLabelPos := fyne.NewPos(float32(0), float32(95))
	bodyLabel.Move(bodyLabelPos)

	bodyInput := widget.NewMultiLineEntry()
	bodyInputPos := fyne.NewPos(float32(150), float32(95))
	bodyInput.Move(bodyInputPos)
	bodyInput.SetPlaceHolder("JSON body")
	bodySize := fyne.NewSize(float32(400), float32(150))
	bodyInput.Resize(bodySize)

	return bodyLabel, bodyInput
}

func createHeaderElement() (*widget.Label, *widget.Entry) {
	headerLabel := widget.NewLabel("Request header:")
	headerLabelPos := fyne.NewPos(float32(0), float32(260))
	headerLabel.Move(headerLabelPos)

	headerInput := widget.NewMultiLineEntry()
	headerInputPos := fyne.NewPos(float32(150), float32(260))
	headerInput.Move(headerInputPos)
	headerInput.SetPlaceHolder("Request header")
	headerSize := fyne.NewSize(float32(400), float32(100))
	headerInput.Resize(headerSize)

	return headerLabel, headerInput
}

func createSendingElement(args []string, urlInput *widget.Entry, methodInput *widget.Entry, bodyInput *widget.Entry, headerInput *widget.Entry, responseJSONLabel *widget.TextGrid, y int) *widget.Button {
	sendBT := widget.NewButton("Send request", func() {
		returnedData, args := controllers.SendingBTrequestUIelement(args, urlInput, methodInput, bodyInput, headerInput)
		log.Println(returnedData)
		log.Println(args)
		responseJSONLabel.SetText(returnedData)
	})
	sendBTS := fyne.NewSize(100, 50)
	sendBT.Resize(sendBTS)
	buttonX := float32(0)
	buttonY := float32(y - 650)
	sendBTPos := fyne.NewPos(buttonX, buttonY)
	sendBT.Move(sendBTPos)

	return sendBT
}

func createResponseElement() (*widget.Label, *widget.TextGrid) {
	responseLabel := widget.NewLabel("Response: ")
	responseLabelPos := fyne.NewPos(float32(0), float32(400))
	responseLabel.Move(responseLabelPos)

	responseJSONLabel := widget.NewTextGridFromString("")
	responseJSONLabelPos := fyne.NewPos(float32(150), float32(410))
	responseJSONLabel.Move(responseJSONLabelPos)

	return responseLabel, responseJSONLabel
}
