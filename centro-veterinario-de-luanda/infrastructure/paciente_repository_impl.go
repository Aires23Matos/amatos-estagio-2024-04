package infrastructure

import (
	"vet-clinic/domain/entities"
	"errors"
)

type PacienteRepositoryImpl struct {
	pacientes map[string]*entities.Paciente
}

func NewPacienteRepositoryImpl() *PacienteRepositoryImpl {
	return &PacienteRepositoryImpl{
		pacientes: make(map[string]*entities.Paciente),
	}
}

func (p *PacienteRepositoryImpl) Historico(paciente *entities.Paciente) error {
	// p.mu.Lock()
	// defer p.mu.Unlock()

	// if _, ok := p.pacientes[paciente.ID]; ok {
	// 	return errors.New("paciente já existe")
	// }
	// p.pacientes[paciente.ID] = paciente
	// return nil
	if _, existe := p.pacientes[paciente.ID]; existe {
        return errors.New("paciente já existe")
    }
    p.pacientes[paciente.ID] = paciente
    return nil
}

func (p *PacienteRepositoryImpl) BuscarId(id string) (*entities.Paciente, error) {
	// p.mu.Lock()
	// defer p.mu.Unlock()

	// if paciente, ok := p.pacientes[id]; ok {
	// 	return paciente, nil
	// }
	// return nil, errors.New("paciente não encontrado")
	paciente, existe := p.pacientes[id]
    if !existe {
        return nil, errors.New("paciente não encontrado")
    }
    return paciente, nil
}
