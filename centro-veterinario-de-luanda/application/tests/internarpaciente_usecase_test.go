package tests_test

import (
	"vet-clinic/application/usecase"
	"vet-clinic/domain/repository/mocks"
	"testing"
)

func TestInternarPaciente(t *testing.T){
	//Arrange
	mockpacienteRepo := mocks.NewMockPacienteRepository()
	service := usecase.NewPacienteUseCase(mockpacienteRepo)

	//Act
	 err := service.InternarPaciente("123", "Boby")
	//Assert
	if err != nil{
		t.Errorf("Esperado nenhum erro, mas recebeu: %v", err)
	}

	t.Run("Deve encontrar paciente", func(t *testing.T) {
		//Arrange
		mockpacienteRepo := mocks.NewMockPacienteRepository()
		service := usecase.NewPacienteUseCase(mockpacienteRepo)

		//Act
		service.InternarPaciente("123", "Boby")
		_, err := mockpacienteRepo.BuscarId("123")
		//Assert
		if err != nil {
			t.Errorf("Esperado encontrar paciente, mas recebeu: %v", err)
		}
	})

	t.Run("Deve saber se existe paciente internado", func(t *testing.T) {
		//Arrange
		mockpacienteRepo := mocks.NewMockPacienteRepository()
		service := usecase.NewPacienteUseCase(mockpacienteRepo)

		//Act
		service.InternarPaciente("123", "Boby")

		//Assert
		if len(mockpacienteRepo.Pacientes) != 1{
			t.Error("Erro não foram encontrados pacientes internados, recebeu: ", len(mockpacienteRepo.Pacientes))
		}
	})
}