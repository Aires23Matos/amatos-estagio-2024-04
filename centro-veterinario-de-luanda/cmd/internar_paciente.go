package cmd

import (
	"fmt"
	"vet-clinic/application/usecase"
	"vet-clinic/infrastructure"

	"github.com/spf13/cobra"
)

var internarPacienteCmd = &cobra.Command{
	Use:   "internar-paciente",
	Short: "Internar um novo paciente",
	Long:  `Internar um novo paciente no sistema do centro Veterinario de Luanda`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Uso: internar-paciente <ID> <Nome>")
			return
		}

		id := args[0]
		nome := args[1]

		pacienteRepo := infrastructure.NewPacienteRepositoryImpl()
		pacienteUsecase := usecase.NewPacienteUseCase(pacienteRepo)

		err := pacienteUsecase.InternarPaciente(id, nome)
		
		if err != nil {
			fmt.Printf("Erro ao internado paciente: %v\n", err)
		} else {
			fmt.Println("Paciente internado com sucesso!")
		}
	},
}

func init() {
	rootCmd.AddCommand(internarPacienteCmd)
}
