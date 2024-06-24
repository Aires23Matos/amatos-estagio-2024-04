package usecase

import (
	"vet-clinic/domain/entities"
	"vet-clinic/domain/repository"
)

type PacienteUseCase struct {
	pacienteRepo repository.PacienteRepository
}

func NewPacienteUseCase(pacienteRepo repository.PacienteRepository) *PacienteUseCase {
	return &PacienteUseCase{
		pacienteRepo: pacienteRepo,
	}
}

func (p *PacienteUseCase) InternarPaciente(id , nome string) error{
	paciente := entities.NewPaciente(id, nome)
	err := p.pacienteRepo.Historico(paciente)

	if err != nil{
		return err
	}
	return nil
}

func (p *PacienteUseCase) BuscarPaciente(id string)(*entities.Paciente, error){
	paciente, err := p.pacienteRepo.BuscarId(id)
	if err != nil{
		return nil, err
	}
	return paciente, nil
}