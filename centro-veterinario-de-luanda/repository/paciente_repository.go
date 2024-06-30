package repository

import (
	"vet-clinic/domain/entities"
	"errors"
)

var (
	ErrPacienteExiste = errors.New("paciente com o mesmo ID já existe")
)

type pacienterepository struct {
	data map[string]*entities.Paciente
}

func NewPacienteRepository() *pacienterepository {
	return &pacienterepository{
		data: make(map[string]*entities.Paciente),
	}
}

func (repo *pacienterepository) Adicionar(paciente *entities.Paciente) error {

	if _, existe := repo.data[paciente.ID]; existe {
		return ErrPacienteExiste
	}
	repo.data[paciente.ID] = paciente
	return nil
}

func (repo *pacienterepository) ObterPorID(id string) (*entities.Paciente, error) {
	paciente, existe := repo.data[id]
	if !existe {
		return nil, nil
	}
	return paciente, nil
}
func (repo *pacienterepository) ListarTodos() ([]*entities.Paciente, error) {
    var lista []*entities.Paciente
    for _, paciente := range repo.data {
        lista = append(lista, paciente)
    }
    return lista, nil
}
