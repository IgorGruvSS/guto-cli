<p align="center">
  <img src="assets/guto-gemini.png" width="880" alt="Guto Logo">
</p>

<p align="center">
  <samp>
    <b>VERBA VOLANT, SCRIPTA MANENT</b><br>
    (Words fly away, the written word remains)
  </samp>
</p>

<p align="center">
  <strong>Guto</strong> is a command-line tool (CLI) developed in Go, inspired by Johannes Gutenberg's legacy. Its mission is to capture, transcribe, and "press" (summarize) the knowledge generated in meetings and conversations, transforming volatile audio into permanent and organized records.
</p>

---

## ✨ Features

- **`guto listen`**: Captures system audio (sink monitor) and microphone simultaneously using FFmpeg.
- **`guto scribe`**: Transcribes audio files into text using OpenAI Whisper (via `whisper-ctranslate2`).
- **`guto press`**: Processes transcriptions and generates executive summaries in Markdown via Google Gemini.
- **`guto scripta`**: Complete interactive flow (Record -> Name -> Transcribe -> Summarize).
- **`guto config`**: Local configuration management via CLI.

## 🚀 Quick Installation

Guto features a universal installer for major Linux distributions (**Ubuntu/Debian, Fedora, Arch, openSUSE**).

### 1. System Requirements
Ensure you have `git` and `go` installed (the script will attempt to install them if missing).

### 2. Run the Installer
```bash
git clone https://github.com/IgorGruvSS/guto-cli.git
cd guto-cli
chmod +x install.sh
sudo ./install.sh
```

The script will:
1. Install system dependencies (`ffmpeg`, `pulseaudio-utils`, `python3`).
2. Compile the Go binary and move it to `/usr/local/bin/guto`.
3. Create an isolated virtual environment for Whisper in `/opt/guto/whisper-env`.

---

## ⚙️ Post-Installation Configuration

After running the installer, you need to configure keys and paths in your user profile (without sudo):

### 1. Configure Scribe (Whisper)
Point to the Python environment created by the installer:
```bash
guto config set scribe.python_bin /opt/guto/whisper-env/bin/python3
```

### 2. Configure Press (Gemini AI)
1. Obtain a free API key at [Google AI Studio](https://aistudio.google.com/app/apikey).
2. Configure it in Guto:
```bash
guto config set press.api_key YOUR_API_KEY_HERE
```

### 3. List and Choose an AI Model
Guto allows you to choose which Gemini model you want to use (default is `gemini-2.5-flash`). To see available models:
```bash
guto config models
```
To select a model from the list:
```bash
guto config set press.model gemini-1.5-pro
```

### 4. Verify Settings
```bash
guto config get
```

---

## 📖 How to Use (The Scripta Flow)

The main command for daily use is `guto scripta`, which guides you through the entire process:

1. **Recording**: Guto starts listening to your microphone and system audio (perfect for Zoom/Teams/Meet calls).
2. **Stopping**: Press `Enter` to stop recording.
3. **Naming**: Give the meeting a name (e.g., `Daily-Sync`).
4. **Transcription**: Guto asks if you want to transcribe it now.
5. **Summary**: Guto generates a `.md` file with the executive summary, decisions, and next steps.

Files are automatically organized in the `Output/` directory:
- `Output/audio/`: Original `.wav` master files.
- `Output/scribe/`: Raw `.txt` transcriptions.
- `Output/press/`: Final `.md` summaries.

---

## 🏗️ Architecture

The project follows the principles of **Hexagonal Architecture (Ports and Adapters)**:
- `internal/ports/`: Defines audio, transcription, and AI interfaces.
- `internal/adapters/`: Concrete implementations (FFmpeg, Whisper, Gemini).
- `cmd/`: Orchestration via Cobra CLI.

This facilitates swapping providers in the future (e.g., using OpenAI GPT instead of Gemini, or local Ollama).

---

## 🛠️ Local Development

To compile manually:
```bash
go build -o guto main.go
./guto --help
```

---

## 📄 License
This project is under the MIT license. See the `LICENSE` file for details.

---
*Guto: Giving permanence back to the spoken word.*
