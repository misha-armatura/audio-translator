package audiodevice

import (
	"fmt"
	"os"
	"os/signal"

	"translator-app/translator"

	"github.com/gordonklaus/portaudio"
)

const (
	sampleRate = 44100
	channels   = 1
	frameSize  = 1024
)

var (
	inputBuffer  = make([]float32, frameSize)
	outputBuffer = make([]float32, frameSize)
)

// this method handles audio input and output to the default audio device
func Record(translationService *translator.TranslationService, done chan bool) {
	fmt.Println("Recording. Press Ctrl-C to stop.")

	hostApi, err := portaudio.DefaultHostApi()
	if err != nil {
		fmt.Printf("Error getting default host API: %v\n", err)
		return
	}

	// setup input stream
	inputParams := portaudio.HighLatencyParameters(hostApi.DefaultInputDevice, nil)
	inputParams.Input.Channels = channels
	inputParams.SampleRate = sampleRate
	inputParams.FramesPerBuffer = frameSize

	inputStream, err := portaudio.OpenStream(inputParams, &inputBuffer)
	if err != nil {
		fmt.Printf("Error opening input stream: %v\n", err)
		return
	}
	defer inputStream.Close()

	err = inputStream.Start()
	if err != nil {
		fmt.Printf("Error starting input stream: %v\n", err)
		return
	}

	// setup output stream
	outputParams := portaudio.HighLatencyParameters(nil, hostApi.DefaultOutputDevice)
	outputParams.Output.Channels = channels
	outputParams.SampleRate = sampleRate
	outputParams.FramesPerBuffer = frameSize

	outputStream, err := portaudio.OpenStream(outputParams, &outputBuffer)
	if err != nil {
		fmt.Printf("Error opening output stream: %v\n", err)
		return
	}
	defer outputStream.Close()

	err = outputStream.Start()
	if err != nil {
		fmt.Printf("Error starting output stream: %v\n", err)
		return
	}

	// handle ctrl+c
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	for {
		select {
		case <-sig:
			inputStream.Stop()
			outputStream.Stop()
			done <- true
			return
		default:
			err := inputStream.Read()
			if err != nil {
				fmt.Printf("Error reading from input stream: %v\n", err)
				continue
			}

			// process audio through translation service
			translatedBuffer, err := translationService.ProcessAudioBuffer(inputBuffer)
			if err != nil {
				fmt.Printf("Error processing audio: %v\n", err)
				continue
			}

			// copy translated audio to output buffer
			copy(outputBuffer, translatedBuffer)

			err = outputStream.Write()
			if err != nil {
				fmt.Printf("Error writing to output stream: %v\n", err)
				continue
			}
		}
	}
}

// initialize PortAudio
func Init() error {
	return portaudio.Initialize()
}

// terminate PortAudio
func Terminate() error {
	return portaudio.Terminate()
}
