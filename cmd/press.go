package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/IgorGruvSS/guto/internal/adapters/press"
	"github.com/spf13/cobra"
)

var pressCmd = &cobra.Command{
	Use:   "press [file.txt]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Guto presses (Summarizes) the information",
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]
		fmt.Printf("📚 Guto's Press is extracting the essence of: %s...\n", inputPath)

		content, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Printf("❌ Error reading file: %v\n", err)
			return
		}

		p := &press.GeminiAdapter{}
		summary, err := p.Summarize(string(content))

		if err != nil {
			fmt.Printf("❌ Error in summarization: %v\n", err)
			return
		}

		pressDir := getOutputPath("press")
		os.MkdirAll(pressDir, 0755)

		mdName := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath)) + ".md"
		outputPath := filepath.Join(pressDir, mdName)

		if err := os.WriteFile(outputPath, []byte(summary), 0644); err != nil {
			fmt.Printf("❌ Error saving summary: %v\n", err)
			return
		}

		fmt.Printf("✅ Summary saved to: %s\n", outputPath)
	},
}

func init() {
	rootCmd.AddCommand(pressCmd)
}
