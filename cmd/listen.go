package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/IgorGruvSS/guto/internal/adapters/audio"
	"github.com/IgorGruvSS/guto/internal/ports"
	"github.com/spf13/cobra"
)

var audioRecorder ports.AudioRecorder = &audio.FFmpegAdapter{}

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Guto listens to the meeting",
	Run: func(cmd *cobra.Command, args []string) {
		outputDir := getOutputPath("audio")
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "❌ Error creating directory: %v\n", err)
			return
		}

		now := time.Now().Format("2006-01-02_15-04-05")
		filename := filepath.Join(outputDir, fmt.Sprintf("meeting_%s_mixed.wav", now))

		fmt.Fprintf(cmd.OutOrStdout(), "🎙️  Guto is listening...\n📂 Destination: %s\n", filename)

		if err := audioRecorder.Listen(filename); err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "❌ Error starting recording: %v\n", err)
			return
		}

		fmt.Fprintln(cmd.OutOrStdout(), "✅ Recording. Press Enter to stop...")
		var input string
		fmt.Fscanln(cmd.InOrStdin(), &input)

		if err := audioRecorder.Stop(); err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "⚠️ Error stopping recording: %v\n", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), "🛑 Recording finished.")
	},
}

func init() {
	rootCmd.AddCommand(listenCmd)
}
