package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Acme",
	Short: "Aplicação CLI para expedir mercadorias",
	Long:  `Aplicação CLI para expedir mercadorias usando Cobra em Go.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(expedirCmd)
}
