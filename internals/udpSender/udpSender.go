
package udpSender

import (
	"fmt"
)

func StartUDPSender() {
	fmt.Println("UDP Sender started")

	// Block goroutine from exiting
	select {}
}
