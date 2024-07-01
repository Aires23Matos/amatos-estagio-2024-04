package repository

import (
	"errors"
	"vet-clinic/domain/entities"
)

type memoryrepository struct {
	pacientes map[string]*entities.Paciente
	rondas    map[string][]*entities.Ronda
}

func NewMemoryRepository() *memoryrepository {
	return &memoryrepository{
		pacientes: make(map[string]*entities.Paciente),
		rondas:    make(map[string][]*entities.Ronda),
	}
}

func (repo *memoryrepository) SalvarPaciente(paciente *entities.Paciente) error {
	repo.pacientes[paciente.ID] = paciente
	return nil
}

func (repo *memoryrepository) BuscarPacienteporID(id string) (*entities.Paciente, error) {
	paciente, ok := repo.pacientes[id]
	if !ok {
		return nil, errors.New("paciente não encontrado")
	}
	return paciente, nil
}

func (repo *memoryrepository) SalvarRonda(ronda *entities.Ronda) error {
	if _, ok := repo.rondas[ronda.ID]; !ok {
		repo.rondas[ronda.ID] = []*entities.Ronda{}
	}
	repo.rondas[ronda.ID] = append(repo.rondas[ronda.ID], ronda)
	return nil
}

func (repo *memoryrepository) BuscarRondaPcienteporID(pacienteID string) ([]*entities.Ronda, error) {
	rondas, ok := repo.rondas[pacienteID]
	if !ok {
		return nil, errors.New("rondas não encontradas para o paciente")
	}
	return rondas, nil
}

func (repo *memoryrepository) BuscarTodosPaciente() ([]*entities.Paciente, error) {
	var pacientes []*entities.Paciente
	for _, paciente := range repo.pacientes {
		pacientes = append(pacientes, paciente)
	}
	return pacientes, nil
}
