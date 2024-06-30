package cmd

import (
	"fmt"
	"vet-clinic/application/usecase"
	"vet-clinic/domain/entities"
	"vet-clinic/repository"

	"github.com/spf13/cobra"
)

var pacienteID string
var tutorID string
var especie string
var diagnostico string
var queixas []string

var internarCmd = &cobra.Command{
    Use:   "internar",
    Short: "Internar um paciente",
    Run: func(cmd *cobra.Command, args []string) {
        pacienteRepo := repository.NewPacienteRepository()
        tutorRepo := repository.NewTutorRepository()
        service := usecase.NewSevicePaciente(pacienteRepo, tutorRepo)

        paciente := &entities.Paciente{
            ID:      pacienteID,
            Nome:    args[0],
            Especie: entities.Especie(especie),
            Queixa: queixas,
        }
        tutor := &entities.Tutor{
            TutorID:   tutorID,
        }

        mensagem, err := service.InternarPaciente(paciente, tutor, diagnostico)
        pacientes, _ := service.ListarPacientesInternados()
        if err != nil {
            fmt.Printf("Erro: %v\n", err)
        } else {
            fmt.Println(mensagem)
        }
        for _, paciente := range pacientes {
            fmt.Printf("ID: %s\n, Nome: %s\n, Especie: %s\n, Diagnóstico: %s\n, Queixas: %v\n, TutorID: %s\n",
                paciente.ID, paciente.Nome, paciente.Especie, paciente.Diagnostico, paciente.Queixa, tutor.TutorID)
        }
    },
    
}

func init() {
    rootCmd.AddCommand(internarCmd)
    internarCmd.Flags().StringVarP(&pacienteID, "pacienteID", "p", "", "ID do Paciente")
    internarCmd.Flags().StringVarP(&tutorID, "tutorID", "t", "", "ID do Tutor")
    internarCmd.Flags().StringVarP(&especie, "especie", "e", "", "Espécie do Paciente")
    internarCmd.Flags().StringVarP(&diagnostico, "diagnostico", "d", "", "Diagnóstico do Paciente")
    internarCmd.Flags().StringArrayVarP(&queixas, "queixas", "q", []string{}, "Queixas do Paciente")

}
