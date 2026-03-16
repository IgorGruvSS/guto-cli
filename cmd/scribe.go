package cmd

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/spf13/cobra"
    "github.com/IgorGruvSS/guto/internal/adapters/scribe"
)

var scribeCmd = &cobra.Command{
    Use:   "scribe [arquivo.wav]",
    Args:  cobra.MinimumNArgs(1),
    Short: "O Guto transcreve o áudio",
    Run: func(cmd *cobra.Command, args []string) {
        inputPath := args[0]
        fmt.Printf("✍️  Guto Scribe está transcrevendo: %s...\n", inputPath)
        
        s := &scribe.WhisperAdapter{}
        txtPath, err := s.Transcribe(inputPath)
        
        if err != nil {
            fmt.Printf("❌ Erro na transcrição: %v\n", err)
            return
        }

        // Move to Output/scribe
        scribeDir := "Output/scribe"
        os.MkdirAll(scribeDir, 0755)
        newTxtPath := filepath.Join(scribeDir, filepath.Base(txtPath))
        if err := os.Rename(txtPath, newTxtPath); err == nil {
            txtPath = newTxtPath
        }
        
        fmt.Printf("✅ Transcrição concluída: %s\n", txtPath)
        
        // Funcionalidade de Limpeza de Áudio
        fmt.Print("🗑️  Deseja descartar a matriz de áudio original? (s/n): ")
        var confirm string
        fmt.Scanln(&confirm)
        
        if confirm == "s" || confirm == "S" {
            if err := os.Remove(inputPath); err != nil {
                fmt.Printf("⚠️ Erro ao remover áudio: %v\n", err)
            } else {
                fmt.Println("✅ Matriz descartada. Espaço recuperado.")
            }
        }
    },
}

func init() {
    rootCmd.AddCommand(scribeCmd)
}
