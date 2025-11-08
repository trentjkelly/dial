package main

import (
	"fmt"
	"dial/udpServer"
	"dial/udpSender"
	"dial/audioInput"
	"dial/audioOutput"
)

func main() {
	fmt.Println("Starting dial...")

	// Make data channels

	// Start go routines
	go udpServer.StartUDPServer()
	go udpSender.StartUDPSender()
	go audioInput.StartAudioInput()
	go audioOutput.StartAudioOutput()

	// Block main from exiting, so that goroutines can finish
	select {}
}
