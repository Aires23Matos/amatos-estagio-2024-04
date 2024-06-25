package cmd

import (
	"github.com/spf13/cobra"
	"os"	
)

var rootCmd = &cobra.Command{
	Use:   "vet-clinic",
	Short: "vet Clinic CLI",
	Long:  "Vet Clinic é uma ferramenta paragerenciar pacientes e rondas no Centro Veterinario de Luanda",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
