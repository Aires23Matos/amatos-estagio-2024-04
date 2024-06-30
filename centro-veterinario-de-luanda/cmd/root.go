package cmd

import (
    "os"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "vet-clinic",
    Short: "vet-clinic",
    Long:  `CLI para gerenciar o internamento de pacientes no Centro Veterinário de Luanda.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize()
}
