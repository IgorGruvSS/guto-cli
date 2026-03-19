<p align="center">
  <img src="assets/guto-main.png" width="880" alt="Guto Logo">
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
sudo make setup
```

The `Makefile` is the primary entry point and is robust enough to be executed from any directory. It will:
1. **`deps`**: Install system dependencies (`ffmpeg`, `pulseaudio-utils`, `python3`, `make`).
2. **`whisper`**: Create an isolated virtual environment for Whisper in `/opt/guto/whisper-env`.
3. **`install`**: Compile the Go binary and move it to `/usr/local/bin/guto`.

### 🔄 How to Update or Uninstall
If you have already installed Guto and want to update to the latest version or apply your local changes:

1. **Full Update (System dependencies + Whisper + Binary):**
   ```bash
   sudo make setup
   ```

2. **Quick Update (Go code changes only):**
   ```bash
   sudo make install
   ```

3. **Uninstall:**
   To remove the `guto` binary, settings, and Whisper environment:
   ```bash
   # Binary only
   sudo make uninstall

   # Interactive deep cleanup (settings, venv, cache)
   ./scripts/uninstall.sh
   ```

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

---

## 🏗️ Architecture

The project follows the principles of **Hexagonal Architecture (Ports and Adapters)**:
- `internal/ports/`: Defines audio, transcription, and AI interfaces.
- `internal/adapters/`: Concrete implementations (FFmpeg, Whisper, Gemini).
- `cmd/`: Orchestration via Cobra CLI.

All core Go code is located in the root directory following standard Go conventions, while auxiliary files are organized as follows:
- `scripts/`: System installation and uninstallation scripts.
- `docs/`: Project documentation, licenses, and community guidelines.

---

## 🤝 Contributing

Contributions are welcome! As Guto is a solo-led project, please ensure you read the [Contributing Guidelines](docs/CONTRIBUTING.md) and the [Code of Conduct](docs/CODE_OF_CONDUCT.md) before submitting a Pull Request.

### Development Workflow

1.  **Install Development Tools:**
    Ensure you have `golangci-lint` and `govulncheck` installed:
    ```bash
    make tools
    ```

2.  **Run CI Locally:**
    Before pushing your changes, run the full CI suite (Format, Lint, Vuln Check, Test):
    ```bash
    make ci
    ```

3.  **Enable Git Hooks (Recommended):**
    Install a `pre-push` hook that automatically runs `make ci` before you push:
    ```bash
    make hooks
    ```

For security-related issues, please refer to our [Security Policy](docs/SECURITY.md).

## 📄 License

This project is licensed under the **MIT License**. See the [LICENSE](docs/LICENSE) file for the full text.

---
*Guto: Giving permanence back to the spoken word.*
