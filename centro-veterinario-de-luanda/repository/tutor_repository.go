package repository

import (
	"vet-clinic/domain/entities"
)

type tutorrepository struct {
	data map[string]*entities.Tutor
}

func NewTutorRepository()*tutorrepository{
	return &tutorrepository{
		data: make(map[string]*entities.Tutor),
	}
}

func(repo *tutorrepository) Adicionar(tutor *entities.Tutor)error{

	repo.data[tutor.TutorID] = tutor
	return nil
}

func (repo *tutorrepository) ObterPorID(id string)(*entities.Tutor, error){

	tutor, existe := repo.data[id]
	if !existe{
		return nil, nil
	}
	return tutor, nil
}
