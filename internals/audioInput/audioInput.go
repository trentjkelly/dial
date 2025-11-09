package audioInput

import (
	"fmt"
	"encoding/binary"
	"os"
	"os/signal"
	"github.com/gordonklaus/portaudio"
	"github.com/hraban/opus"
)

func StartAudioInput() {
	fmt.Println("Audio Input started")
	// portaudio setup
	portaudio.Initialize() 
	defer portaudio.Terminate()

	// Variables for streamer and encoder inputs
	const sampleRate = 44100
	const opusSampleRate = 48000
	const channels = 1
	const maxPacketSize = 1500
	const framesPerBuffer = 960

	// Buffer that holds audio from microphone
	buffer := make([]int16, framesPerBuffer)

	// Streamer that records audio from microphone
	stream, err := portaudio.OpenDefaultStream(
		channels, 0, sampleRate, framesPerBuffer, &buffer)
	if err != nil {
		fmt.Println("STREAM: ",err)
		return;
	}
	defer stream.Close()

	//Start the streamer
	err = stream.Start()
	if err != nil {
		fmt.Println("STREAM START: ", err)
		return
	}

	// Creates a new opus encoder
	enc, err := opus.NewEncoder(opusSampleRate, channels, opus.AppVoIP)
	if err != nil {
		fmt.Println("NEWENCODER: ",err)
		return;
	}

	// Looks for the kill signal and ends for loop
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	// Creates buffer for the encoded data
	encodedData := make([]byte, maxPacketSize)

	// Creaes audio file
	audioFile, err := os.Create("audio.opus")
	if err != nil {
		fmt.Println("CREATE: ",err)
		return;
	}
	defer audioFile.Close()

	fmt.Println("Recording... Press Ctrl+C to stop")

	// Block goroutine from exiting
	for {
		select {
		case <-sigChan:
			fmt.Println("Stopping recording")
			return
		default:
			err := stream.Read()
			if err != nil {
				fmt.Println("READ: ",err)
				continue
			}
			// Encodes the buffer into the encodedData buffer and return length of buffer (n) in bytes
			n, err := enc.Encode(buffer, encodedData)
			if err != nil {
				fmt.Println("ENCODE: ",err)
				continue
			}
			
			// Writes to the audio file data from 0 - n from the encodedData buffer
			err = binary.Write(audioFile, binary.LittleEndian, encodedData[:n])
			if err != nil {
				fmt.Println("WRITE: ",err)
				return;
			}
		}
	}
}