package main

import (
	"dial/internals/audioInput"
	"dial/internals/audioOutput"
	"dial/internals/udpSender"
	"dial/internals/udpServer"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting dial...")

	myIp, myPort, friendIp, friendPort := getEnvironmentVariables()
	addresses := fmt.Sprintf("My IP: %s, My Port: %s, Friend IP: %s, Friend Port: %s", myIp, myPort, friendIp, friendPort)
	fmt.Println(addresses)

	// Make data channels

	// Start go routines
	go udpServer.StartUDPServer(myIp, myPort)
	go udpSender.StartUDPSender(friendIp, friendPort)
	go audioInput.StartAudioInput()
	go audioOutput.StartAudioOutput()

	// Block main from exiting, so that goroutines can finish
	select {}
}

func getEnvironmentVariables() (string, string, string, string) {
	fmt.Println("Loading .env file...")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	myIp := os.Getenv("MY_IP")
	myPort := os.Getenv("MY_PORT")
	friendIp := os.Getenv("FRIEND_IP")
	friendPort := os.Getenv("FRIEND_PORT")

	return myIp, myPort, friendIp, friendPort
}