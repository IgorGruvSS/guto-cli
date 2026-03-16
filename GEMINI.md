# GEMINI.md - Guto CLI Project Context

## 📚 Project Overview
**Guto** (Guto: Your Personal Archivist) is a command-line tool (CLI) developed in Go, inspired by Johannes Gutenberg's legacy. Its mission is to capture, transcribe, and "press" (summarize) the knowledge generated in meetings and conversations, transforming volatile audio into permanent and organized records.

### 🛠️ Core Technologies
- **Language:** Go (v1.25.6+)
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
- **Audio Capture:** FFmpeg (via PulseAudio/ALSA)
- **Transcription:** OpenAI Whisper (via `scribe.WhisperAdapter`)
- **Summarization:** Google Gemini (via `press.GeminiAdapter`)

## 🏗️ Architecture
The project follows the principles of **Hexagonal Architecture (Ports and Adapters)**, facilitating the exchange of audio, transcription, or AI providers.

- `internal/ports/`: Defines the main interfaces (`AudioRecorder`, `Scribe`, `Press`).
- `internal/adapters/`: Concrete implementations of the ports (e.g., `ffmpeg.go`, `whisper.go`, `gemini.go`).
- `cmd/`: Defines CLI commands and orchestrates the data flow.

## 🚀 Main Commands

| Command | Description |
| :--- | :--- |
| `guto listen` | Starts audio capture from the system and microphone (auto-detection or via config). |
| `guto scribe` | Transcribes an audio file into text. |
| `guto press` | Processes text and generates an executive summary in Markdown. |
| `guto scripta` | The full interactive flow: Record -> Name -> Transcribe -> Summarize. |
| `guto config` | Manages Guto configurations (`get`, `set`). |

## 🛠️ Configuration and Installation

### Automatic Installation (Debian/Ubuntu)
The project includes an installation script to facilitate setup:

```bash
sudo ./install.sh
```
This will install dependencies like FFmpeg and Python, compile the binary, and move it to `/usr/local/bin`.

### Configuration (`guto config`)
Guto now uses a configuration file (`~/.config/guto/config.yaml`) managed via the CLI.

Examples:
```bash
# Set audio source manually (if auto-detection fails)
guto config set audio.input_source alsa_input.usb-Logitech_G733...

# Set Python path for Whisper
guto config set scribe.python_bin /path/to/venv/bin/python3

# View current settings
guto config get
```

### Dependencies
- **FFmpeg:** Mandatory for audio capture.
- **Python + Whisper:** Mandatory for transcription.
- **Gemini CLI:** Mandatory for summarization (currently).

## 🏗️ Architecture and Roadmap
The project follows the **Hexagonal Architecture**.
- **Future:** We plan to support multiple AI providers (OpenAI, Claude, Ollama). See `docs/AI_ROADMAP.md` for details.

### Audio Adapters
The `FFmpegAdapter` automatically detects the default sink (system audio) and the default source (microphone) using `pactl`. This can be overridden via configuration.

## 💻 Development

### Compilation and Installation
```bash
# To compile the binary
go build -o guto main.go

# To run directly
go run main.go [command]
```

### Code Conventions
- **Ports First:** Always define new functionalities in `internal/ports` before implementing the adapter.
- **Idiomatic CLI:** Add new commands in the `cmd/` directory following the Cobra pattern.
- **Verba Volant, Scripta Manent:** The entire design is focused on the persistence and clarity of the written record.

---
*This file serves as a context guide for AI assistants and developers interacting with the repository.*
