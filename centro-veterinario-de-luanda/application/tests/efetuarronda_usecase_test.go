package tests_test

import (
	"testing"
	"vet-clinic/application/usecase"
	"vet-clinic/domain/entities"
	"vet-clinic/repository"
)

func TestEfetuarRonda(t *testing.T) {
	t.Run("Deve efetuar ronda", func(t *testing.T) {
		//Arrange
		rondaRepo := repository.NewRondaRepository()
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewRondaCaseUse(rondaRepo, pacienteRepo)

		paciente := entities.NewPaciente("123", "Bobby")
		pacienteRepo.Historico(paciente)

		exame := entities.ExameGeral{
			Glicemia:           80,
			Temperatura:        37.5,
			PressaoArterial:    120,
			FrequenciaCardiaca: 80,
			Batimento:          80,
		}

		//Act
		err := service.EfetuarRonda("123", exame)
		//Assert
		if err != nil {
			t.Errorf("Esperado nenhum erro, mas recebeu: %v", err)
		}

	})

	t.Run("deve verificar quantas rondas foram feitas", func(t *testing.T) {
		//Arrange
		rondaRepo := repository.NewRondaRepository()
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewRondaCaseUse(rondaRepo, pacienteRepo)

		paciente := entities.NewPaciente("123", "Bobby")
		pacienteRepo.Historico(paciente)

		exame := entities.ExameGeral{
			Glicemia:           80,
			Temperatura:        37.5,
			PressaoArterial:    120,
			FrequenciaCardiaca: 80,
			Batimento:          80,
		}

		//Act
		service.EfetuarRonda("123", exame)
		//Assert
		if len(rondaRepo.Rondas) != 1 {
			t.Errorf("Esperado 1 ronda, mas recebeu: %v", len(rondaRepo.Rondas))
		}
	})

	t.Run("verificar se no exame geral foram cumpridas os 5 parâmetros", func(t *testing.T) {
		//Arrange
		rondaRepo := repository.NewRondaRepository()
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewRondaCaseUse(rondaRepo, pacienteRepo)

		paciente := entities.NewPaciente("123", "Bobby")
		pacienteRepo.Historico(paciente)

		exame := entities.ExameGeral{
			Glicemia:           80,
			Temperatura:        37.5,
			PressaoArterial:    120,
			FrequenciaCardiaca: 80,
			Batimento:          80,
		}

		//Act
		service.EfetuarRonda("123", exame)
		//Assert
		ultmatRonda := rondaRepo.Rondas[len(rondaRepo.Rondas)-1]
		if ultmatRonda.Exame.Glicemia != 80 ||
			ultmatRonda.Exame.Temperatura != 37.5 ||
			ultmatRonda.Exame.PressaoArterial != 120 ||
			ultmatRonda.Exame.FrequenciaCardiaca != 80 ||
			ultmatRonda.Exame.Batimento != 80 {
			t.Error("Os parametros do exame geral não foram registrados corretamente")
		}

	})

	t.Run("deve saber dar erro quando for feito mais de duas rondas", func(t *testing.T) {
		//Arrange
		rondaRepo := repository.NewRondaRepository()
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewRondaCaseUse(rondaRepo, pacienteRepo)

		paciente := entities.NewPaciente("123", "Bobby")
		pacienteRepo.Historico(paciente)

		exame := entities.ExameGeral{
			Glicemia:           80,
			Temperatura:        37.5,
			PressaoArterial:    120,
			FrequenciaCardiaca: 80,
			Batimento:          80,
		}

		//Act
		service.EfetuarRonda("123", exame)

		//Assert
		er := service.EfetuarRonda("123", exame)
		if er != nil {
			t.Errorf("Esperado erro ao efetuar mais de duas rondas no mesmo dia")
		}
	})

}
