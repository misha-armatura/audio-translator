package main

import (
	"fmt"
	"translator-app/audiodevice"
	"translator-app/translator"
)

func main() {
	err := audiodevice.Init()
	if err != nil {
		fmt.Printf("Error initializing audio: %v\n", err)
		return
	}
	defer audiodevice.Terminate()

	// create a translation service (for future use e.g Whisper or similar)
	translationService, err := translator.NewTranslationService(44100)
	if err != nil {
		fmt.Printf("Error creating translation service: %v\n", err)
		return
	}

	// call when we are done
	done := make(chan bool)

	// start working
	go audiodevice.Record(translationService, done)

	// Wait for done signal
	<-done
}
