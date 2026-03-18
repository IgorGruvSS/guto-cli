# Contributing to Guto

First off, thank you for considering contributing to Guto! It's people like you that make Guto such a great tool.

## Code of Conduct
By participating in this project, you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md).

## How Can I Contribute?

### Reporting Bugs
* **Check the FAQ** to see if your issue is already addressed.
* **Search existing issues** to see if the bug has already been reported.
* **Use the Bug Report template** when opening a new issue.

### Suggesting Enhancements
* **Check the Roadmap** in `docs/AI_ROADMAP.md` to see what's planned.
* **Open an issue** with the "Enhancement" tag to discuss your idea before implementing it.

### Pull Requests
1. **Fork the repository** and create your branch from `main`.
2. **Implement your changes**:
   * Follow the **Hexagonal Architecture**: Define ports in `internal/ports/` and adapters in `internal/adapters/`.
   * Write idiomatic Go code.
   * Add or update tests as necessary.
3. **Ensure the code builds**: Run `go build ./...` and `make install` to test locally.
4. **Submit a Pull Request**:
   * Use a clear and descriptive title.
   * Follow the [Semantic Commits](https://www.conventionalcommits.org/) standard (e.g., `feat:`, `fix:`, `docs:`).
   * Link to any related issues.

## Development Setup

### Requirements
* **Go** (v1.23.1+)
* **FFmpeg** (with PulseAudio/ALSA support)
* **Python 3** (for Whisper)
* **Google Gemini API Key**

### Local Build
```bash
go build -o guto main.go
./guto --help
```

### Architectural Guidelines
Guto follows a strict **Hexagonal Architecture**. 
- **Ports First**: Define new functionalities in `internal/ports/interfaces.go` first.
- **Adapters**: Implement concrete logic in `internal/adapters/`.
- **CLI**: Keep the `cmd/` directory focused on orchestration and user interaction.

## Style Guide
* Follow [Effective Go](https://golang.org/doc/effective_go.html).
* Use `gofmt` or `goimports` to format your code.
* Document exported functions and types.

---
*Guto: Giving permanence back to the spoken word.*
