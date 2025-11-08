package audioOutput

import (
	"fmt"
)

func StartAudioOutput() {
	fmt.Println("Audio Output started")

	// Block goroutine from exiting
	for {}
}
