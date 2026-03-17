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
	// 1. Tenta pegar da configuração
	outSource := viper.GetString("audio.output_monitor")
	micSource := viper.GetString("audio.input_source")

	// 2. Se não estiver configurado, tenta detectar automaticamente
	if outSource == "" {
		detected, err := getDefaultSinkMonitor()
		if err == nil {
			outSource = detected
		} else {
			// Fallback ou erro
			return fmt.Errorf("não foi possível detectar o áudio do sistema (sink monitor). Configure manualmente em 'audio.output_monitor' ou verifique o PulseAudio: %v", err)
		}
	}

	if micSource == "" {
		detected, err := getDefaultSource()
		if err == nil {
			micSource = detected
		} else {
			return fmt.Errorf("não foi possível detectar o microfone (source). Configure manualmente em 'audio.input_source': %v", err)
		}
	}

	fmt.Printf("🎙️  Capturando de:\n  - Sistema: %s\n  - Microfone: %s\n", outSource, micSource)

	a.Cmd = exec.Command("ffmpeg", "-hide_banner", "-loglevel", "error", "-nostdin",
		"-f", "pulse", "-i", outSource,
		"-f", "pulse", "-i", micSource,
		"-filter_complex", "amix=inputs=2:duration=first",
		"-ac", "1", "-ar", "16000", filename)

	return a.Cmd.Start()
}

func (a *FFmpegAdapter) Stop() error {
	if a.Cmd != nil && a.Cmd.Process != nil {
		// Tenta encerrar graciosamente com Interrupt
		err := a.Cmd.Process.Signal(os.Interrupt)
		if err != nil {
			return fmt.Errorf("erro ao enviar sinal de interrupção: %v", err)
		}

		// Aguarda o processo terminar ou força o encerramento após 2 segundos
		done := make(chan error, 1)
		go func() {
			done <- a.Cmd.Wait()
		}()

		select {
		case err := <-done:
			return err
		case <-time.After(2 * time.Second):
			// Se não fechar em 2s, mata o processo
			a.Cmd.Process.Kill()
			return fmt.Errorf("ffmpeg não encerrou a tempo e foi forçado a parar")
		}
	}
	return fmt.Errorf("nenhum processo de gravação ativo")
}

// Helpers para detecção via pactl

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
