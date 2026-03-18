package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/IgorGruvSS/guto/internal/adapters/press"
	"github.com/IgorGruvSS/guto/internal/ports"
	"github.com/spf13/cobra"
)

var pressAdapter ports.Press = &press.GeminiAdapter{}

var pressCmd = &cobra.Command{
	Use:   "press [file.txt]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Guto presses (Summarizes) the information",
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]
		fmt.Fprintf(cmd.OutOrStdout(), "📚 Guto's Press is extracting the essence of: %s...\n", inputPath)

		content, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "❌ Error reading file: %v\n", err)
			return
		}

		summary, err := pressAdapter.Summarize(string(content))

		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "❌ Error in summarization: %v\n", err)
			return
		}

		pressDir := getOutputPath("press")
		os.MkdirAll(pressDir, 0755)

		mdName := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath)) + ".md"
		outputPath := filepath.Join(pressDir, mdName)

		if err := os.WriteFile(outputPath, []byte(summary), 0644); err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "❌ Error saving summary: %v\n", err)
			return
		}

		fmt.Fprintf(cmd.OutOrStdout(), "✅ Summary saved to: %s\n", outputPath)
	},
}

func init() {
	rootCmd.AddCommand(pressCmd)
}
