package repository

import (
	"errors"
	"vet-clinic/domain/entities"
)

type ronda struct {
	Rondas []*entities.Ronda
}

func NewRondaRepository() *ronda {
	return &ronda{
		Rondas: make([]*entities.Ronda, 0),
	}
}

func (r *ronda) Historico(ronda *entities.Ronda) error {
	r.Rondas = append(r.Rondas, ronda)
	return nil
}

func (r *ronda) BuscarPacienteId(pacienteID string) ([]*entities.Ronda, error) {
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
