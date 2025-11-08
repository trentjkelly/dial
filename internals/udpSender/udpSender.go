
package udpSender

import (
	"fmt"
	"net"
)

func StartUDPSender(ip string, port string) {
	fmt.Println("UDP Sender started")
	address := ip + ":" + port
	udpAddr, err := net.ResolveUDPAddr("udp", address) // Change the IP Address
	if err != nil {
		fmt.Println(err)
		return;
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println(err)
		return;
	}
	defer conn.Close()

	message :=[]byte("TEST UDP CONN MESSAGE")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Failed to write message to connection")
		return;
	}


	// Block goroutine from exiting
	for {}

}
