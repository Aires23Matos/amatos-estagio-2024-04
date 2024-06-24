package mocks

import (
	"vet-clinic/domain/entities"
	"errors"
)

type MockRondaRepository struct {
	Rondas []*entities.Ronda
}

func NewMockRondaRepository() *MockRondaRepository {
	return &MockRondaRepository{
		Rondas: make([]*entities.Ronda, 0),
	}
}

func (r *MockRondaRepository) Historico(ronda *entities.Ronda) error {
	r.Rondas = append(r.Rondas, ronda)
	return nil
}

func (r *MockRondaRepository) BuscarPacienteId(pacienteID string) ([]*entities.Ronda, error) {
	rondasDoPaciente := make([]*entities.Ronda, 0)
	for _, ronda := range r.Rondas {
		if ronda.PacienteID == pacienteID {
			rondasDoPaciente = append(rondasDoPaciente, ronda)
		}
	}
	if len(rondasDoPaciente) == 0 {
		return nil, errors.New("rondas não encontradas para o paciente")
	}
	return rondasDoPaciente, nil
}
