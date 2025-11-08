package udpServer

import (
	"fmt"
	"net"
)

func StartUDPServer(receiverIp string, receiverPort string) {
	fmt.Println("UDP Server started")

	address := fmt.Sprintf("%s:%s", receiverIp, receiverPort)
	fmt.Println("Address:", address)

	_, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Block goroutine from exiting
	for {}
}