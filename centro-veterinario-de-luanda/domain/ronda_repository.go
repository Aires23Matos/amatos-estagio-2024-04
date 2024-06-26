package domain

import "vet-clinic/domain/entities"

type RondaRepository interface {
	Historico(ronda *entities.Ronda) error
	BuscarPacienteId(pacienteID string) ([]*entities.Ronda, error)
}