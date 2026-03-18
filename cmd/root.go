package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "guto",
	Short: "Guto: Your Personal Archivist (Gutenberg's heir)",
	Long:  `📚 Guto captures your meetings, transcribes with precision, and presses the knowledge into executive summaries.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/guto/config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		configPath := home + "/.config/guto"
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			os.MkdirAll(configPath, 0755)
		}

		viper.AddConfigPath(configPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
	}
}

// getOutputPath ensures all meeting records are centralized and persistent.
// It allows defining a project-specific or cloud-synced base directory (via output.base_dir),
// preventing volatile data from being scattered across different execution paths.
func getOutputPath(subDir string) string {
	base := viper.GetString("output.base_dir")
	if base == "" {
		base = "Output"
	}

	fullPath := filepath.Join(base, subDir)
	return fullPath
}
