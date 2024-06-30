package domain

import "vet-clinic/domain/entities"

type TutorRepository interface {
	Adicionar(tutor *entities.Tutor) error
	ObterPorID(id string)(*entities.Tutor, error)
}