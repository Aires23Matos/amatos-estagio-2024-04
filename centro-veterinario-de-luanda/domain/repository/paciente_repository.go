package repository

import (
	"vet-clinic/domain/entities"
)

type PacienteRepository interface {
	Historico(paciente *entities.Paciente) error
	BuscarId(id string)(*entities.Paciente, error)
}