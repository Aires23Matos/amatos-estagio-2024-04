package domain

import "vet-clinic/domain/entities"

type RondaRepository interface {
	Salvar(ronda *entities.Ronda) error
	BuscarID(id string) (*entities.Ronda, error)
	ListarTodos() ([]*entities.Ronda, error)
}
