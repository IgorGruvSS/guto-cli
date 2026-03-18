package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/IgorGruvSS/guto/internal/adapters/scribe"
	"github.com/IgorGruvSS/guto/internal/ports"
	"github.com/spf13/cobra"
)

var scribeAdapter ports.Scribe = &scribe.WhisperAdapter{}

var scribeCmd = &cobra.Command{
	Use:   "scribe [file.wav]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Guto transcribes the audio",
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]
		fmt.Fprintf(cmd.OutOrStdout(), "✍️  Guto Scribe is transcribing: %s...\n", inputPath)

		txtPath, err := scribeAdapter.Transcribe(inputPath)

		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "❌ Error in transcription: %v\n", err)
			return
		}

		scribeDir := getOutputPath("scribe")
		os.MkdirAll(scribeDir, 0755)
		newTxtPath := filepath.Join(scribeDir, filepath.Base(txtPath))
		if err := os.Rename(txtPath, newTxtPath); err == nil {
			txtPath = newTxtPath
		}

		fmt.Fprintf(cmd.OutOrStdout(), "✅ Transcription completed: %s\n", txtPath)

		fmt.Fprint(cmd.OutOrStdout(), "🗑️  Do you want to discard the original audio master? (y/n): ")
		var confirm string
		fmt.Fscanln(cmd.InOrStdin(), &confirm)

		if confirm == "y" || confirm == "Y" {
			if err := os.Remove(inputPath); err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "⚠️ Error removing audio: %v\n", err)
			} else {
				fmt.Fprintln(cmd.OutOrStdout(), "✅ Master discarded. Space recovered.")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(scribeCmd)
}
