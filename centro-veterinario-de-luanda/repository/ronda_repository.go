package repository

import (
	"errors"
	"vet-clinic/domain/entities"
)

type rondarepository struct {
    rondas map[string]*entities.Ronda
}

func NewRondaRepository() *rondarepository {
    return &rondarepository{
        rondas: make(map[string]*entities.Ronda),
    }
}

func (repo *rondarepository) Salvar(ronda *entities.Ronda) error {
	repo.rondas[ronda.ID] = ronda
    return nil
}

func (repo *rondarepository) BuscarID(id string) (*entities.Ronda, error) {
    ronda, ok := repo.rondas[id]
    if !ok {
        return nil, errors.New("ronda não encontrada")
    }
    return ronda, nil
}

func (repo *rondarepository) ListarTodos() ([]*entities.Ronda, error) {
    var rondas []*entities	.Ronda
    for _, ronda := range repo.rondas {
        rondas = append(rondas, ronda)
    }
    return rondas, nil
}