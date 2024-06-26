package repository

import (
	"errors"
	"vet-clinic/domain/entities"
)

type paciente struct {
	Pacientes map[string]*entities.Paciente
}

func NewPacienteRepository() *paciente {
	return &paciente{
		Pacientes: make(map[string]*entities.Paciente),
	}
}

func (p *paciente) Historico(paciente *entities.Paciente) error {
	if _, ok := p.Pacientes[paciente.ID]; ok {
		return errors.New("paciente já existe")
	}
	p.Pacientes[paciente.ID] = paciente
	return nil
}

func (p *paciente) BuscarId(id string) (*entities.Paciente, error) {
	paciente, existe := p.Pacientes[id]
    if !existe {
        return nil, errors.New("paciente não encontrado")
    }
	return paciente,nil
}
