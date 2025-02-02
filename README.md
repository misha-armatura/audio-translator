# audio-translator
# Real-Time Audio Pass-Through 🎤 ➡️ 🔊

A lightweight Go application that captures audio from your microphone and plays it through your speakers in real-time. Built with PortAudio for high-performance audio I/O.

## Features 🌟

- Real-time audio streaming
- Low latency audio processing
- Cross-platform support (macOS, Linux, Windows)
- Simple and clean architecture
- Ready for future audio translation capabilities

## Prerequisites 📋

- Go 1.21 or later
- PortAudio library

## Setup 🛠️

1. Install PortAudio:

   **macOS:**
   ```bash
   brew install portaudio
   ```

   **Linux:**
   ```bash
   sudo apt-get install portaudio19-dev
   ```

   **Windows:**
   ```bash
   choco install portaudio
   ```

2. Clone the repository:
   ```bash
   git clone https://github.com/misha-armatura/audio-translator.git
   cd audio-translator
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage 🚀

Run the application:
```bash
go run main 
