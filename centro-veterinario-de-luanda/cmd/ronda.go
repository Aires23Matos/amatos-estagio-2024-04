package cmd

import (
	"log"
	"time"
	"vet-clinic/application/usecase"
	"vet-clinic/repository"
	"github.com/spf13/cobra"
)

var (
    id                string
    frequenciaCardiaca int
    glicemia          float64
    batimento         int
    temperatura       float64
    pressaoArterial   string
)

var efetuarRondaCmd = &cobra.Command{
    Use:   "efetuar_ronda",
    Short: "Efetua uma ronda",
    Run: func(cmd *cobra.Command, args []string) {
        repo := repository.NewRondaRepository()
        service := usecase.NewServiceRonda(repo)

        horario := time.Now()
        err := service.EfetuarRonda(id, horario, frequenciaCardiaca, glicemia, batimento, temperatura, pressaoArterial)
        if err != nil {
            log.Fatalf("Erro ao efetuar ronda: %v", err)
        }

        log.Println("Ronda efetuada com sucesso")
		
        internamentos, err := repo.ListarTodos()
        if err != nil {
            log.Fatalf("Erro ao listar internamentos: %v", err)
        }

        for _, internamento := range internamentos {
            log.Printf("ID: %s\n, Horário: %s\n, Realizada: %t\n, Frequência Cardíaca: %d\n, Glicemia: %.1f\n, Batimento: %d\n, Temperatura: %.1f\n, Pressão Arterial: %s\n",
                internamento.ID, internamento.Horario, internamento.Realizada, internamento.FrequenciaCardiaca, internamento.Glicemia, internamento.Batimento, internamento.Temperatura, internamento.PressaoArterial)
        }
    },
}

func init() {
    rootCmd.AddCommand(efetuarRondaCmd)

    efetuarRondaCmd.Flags().StringVarP(&id, "id", "i", "", "ID do paciente")
    efetuarRondaCmd.Flags().IntVarP(&frequenciaCardiaca, "frequenciaCardiaca", "f", 0, "Frequência Cardíaca")
    efetuarRondaCmd.Flags().Float64VarP(&glicemia, "glicemia", "g", 0, "Glicemia")
    efetuarRondaCmd.Flags().IntVarP(&batimento, "batimento", "b", 0, "Batimento")
    efetuarRondaCmd.Flags().Float64VarP(&temperatura, "temperatura", "t", 0, "Temperatura")
    efetuarRondaCmd.Flags().StringVarP(&pressaoArterial, "pressaoArterial", "p", "", "Pressão Arterial")

    efetuarRondaCmd.MarkFlagRequired("id")
    efetuarRondaCmd.MarkFlagRequired("frequenciaCardiaca")
    efetuarRondaCmd.MarkFlagRequired("glicemia")
    efetuarRondaCmd.MarkFlagRequired("batimento")
    efetuarRondaCmd.MarkFlagRequired("temperatura")
    efetuarRondaCmd.MarkFlagRequired("pressaoArterial")
}
