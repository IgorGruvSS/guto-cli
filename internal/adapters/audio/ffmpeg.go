package audio

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type FFmpegAdapter struct {
	Cmd *exec.Cmd
}

func (a *FFmpegAdapter) Listen(filename string) error {
	outSource := viper.GetString("audio.output_monitor")
	micSource := viper.GetString("audio.input_source")

	if outSource == "" {
		detected, err := getDefaultSinkMonitor()
		if err == nil {
			outSource = detected
		} else {
			return fmt.Errorf("could not detect system audio (sink monitor). Configure manually in 'audio.output_monitor' or check PulseAudio: %v", err)
		}
	}

	if micSource == "" {
		detected, err := getDefaultSource()
		if err == nil {
			micSource = detected
		} else {
			return fmt.Errorf("could not detect microphone (source). Configure manually in 'audio.input_source': %v", err)
		}
	}

	fmt.Printf("🎙️  Capturing from:\n  - System: %s\n  - Microphone: %s\n", outSource, micSource)

	a.Cmd = exec.Command("ffmpeg", "-hide_banner", "-loglevel", "error", "-nostdin",
		"-f", "pulse", "-i", outSource,
		"-f", "pulse", "-i", micSource,
		"-filter_complex", "amix=inputs=2:duration=first",
		"-ac", "1", "-ar", "16000", filename)

	return a.Cmd.Start()
}

func (a *FFmpegAdapter) Stop() error {
	if a.Cmd != nil && a.Cmd.Process != nil {
		// Try to terminate gracefully with Interrupt before forcing a kill
		err := a.Cmd.Process.Signal(os.Interrupt)
		if err != nil {
			return fmt.Errorf("error sending interrupt signal: %v", err)
		}

		done := make(chan error, 1)
		go func() {
			done <- a.Cmd.Wait()
		}()

		select {
		case err := <-done:
			return err
		case <-time.After(2 * time.Second):
			a.Cmd.Process.Kill()
			return fmt.Errorf("ffmpeg did not exit in time and was forced to stop")
		}
	}
	return fmt.Errorf("no active recording process")
}

func getDefaultSinkMonitor() (string, error) {
	cmd := exec.Command("sh", "-c", "pactl get-default-sink")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	sinkName := strings.TrimSpace(string(out))
	return sinkName + ".monitor", nil
}

func getDefaultSource() (string, error) {
	cmd := exec.Command("sh", "-c", "pactl get-default-source")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
