package cmd

import (
	"fmt"
	"strconv"
	"time"
	"vet-clinic/application/usecase"
	"vet-clinic/domain/entities"
	"vet-clinic/infrastructure"
	"github.com/spf13/cobra"
)

var efetuarRondaCmd = &cobra.Command{
	Use:   "efetuar-ronda",
	Short: "Efetuar uma nova Ronda",
	Long:  `Efetuar uma nova ronda para um paciente no Centro Veterinário de Luanda`,

	Run: func(cmd *cobra.Command, args []string) {
		agora := time.Now()
		if agora.Hour() < 8 || agora.Hour() >= 22 {
			fmt.Println("Rondas só podem ser efetuadas entre 8h e 22h.")
			return
		}

		if len(args) < 6 {
			fmt.Println("Uso: efetuar-ronda <PacienteID> <Glicemia> <Temperatura> <PressaoArterial> <FrequenciaCardiaca> <Batimento>")
			return
		}

		pacienteID := args[0]
		glicemia, _ := strconv.ParseFloat(args[1], 64)
		temperatura, _ := strconv.ParseFloat(args[2], 64)
		pressaoArterial, _ := strconv.ParseFloat(args[3], 64)
		frequenciaCardiaca, _ := strconv.Atoi(args[4])
		batimento, _ := strconv.Atoi(args[5])

		exame := entities.ExameGeral{
			Glicemia:           glicemia,
			Temperatura:        temperatura,
			PressaoArterial:    pressaoArterial,
			FrequenciaCardiaca: frequenciaCardiaca,
			Batimento:          batimento,
		}

		rondaRepo := infrastructure.NewRondaRepositoryImpl()
		pacienteRepo := infrastructure.NewPacienteRepositoryImpl()
		rondaUsecase := usecase.NewRondaCaseUse(rondaRepo, pacienteRepo)

		rondas, _ := rondaUsecase.BuscarRondas(pacienteID)
		contagemRondas := 0

		for _, ronda := range rondas {
			if ronda.Timestamp.Day() == agora.Day() && ronda.Timestamp.Month() == agora.Month() && ronda.Timestamp.Year() == agora.Year() {
				contagemRondas++
			}
		}
		if contagemRondas >= 2 {
			fmt.Println("Não é permitido efetuar mais de duas rondas por dia para o mesmo paciente.")
			return
		}
		err := rondaUsecase.EfetuarRonda(pacienteID, exame)

		if err != nil {
			fmt.Printf("Erro ao efetuar ronda: %v\n", err)
		} else {
			fmt.Println("Ronda efetuada com sucesso!")
		}

	},
}

func init() {
	rootCmd.AddCommand(efetuarRondaCmd)
}
