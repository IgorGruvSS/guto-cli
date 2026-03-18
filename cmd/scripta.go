package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/IgorGruvSS/guto/internal/adapters/audio"
	"github.com/IgorGruvSS/guto/internal/adapters/press"
	"github.com/IgorGruvSS/guto/internal/adapters/scribe"
	"github.com/spf13/cobra"
)

var scriptaCmd = &cobra.Command{
	Use:   "scripta",
	Short: "Complete registration cycle: Verba volant, scripta manent",
	Run: func(cmd *cobra.Command, args []string) {
		audioDir := getOutputPath("audio")
		scribeDir := getOutputPath("scribe")
		pressDir := getOutputPath("press")

		os.MkdirAll(audioDir, 0755)
		os.MkdirAll(scribeDir, 0755)
		os.MkdirAll(pressDir, 0755)

		// --- PHASE 1: LISTEN ---
		now := time.Now()
		dateStr := now.Format("2006-01-02")
		tempName := fmt.Sprintf("meeting_%s_%s_mixed.wav", dateStr, now.Format("15-04-05"))
		tempPath := filepath.Join(audioDir, tempName)

		recorder := &audio.FFmpegAdapter{}
		fmt.Printf("🎙️  Guto has started capturing. The press is ready...\n")
		if err := recorder.Listen(tempPath); err != nil {
			fmt.Printf("❌ Error starting capture: %v\n", err)
			return
		}

		fmt.Println("✅ Listening... Press Enter to end the verbal registration.")
		var input string
		fmt.Scanln(&input)
		recorder.Stop()
		fmt.Println("🛑 Verbal registration finished.")

		// --- PHASE 2: RENAME (TITLING) ---
		fmt.Print("📝 Title for this Scripta (e.g., Daily-Sync) or Enter for default: ")
		var title string
		fmt.Scanln(&title)

		finalPath := tempPath
		if title != "" {
			cleanName := strings.ReplaceAll(title, " ", "-")
			newName := fmt.Sprintf("%s-%s.wav", dateStr, cleanName)
			finalPath = filepath.Join(audioDir, newName)
			os.Rename(tempPath, finalPath)
			fmt.Printf("📂 Official file: %s\n", newName)
		}

		// --- PHASE 3: SCRIBE (WRITING) ---
		fmt.Print("💡 Do you want Guto Scribe to write this registration now? (y/n): ")
		var confirmScribe string
		fmt.Scanln(&confirmScribe)

		if confirmScribe == "y" || confirmScribe == "Y" {
			s := &scribe.WhisperAdapter{}
			txtPath, err := s.Transcribe(finalPath)
			if err != nil {
				fmt.Printf("❌ Error in Scribe: %v\n", err)
				fmt.Println("⚠️  Flow interrupted due to transcription error.")
				return
			}

			newTxtPath := filepath.Join(scribeDir, filepath.Base(txtPath))
			if err := os.Rename(txtPath, newTxtPath); err == nil {
				txtPath = newTxtPath
			}
			fmt.Printf("✅ Scripta transcribed: %s\n", txtPath)

			fmt.Print("🗑️  Do you want to discard the audio master (.wav)? (y/n): ")
			var confirmClean string
			fmt.Scanln(&confirmClean)
			if confirmClean == "y" || confirmClean == "Y" {
				os.Remove(finalPath)
				fmt.Println("✅ Master discarded. Space recovered.")
			}

			// --- PHASE 4: PRESS (PRESSING) ---
			fmt.Print("🤖 Guto Press: Generate executive summary in Markdown? (y/n): ")
			var confirmPress string
			fmt.Scanln(&confirmPress)
			if confirmPress == "y" || confirmPress == "Y" {
				content, err := os.ReadFile(txtPath)
				if err != nil {
					fmt.Printf("❌ Error reading transcription: %v\n", err)
					return
				}
				p := &press.GeminiAdapter{}
				summary, err := p.Summarize(string(content))
				if err != nil {
					fmt.Printf("❌ Error in Press: %v\n", err)
					return
				}

				mdName := strings.TrimSuffix(filepath.Base(txtPath), ".txt") + ".md"
				mdPath := filepath.Join(pressDir, mdName)
				err = os.WriteFile(mdPath, []byte(summary), 0644)
				if err != nil {
					fmt.Printf("❌ Error saving summary: %v\n", err)
					return
				}
				fmt.Printf("✅ Essence pressed at: %s\n", mdPath)
			}
		} else {
			fmt.Println("ℹ️  Transcription skipped by user.")
		}
		fmt.Println("🏛️  Scripta completed successfully. Verba volant, scripta manent.")
	},
}

func init() {
	rootCmd.AddCommand(scriptaCmd)
}
