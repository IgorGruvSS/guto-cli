package cmd

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "github.com/spf13/cobra"
    "github.com/igorsousasilva/guto/internal/adapters/press"
)

var pressCmd = &cobra.Command{
    Use:   "press [arquivo.txt]",
    Args:  cobra.MinimumNArgs(1),
    Short: "O Guto prensa (Sumariza) a informação",
    Run: func(cmd *cobra.Command, args []string) {
        inputPath := args[0]
        fmt.Printf("📚 A Prensa do Guto está extraindo a essência de: %s...\n", inputPath)
        
        content, err := os.ReadFile(inputPath)
        if err != nil {
            fmt.Printf("❌ Erro ao ler arquivo: %v\n", err)
            return
        }
        
        p := &press.GeminiAdapter{}
        summary, err := p.Summarize(string(content))
        
        if err != nil {
            fmt.Printf("❌ Erro na sumarização: %v\n", err)
            return
        }
        
        pressDir := "Output/press"
        os.MkdirAll(pressDir, 0755)
        
        mdName := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath)) + ".md"
        outputPath := filepath.Join(pressDir, mdName)
        
        if err := os.WriteFile(outputPath, []byte(summary), 0644); err != nil {
            fmt.Printf("❌ Erro ao salvar sumário: %v\n", err)
            return
        }
        
        fmt.Printf("✅ Sumário salvo em: %s\n", outputPath)
    },
}

func init() {
    rootCmd.AddCommand(pressCmd)
}
