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

	trentIp, trentPort, kareemIp, kareemPort := getEnvironmentVariables()
	addresses := fmt.Sprintf("Trent IP: %s, Trent Port: %s, Kareem IP: %s, Kareem Port: %s", trentIp, trentPort, kareemIp, kareemPort)
	fmt.Println(addresses)

	// Make data channels

	// Start go routines
	go udpServer.StartUDPServer(kareemIp, kareemPort)
	go udpSender.StartUDPSender()
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

	trentIp := os.Getenv("TRENT_IP")
	trentPort := os.Getenv("TRENT_PORT")
	kareemIp := os.Getenv("KAREEM_IP")
	kareemPort := os.Getenv("KAREEM_PORT")

	return trentIp, trentPort, kareemIp, kareemPort
}
