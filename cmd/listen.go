package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/IgorGruvSS/guto/internal/adapters/audio"
	"github.com/spf13/cobra"
)

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Guto listens to the meeting",
	Run: func(cmd *cobra.Command, args []string) {
		outputDir := getOutputPath("audio")
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			fmt.Printf("❌ Error creating directory: %v\n", err)
			return
		}

		now := time.Now().Format("2006-01-02_15-04-05")
		filename := filepath.Join(outputDir, fmt.Sprintf("meeting_%s_mixed.wav", now))

		recorder := &audio.FFmpegAdapter{}

		fmt.Printf("🎙️  Guto is listening...\n📂 Destination: %s\n", filename)

		if err := recorder.Listen(filename); err != nil {
			fmt.Printf("❌ Error starting recording: %v\n", err)
			return
		}

		fmt.Println("✅ Recording. Press Enter to stop...")
		var input string
		fmt.Scanln(&input)

		if err := recorder.Stop(); err != nil {
			fmt.Printf("⚠️ Error stopping recording: %v\n", err)
		}
		fmt.Println("🛑 Recording finished.")
	},
}

func init() {
	rootCmd.AddCommand(listenCmd)
}
