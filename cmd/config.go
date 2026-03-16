package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Gerencia as configurações do Guto",
	Long:  `Permite visualizar e editar as configurações armazenadas em ~/.config/guto/config.yaml`,
}

var configGetCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Exibe o valor de uma configuração ou todas se a chave for omitida",
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
			fmt.Printf("Configuração '%s' não encontrada.\n", key)
		} else {
			fmt.Printf("%s: %v\n", key, val)
		}
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Define o valor de uma configuração",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]
		viper.Set(key, value)
		err := viper.WriteConfig()
		if err != nil {
			// Tenta criar se não existir (SafeWriteConfig seria melhor, mas WriteConfig cria se o arquivo existir)
			// Se falhar porque não existe, usamos SafeWriteConfigAs ou WriteConfigAs
			err = viper.SafeWriteConfig()
			if err != nil {
                 // Se ja existe, ele reclama, entao eh so write mesmo
                 if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
                     fmt.Printf("Erro ao salvar configuração: %v\n", err)
                     return
                 }
			}
		}
		fmt.Printf("Configuração atualizada: %s = %s\n", key, value)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configSetCmd)
}
