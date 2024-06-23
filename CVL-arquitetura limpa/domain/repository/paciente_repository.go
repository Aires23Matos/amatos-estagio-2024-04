package repository

import (
	"cvl_centro_veterinario_luanda/domain/entities"
)

type PacienteRepository interface {
	Histrico(paciente *entities.Paciente) error
	Buscar(id string)(*entities.Paciente, error)
}