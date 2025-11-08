package udpServer

import (
	"fmt"
	"net"
)

func StartUDPServer(receiverIp string, receiverPort string) {
	fmt.Println("UDP Server started")

	address := fmt.Sprintf("%s:%s", receiverIp, receiverPort)
	fmt.Println("Address:", address)

	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Listen for incoming UDP packets
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error listening for UDP packets:", err)
		return
	}
	defer conn.Close()

	// Read incoming UDP packets
	for {
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP connection:", err)
			return
		}

		// Print the received data
		fmt.Printf("Received %d bytes from %s: %s\n", n, addr, buffer[:n])
	}
}