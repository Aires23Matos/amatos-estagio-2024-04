package tests_test

import (
	"testing"
	"vet-clinic/application/usecase"
	"vet-clinic/repository"
)

func TestInternarPaciente(t *testing.T) {

	t.Run("Deve internar paciente", func(t *testing.T) {
		//Arrange
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewPacienteUseCase(pacienteRepo)

		//Act
		err := service.InternarPaciente("123", "Boby")
		//Assert
		if err != nil {
			t.Errorf("Esperado nenhum erro, mas recebeu: %v", err)
		}
	})

	t.Run("Ao internar um novo pacinte deve verificar se o ID já existe", func(t *testing.T) {
		//Arrenge
		var expectedError = "paciente já existe"
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewPacienteUseCase(pacienteRepo)

		//Act
		err := service.InternarPaciente("123", "Boby")
	
		//Assert
		err = service.InternarPaciente("123", "Boby")
		if err == nil {
			t.Errorf("Esperado erro ao tentar internar paciente com o mesmo ID: %v", err)
		}

		if err.Error() != expectedError {
			t.Errorf("Esperava erro 'Paciente já existe', mas recebeu: %v", err)
		}
	})

	t.Run("Deve encontrar paciente", func(t *testing.T) {
		//Arrange
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewPacienteUseCase(pacienteRepo)

		//Act
		service.InternarPaciente("123", "Boby")
		_, err := pacienteRepo.BuscarId("123")
		//Assert
		if err != nil {
			t.Errorf("Esperado encontrar paciente, mas recebeu: %v", err)
		}
	})

	t.Run("Deve saber se existe paciente internado", func(t *testing.T) {
		//Arrange
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewPacienteUseCase(pacienteRepo)

		//Act
		service.InternarPaciente("123", "Boby")

		//Assert
		if len(pacienteRepo.Pacientes) != 1 {
			t.Error("Erro não foram encontrados pacientes internados, recebeu: ", len(pacienteRepo.Pacientes))
		}
	})
}
