package audioInput

import (
	"fmt"
)

func StartAudioInput() {
	fmt.Println("Audio Input started")

	// Block goroutine from exiting
	for {}
}