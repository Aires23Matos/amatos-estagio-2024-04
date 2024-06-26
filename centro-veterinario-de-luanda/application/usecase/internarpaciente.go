package usecase

import (
	"errors"
	"vet-clinic/domain"
	"vet-clinic/domain/entities"
)

type PacienteUseCase struct {
	pacienteRepo domain.PacienteRepository
}

func NewPacienteUseCase(pacienteRepo domain.PacienteRepository) *PacienteUseCase {
	return &PacienteUseCase{
		pacienteRepo: pacienteRepo,
	}
}

func (p *PacienteUseCase) InternarPaciente(id, nome string) error {
	_,err := p.pacienteRepo.BuscarId(id)
	if err == nil {
		return errors.New("paciente já existe")
	}

	paciente := entities.NewPaciente(id, nome)
	err = p.pacienteRepo.Historico(paciente)
	if err != nil {
		return err
	}
	return nil
}

func (p *PacienteUseCase) BuscarPaciente(id string) (*entities.Paciente, error) {
	paciente, err := p.pacienteRepo.BuscarId(id)
	if err != nil {
		return nil, err
	}
	return paciente, nil
}
