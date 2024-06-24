package main

import (
	"vet-clinic/application/usecase"
	"vet-clinic/domain/entities"
	"vet-clinic/infrastructure"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	pacienteRepo := infrastructure.NewPacienteRepositoryImpl()
	rondaRepo := infrastructure.NewRondaRepositoryImpl()

	pacienteUsecase := usecase.NewPacienteUseCase(pacienteRepo)
	rondaUsecase := usecase.NewRondaCaseUse(rondaRepo, pacienteRepo)

	args := os.Args
	if len(args)<2{
		fmt.Println("Uso: ./vet-clinic <commando>")
		os.Exit(1)
	}

	comando := args[1]
	switch comando{
	case "inserir-paciente":
		if len(args) < 4{
			fmt.Println("Uso: ./vet-clinic inserir-paciente <ID> <Nome>")
			os.Exit(1)
		}
		id := args[2]
		nome := args[3]
		err := pacienteUsecase.InternarPaciente(id, nome)
		if err != nil{
			fmt.Printf("Erro ao inserir paciente: %v\n", err)
		}else{
			fmt.Println("Paciente inserido com sucesso!")
		}
	case "realizar-ronda":
		agora := time.Now()
		if agora.Hour() < 8 || agora.Hour() >= 22 {
			fmt.Println("Rondas só podem ser realizadas entre 8h e 22h.")
			os.Exit(1)
		}

		if len(args)< 9{
			fmt.Println("Uso ./vet-clinic realizar-ronda <PacienteID> <Glicemia> <Temperatura> <PressaoArterial> <FrequenciaCardiaca> <Batimento>")
			os.Exit(1)
		}
		pacienteID := args[2]
		glicemia, _ := strconv.ParseFloat(args[3], 64)
		temperatura, _ := strconv.ParseFloat(args[4], 64) 
		pressaoArterial, _ :=strconv.ParseFloat(args[5], 64) 
		frequenciaCardiaca, _ := strconv.Atoi(args[6])
		batimento, _ := strconv.Atoi(args[7])


		exame := entities.ExameGeral{
			Glicemia: glicemia,
			Temperatura: temperatura,
			PressaoArterial: pressaoArterial,
			FrequenciaCardiaca: frequenciaCardiaca,
			Batimento: batimento,
		}

		rondas, _ := rondaUsecase.BuscarRondas(pacienteID)
		contagemRondas := 0

		for _, ronda := range rondas {
			if ronda.Timestamp.Day() == agora.Day() && ronda.Timestamp.Month() == agora.Month() && ronda.Timestamp.Year() == agora.Year(){
				contagemRondas ++
			}
		}
		if contagemRondas >= 2{
			fmt.Println("Não é permitido realizar mais de duas rondas por dia para o mesmo paciente.")
			os.Exit(1)
		}

		err := rondaUsecase.EfetuarRonda(pacienteID, exame)
		
		if err != nil {
			fmt.Printf("Erro ao realizar ronda: %v\n", err)
		}else{
			fmt.Println("Ronda efetuada com sucesso!")
		}

	default:
		fmt.Println("Comando não reconhecido.")
		os.Exit(1)
	}
}