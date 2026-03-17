package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/IgorGruvSS/guto/internal/adapters/press"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Guto settings",
	Long:  `Allows viewing and editing the settings stored in ~/.config/guto/config.yaml`,
}

var configGetCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Display a setting value or all if the key is omitted",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			settings := viper.AllSettings()
			keys := make([]string, 0, len(settings))
			for k := range settings {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				fmt.Printf("%s: %v\n", k, settings[k])
			}
			return
		}
		key := args[0]
		val := viper.Get(key)
		if val == nil {
			fmt.Printf("Setting '%s' not found.\n", key)
		} else {
			fmt.Printf("%s: %v\n", key, val)
		}
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Define a setting value",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]
		viper.Set(key, value)
		err := viper.WriteConfig()
		if err != nil {
			err = viper.SafeWriteConfig()
			if err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
					fmt.Printf("Error saving setting: %v\n", err)
					return
				}
			}
		}
		fmt.Printf("Setting updated: %s = %s\n", key, value)
	},
}

var configModelsCmd = &cobra.Command{
	Use:   "models",
	Short: "List available AI models",
	Run: func(cmd *cobra.Command, args []string) {
		// Currently fixed to Gemini, could be provider-specific in the future
		adapter := &press.GeminiAdapter{}
		models, err := adapter.ListModels()
		if err != nil {
			fmt.Printf("Error listing models: %v\n", err)
			return
		}

		fmt.Println("Available Gemini models (use guto config set press.model <name>):")
		for _, m := range models {
			// Remove "models/" prefix for cleaner output
			name := strings.TrimPrefix(m, "models/")
			fmt.Printf("- %s\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configModelsCmd)
}
