package mocks

import (
	"vet-clinic/domain/entities"
	"errors"
)

type MockPacienteRepository struct {
	Pacientes map[string]*entities.Paciente
}

func NewMockPacienteRepository() *MockPacienteRepository {
	return &MockPacienteRepository{
		Pacientes: make(map[string]*entities.Paciente),
	}
}

func (m *MockPacienteRepository) Historico(paciente *entities.Paciente) error {
	if _, ok := m.Pacientes[paciente.ID]; ok {
		return errors.New("paciente já existe")
	}
	m.Pacientes[paciente.ID] = paciente
	return nil
}

func (m *MockPacienteRepository) BuscarId(id string) (*entities.Paciente, error) {
	if paciente, ok := m.Pacientes[id]; ok {
		return paciente, nil
	}
	return nil, errors.New("paciente não encontrado")
}
