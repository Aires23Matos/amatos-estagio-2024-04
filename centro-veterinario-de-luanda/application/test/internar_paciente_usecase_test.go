package test_test

import (
	"testing"
	"vet-clinic/application/usecase"
	"vet-clinic/domain/entities"
	"vet-clinic/repository"
)

func TestInternarPaciente(t *testing.T) {
	t.Run("Ao internar um paciente o seu estado deve ser @Internado", func(t *testing.T) {
		//Arrange
		tutorRepo := repository.NewTutorRepository()
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewSevicePaciente(pacienteRepo, tutorRepo)
		paciente := entities.Paciente{
			ID:      "1",
			Nome:    "Fido",
			Especie: entities.Canino,
		}
		tutor := entities.Tutor{
			TutorID: "101",
		}
		diagnostico := "Pneumonia"
		//Act
		service.InternarPaciente(&paciente, &tutor, diagnostico)
		//Assert
		if paciente.Estado != entities.EstadoInternado {
			t.Errorf("Estado do paciente não internado, Esperado: %v, mas obteve %v ", entities.EstadoInternado, paciente.Estado)
		}
	})

	t.Run("Não devo ter na zona de internamento pacientes com o mesmo identificador", func(t *testing.T) {
		//Arrange
		tutorRepo := repository.NewTutorRepository()
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewSevicePaciente(pacienteRepo, tutorRepo)
		paciente := entities.Paciente{
			ID:      "1",
			Nome:    "Figo",
			Queixa:  []string{"Tosse persistente"},
			Especie: entities.Canino,
		}
		paciente2 := entities.Paciente{
			ID:      "1",
			Nome:    "Gigo",
			Queixa:  []string{"Tosse persistente"},
			Especie: entities.Canino,
		}
		tutor := entities.Tutor{
			TutorID: "101",
		}
		diagnostico := "Pneumonia"
		//Act
		_, _ = service.InternarPaciente(&paciente2, &tutor, diagnostico)
		_, err := service.InternarPaciente(&paciente, &tutor, diagnostico)
		//Assert
		if err == nil {
			t.Error("esperado erro, mas obteve sucesso")
		}
		if err != repository.ErrPacienteExiste {
			t.Errorf("Já existe um indentificador com o mesmo valor, Obteve : %v", repository.ErrPacienteExiste)
		}

	})

	t.Run("Ao internar um novo paciente o sistema deve associa-lo ao seu tutor", func(t *testing.T) {
		//Arrange
		tutorRepo := repository.NewTutorRepository()
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewSevicePaciente(pacienteRepo, tutorRepo)
		paciente := entities.Paciente{
			ID:      "1",
			Nome:    "Dar",
			Especie: entities.Canino,
		}
		tutor := entities.Tutor{
			TutorID: "101",
		}
		diagnostico := "Pneumonia"
		//Act
		service.InternarPaciente(&paciente, &tutor, diagnostico)
		//Assert
		if paciente.Tutor == nil || paciente.Tutor.TutorID != tutor.TutorID {
			t.Errorf("Erro ao associar um paciente ao seu Tutor.\n esperado: %v , mas obteve %v", tutor, paciente.Tutor)
		}
	})
	t.Run("Ao internar o paciente caso o seu tutor não exista no sistema, ele é criado", func(t *testing.T) {
		//Arrange
		tutorRepo := repository.NewTutorRepository()
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewSevicePaciente(pacienteRepo, tutorRepo)
		paciente := entities.Paciente{
			ID:      "1",
			Nome:    "das",
			Queixa:  []string{"Tosse persistente"},
			Especie: entities.Canino,
		}
		tutor := entities.Tutor{
			TutorID: "101",
		}
		diagnostico := "Pneumonia"
		//Act
		mensagem, err := service.InternarPaciente(&paciente, &tutor, diagnostico)
		//Assert
		if err != nil {
			t.Errorf("Esperado erro nil, mas obteve %v", err)
		}
		tutorExistente, _ := tutorRepo.ObterPorID(tutor.TutorID)
		if tutorExistente == nil {
			t.Error("Esperado tutor criado, mas não encontrado")
		}
		if tutorExistente.TutorID != tutor.TutorID {
			t.Errorf("Erro ao criar tutor. esperado tutor %v, mas obteve %v", tutor, tutorExistente)
		}

		if mensagem != "Paciente internado com sucesso. Queixa: [Tosse persistente]" {
			t.Errorf("esperado mensagem 'Paciente internado com sucesso. Queixa: Tosse persistente', mas obteve %v", mensagem)
		}
	})

	t.Run("Ao internar paciente o tutor pode ser associado por mais de um paciente a ser internado", func(t *testing.T) {
		//Arrange
		tutorRepo := repository.NewTutorRepository()
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewSevicePaciente(pacienteRepo, tutorRepo)
		paciente := entities.Paciente{
			ID:      "1",
			Nome:    "Fig",
			Queixa:  []string{"Tosse persistente"},
			Especie: entities.Canino,
		}
		paciente2 := entities.Paciente{
			ID:      "2",
			Nome:    "Fig",
			Queixa:  []string{"Tosse persistente"},
			Especie: entities.Canino,
		}
		tutor := entities.Tutor{
			TutorID: "1",
		}
		diagnostico := "Pneumonia"
		//Act
		mensagem1, err1 := service.InternarPaciente(&paciente, &tutor, diagnostico)
		mensagem2, err2 := service.InternarPaciente(&paciente2, &tutor, diagnostico)
		//Assert
		if err1 != nil {
			t.Errorf("esperado erro nil ao internar primeiro paciente, mas obteve %v", err1)
		}
		if err2 != nil {
			t.Errorf("esperado erro nil ao internar segundo paciente, mas obteve %v", err2)
		}
		if paciente.Tutor == nil || paciente.Tutor.TutorID != tutor.TutorID {
			t.Errorf("Erro ao associar o primeiro paciente ao tutor %v.Mas Obteve %v", tutor, paciente.Tutor)
		}
		if paciente2.Tutor == nil || paciente2.Tutor.TutorID != tutor.TutorID {
			t.Errorf("Erro ao associar o segundo paciente ao tutor %v.Mas Obteve %v", tutor, paciente2.Tutor)
		}

		if len(paciente.Queixa) != 1 {
			t.Errorf("esperado 1 queixas, mas obteve %v", len(paciente.Queixa))
		}

		expectedQueixas := []string{"Tosse persistente"}
		for i, q := range paciente.Queixa {
			if q != expectedQueixas[i] {
				t.Errorf("esperado queixa %s, mas obteve %s", expectedQueixas[i], q)
			}
		}

		if mensagem1 != "Paciente internado com sucesso. Queixa: [Tosse persistente]" {
			t.Errorf("esperado mensagem 'Paciente internado com sucesso. Queixa: Tosse persistente', mas obteve %v", mensagem1)
		}

		if mensagem2 != "Paciente internado com sucesso. Queixa: [Tosse persistente]" {
			t.Errorf("esperado mensagem 'Paciente internado com sucesso. Queixa: Tosse persistente', mas obteve %v", mensagem2)
		}
	})

	t.Run("Ao internar o paciente, ele deve estar diagnosticado", func(t *testing.T) {
		//Arrange
		tutorRepo := repository.NewTutorRepository()
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewSevicePaciente(pacienteRepo, tutorRepo)
		paciente := entities.Paciente{
			ID:      "1",
			Nome:    "Fig",
			Queixa:  []string{"Tosse persistente"},
			Especie: entities.Canino,
		}
		tutor := entities.Tutor{
			TutorID: "1",
		}
		diagnostico := "Pneumonia"
		//Act
		service.InternarPaciente(&paciente, &tutor, diagnostico)
		//Assert
		if paciente.Diagnostico != diagnostico {
			t.Errorf("esperado diagnóstico %v, mas obteve %v", diagnostico, paciente.Diagnostico)
		}
	})
	t.Run("deve permitir internar apenas espécie canino ou felino", func(t *testing.T) {
		//Arrange
		tutorRepo := repository.NewTutorRepository()
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewSevicePaciente(pacienteRepo, tutorRepo)
		paciente := entities.Paciente{
			ID:      "1",
			Nome:    "Fig",
			Queixa:  []string{"Vômitos frequentes"},
			Especie: entities.Outros,
		}
		tutor := entities.Tutor{
			TutorID: "1",
		}
		diagnostico := "Pneumonia"
		//Act
		_, err := service.InternarPaciente(&paciente, &tutor, diagnostico)
		//Assert
		if err == nil {
			t.Errorf("esperado erro, mas obteve sucesso")
		}
		if err != usecase.ErrEspecieInvalida {
			t.Errorf("esperado erro %v, mas obteve %v", usecase.ErrEspecieInvalida, err)
		}

	})

	t.Run("Internar paciente com sucesso", func(t *testing.T) {
		// Arrange
		tutorRepo := repository.NewTutorRepository()
		pacienteRepo := repository.NewPacienteRepository()
		service := usecase.NewSevicePaciente(pacienteRepo, tutorRepo)

		paciente := &entities.Paciente{
			ID:      "1",
			Nome:    "Fido",
			Queixa:  []string{"Tosse persistente"},
			Especie: entities.Canino,
		}
		tutor := &entities.Tutor{
			TutorID: "1",
		}
		diagnostico := "Pneumonia"

		// Act
		_, err := service.InternarPaciente(paciente, tutor, diagnostico)

		// Assert
		if err != nil {
			t.Fatalf("erro ao internar paciente: %v", err)
		}

	})
}
