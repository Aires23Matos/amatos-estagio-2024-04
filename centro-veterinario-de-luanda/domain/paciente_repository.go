package domain

import "vet-clinic/domain/entities"

type PacienteRepository interface {
	Adicionar(paciente *entities.Paciente) error
	ObterPorID(id string)(*entities.Paciente, error)
	ListarTodos() ([]*entities.Paciente, error)
}			