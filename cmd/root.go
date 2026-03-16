package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "guto",
	Short: "Guto: Seu Arquivista Pessoal (O herdeiro de Gutenberg)",
	Long:  `📚 Guto captura suas reuniões, transcreve com precisão e prensa o conhecimento em sumários executivos.`,
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
		// Ensure directory exists
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			os.MkdirAll(configPath, 0755)
		}

		viper.AddConfigPath(configPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		// Config file loaded successfully
	}
}
