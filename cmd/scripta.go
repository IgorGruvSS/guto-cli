package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/igorsousasilva/guto/internal/adapters/audio"
	"github.com/igorsousasilva/guto/internal/adapters/press"
	"github.com/igorsousasilva/guto/internal/adapters/scribe"
	"github.com/spf13/cobra"
)

var scriptaCmd = &cobra.Command{
	Use:   "scripta",
	Short: "Ciclo completo de registro: Verba volant, scripta manent",
	Run: func(cmd *cobra.Command, args []string) {
		audioDir := "Output/audio"
		scribeDir := "Output/scribe"
		pressDir := "Output/press"
		
		os.MkdirAll(audioDir, 0755)
		os.MkdirAll(scribeDir, 0755)
		os.MkdirAll(pressDir, 0755)

		// --- FASE 1: LISTEN (OUVIR) ---
		now := time.Now()
		dateStr := now.Format("2006-01-02")
		tempName := fmt.Sprintf("meeting_%s_%s_mixed.wav", dateStr, now.Format("15-04-05"))
		tempPath := filepath.Join(audioDir, tempName)

		recorder := &audio.FFmpegAdapter{}
		fmt.Printf("🎙️  Guto iniciou a captura. O prelo está pronto...\n")
		if err := recorder.Listen(tempPath); err != nil {
			fmt.Printf("❌ Erro ao iniciar captura: %v\n", err)
			return
		}

		fmt.Println("✅ Ouvindo... Pressione Enter para encerrar o registro verbal.")
		var input string
		fmt.Scanln(&input)
		recorder.Stop()
		fmt.Println("🛑 Verba finalizada.")

		// --- FASE 2: RENAME (TITULAR) ---
		fmt.Print("📝 Título para este Scripta (ex: Daily-Sync) ou Enter para padrão: ")
		var title string
		fmt.Scanln(&title)

		finalPath := tempPath
		if title != "" {
			cleanName := strings.ReplaceAll(title, " ", "-")
			newName := fmt.Sprintf("%s-%s.wav", dateStr, cleanName)
			finalPath = filepath.Join(audioDir, newName)
			os.Rename(tempPath, finalPath)
			fmt.Printf("📂 Arquivo oficial: %s\n", newName)
		}

		// --- FASE 3: SCRIBE (ESCREVER) ---
		fmt.Print("💡 Deseja que o Guto Scribe escreva este registro agora? (s/n): ")
		var confirmScribe string
		fmt.Scanln(&confirmScribe)

		if confirmScribe == "s" || confirmScribe == "S" {
			s := &scribe.WhisperAdapter{}
			txtPath, err := s.Transcribe(finalPath) // This creates .txt in the same dir as audio
			if err != nil {
				fmt.Printf("❌ Erro no Scribe: %v\n", err)
			} else {
				// Move .txt to Output/scribe
				newTxtPath := filepath.Join(scribeDir, filepath.Base(txtPath))
				if err := os.Rename(txtPath, newTxtPath); err == nil {
					txtPath = newTxtPath
				}
				fmt.Printf("✅ Scripta transcrito: %s\n", txtPath)

				// CLEANUP
				fmt.Print("🗑️  Deseja descartar a matriz de áudio (.wav)? (s/n): ")
				var confirmClean string
				fmt.Scanln(&confirmClean)
				if confirmClean == "s" || confirmClean == "S" {
					os.Remove(finalPath)
					fmt.Println("✅ Matriz descartada. Espaço recuperado.")
				}

				// --- FASE 4: PRESS (PRENSAR) ---
				fmt.Print("🤖 Guto Press: Gerar sumário executivo em Markdown? (s/n): ")
				var confirmPress string
				fmt.Scanln(&confirmPress)
				if confirmPress == "s" || confirmPress == "S" {
					content, _ := os.ReadFile(txtPath)
					p := &press.GeminiAdapter{}
					summary, err := p.Summarize(string(content))
					if err != nil {
						fmt.Printf("❌ Erro no Press: %v\n", err)
					} else {
						mdName := strings.TrimSuffix(filepath.Base(txtPath), ".txt") + ".md"
						mdPath := filepath.Join(pressDir, mdName)
						os.WriteFile(mdPath, []byte(summary), 0644)
						fmt.Printf("✅ Essência prensada em: %s\n", mdPath)
					}
				}
			}
		}
		fmt.Println("🏛️  Scripta concluído. Verba volant, scripta manent.")
	},
}

func init() {
	rootCmd.AddCommand(scriptaCmd)
}
