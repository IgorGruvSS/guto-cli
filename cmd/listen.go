package cmd

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
    "github.com/spf13/cobra"
    "github.com/igorsousasilva/guto/internal/adapters/audio"
)

var listenCmd = &cobra.Command{
    Use:   "listen",
    Short: "O Guto ouve a reunião",
    Run: func(cmd *cobra.Command, args []string) {
        outputDir := "Output/audio"
        if err := os.MkdirAll(outputDir, 0755); err != nil {
            fmt.Printf("❌ Erro ao criar diretório: %v\n", err)
            return
        }
        
        now := time.Now().Format("2006-01-02_15-04-05")
        filename := filepath.Join(outputDir, fmt.Sprintf("meeting_%s_mixed.wav", now))
        
        recorder := &audio.FFmpegAdapter{}
        
        fmt.Printf("🎙️  Guto está ouvindo...\n📂 Destino: %s\n", filename)
        
        if err := recorder.Listen(filename); err != nil {
            fmt.Printf("❌ Erro ao iniciar gravação: %v\n", err)
            return
        }
        
        fmt.Println("✅ Gravando. Pressione Enter para parar...")
        var input string
        fmt.Scanln(&input)
        
        if err := recorder.Stop(); err != nil {
            fmt.Printf("⚠️ Erro ao parar gravação: %v\n", err)
        }
        fmt.Println("🛑 Gravação finalizada.")
    },
}

func init() {
    rootCmd.AddCommand(listenCmd)
}
