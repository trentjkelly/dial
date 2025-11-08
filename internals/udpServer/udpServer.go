package udpServer

import (
	"fmt"
)

func StartUDPServer() {
	fmt.Println("UDP Server started")

	// Block goroutine from exiting
	for {}
}