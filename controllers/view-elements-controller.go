package controllers

import (
	"log"

	"fyne.io/fyne/v2/widget"
)

func SendingBTrequestUIelement(args []string, urlInput *widget.Entry, methodInput *widget.Entry, bodyInput *widget.Entry) (string, []string) {
	log.Println(urlInput.Text, methodInput.Text, bodyInput.Text)
	args = append(args, "-X")
	args = append(args, methodInput.Text)
	args = append(args, "--data")
	args = append(args, bodyInput.Text)
	args = append(args, urlInput.Text)
	returnedData := sendRequest(args)
	args = args[:0]

	return string(returnedData), args
}
