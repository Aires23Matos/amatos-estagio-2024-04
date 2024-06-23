package tests_test

import (
	"cvl_centro_veterinario_luanda/application/usecase"
	"cvl_centro_veterinario_luanda/domain/entities"
	"cvl_centro_veterinario_luanda/domain/repository/mocks"
	"testing"
)

func TestEfetuarRonda(t *testing.T) {
	//Arrange
	mockRondaRepo := mocks.NewMockRondaRepository()
	mockpacienteRepo := mocks.NewMockPacienteRepository()
	service := usecase.NewRondaCaseUse(mockRondaRepo, mockpacienteRepo)

	paciente := entities.NewPaciente("123", "Bobby")
	mockpacienteRepo.Histrico(paciente)

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

	if len(mockRondaRepo.Rondas) != 1 {
		t.Errorf("Esperado 1 ronda, mas recebeu: %v", len(mockRondaRepo.Rondas))
	}

	ultmatRonda := mockRondaRepo.Rondas[len(mockRondaRepo.Rondas)-1]
	if ultmatRonda.Exame.Glicemia != 80 ||
		ultmatRonda.Exame.Temperatura != 37.5 ||
		ultmatRonda.Exame.PressaoArterial != 120 ||
		ultmatRonda.Exame.FrequenciaCardiaca != 80 ||
		ultmatRonda.Exame.Batimento != 80 {
		t.Error("Os parametros do exame geral não foram registrados corretamente")
	}

	er := service.EfetuarRonda("123", exame)
	if er != nil {
		t.Errorf("Esperado erro ao realizar mais de duas rondas no mesmo dia")
	}
}
