package scribe

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/viper"
)

type WhisperAdapter struct{}

func (a *WhisperAdapter) Transcribe(inputPath string) (string, error) {
	pythonBin := viper.GetString("scribe.python_bin")
	if pythonBin == "" {
		// Default to local user path if no specific binary is configured
		pythonBin = filepath.Join(os.Getenv("HOME"), ".local/share/whisper-env/bin/python3")
	}

	model := viper.GetString("scribe.model")
	if model == "" {
		model = "large-v3"
	}

	device := viper.GetString("scribe.device")
	if device == "" {
		device = "cuda"
	}

	computeType := viper.GetString("scribe.compute_type")
	if computeType == "" {
		computeType = "float16"
	}

	outputDir := filepath.Dir(inputPath)

	fmt.Printf("📝 Scribe starting transcription with: Model=%s, Device=%s\n", model, device)

	cmd := exec.Command(pythonBin, "-m", "whisper_ctranslate2.whisper_ctranslate2",
		inputPath,
		"--model", model,
		"--device", device,
		"--compute_type", computeType,
		"--language", "pt",
		"--output_dir", outputDir,
		"--output_format", "txt")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running whisper (check 'scribe.python_bin' in config): %v", err)
	}

	ext := filepath.Ext(inputPath)
	txtPath := inputPath[:len(inputPath)-len(ext)] + ".txt"
	return txtPath, nil
}
