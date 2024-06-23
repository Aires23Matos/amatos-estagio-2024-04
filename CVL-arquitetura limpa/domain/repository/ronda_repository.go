package repository

import "cvl_centro_veterinario_luanda/domain/entities"

type RondaRepository interface {
	Historico(ronda *entities.Ronda) error
	BuscarPacienteId(pacienteID string) ([]*entities.Ronda, error)
}