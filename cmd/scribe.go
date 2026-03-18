package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/IgorGruvSS/guto/internal/adapters/scribe"
	"github.com/spf13/cobra"
)

var scribeCmd = &cobra.Command{
	Use:   "scribe [file.wav]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Guto transcribes the audio",
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]
		fmt.Printf("✍️  Guto Scribe is transcribing: %s...\n", inputPath)

		s := &scribe.WhisperAdapter{}
		txtPath, err := s.Transcribe(inputPath)

		if err != nil {
			fmt.Printf("❌ Error in transcription: %v\n", err)
			return
		}

		scribeDir := getOutputPath("scribe")
		os.MkdirAll(scribeDir, 0755)
		newTxtPath := filepath.Join(scribeDir, filepath.Base(txtPath))
		if err := os.Rename(txtPath, newTxtPath); err == nil {
			txtPath = newTxtPath
		}

		fmt.Printf("✅ Transcription completed: %s\n", txtPath)

		fmt.Print("🗑️  Do you want to discard the original audio master? (y/n): ")
		var confirm string
		fmt.Scanln(&confirm)

		if confirm == "y" || confirm == "Y" {
			if err := os.Remove(inputPath); err != nil {
				fmt.Printf("⚠️ Error removing audio: %v\n", err)
			} else {
				fmt.Println("✅ Master discarded. Space recovered.")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(scribeCmd)
}
