package controllers

import (
	"log"
	"os/exec"
)

func SendRequest(args []string) []byte {
	cmd := exec.Command("curl", args...)
	log.Println("curl ", args)
	data, err := cmd.Output()
	if err != nil {
		data = []byte("ERROR OCCURED WHILE SENDING REQUEST")
	}
	return data

}
