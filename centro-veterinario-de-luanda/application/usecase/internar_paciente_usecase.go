package usecase

import (
	"vet-clinic/domain"
	"vet-clinic/domain/entities"
	"vet-clinic/repository"
	"errors"
	"fmt"
	"time"
)

type ServicePaciente struct {
	repo domain.PacienteRepository
	tutorRepo domain.TutorRepository
}

var (
    ErrEspecieInvalida = errors.New("espécie inválida. Permitido apenas Canino ou Felino")
)

func NewSevicePaciente(repo domain.PacienteRepository, tutorRepo domain.TutorRepository) *ServicePaciente {
	return &ServicePaciente{
		repo: repo,
		tutorRepo: tutorRepo,
	}
}

func (i *ServicePaciente) InternarPaciente(paciente *entities.Paciente, tutor *entities.Tutor, Diagnostico string) (string, error){
	existente, _ := i.repo.ObterPorID(paciente.ID)
	if existente != nil{
		return "",repository.ErrPacienteExiste
	}

	tutorExiste, _:= i.tutorRepo.ObterPorID(tutor.TutorID)
	if tutorExiste == nil{
		err := i.tutorRepo.Adicionar(tutor)
		if err != nil{
			return "",err
		}
	}

	if paciente.Especie != entities.Canino && paciente.Especie != entities.Felino {
        return "", ErrEspecieInvalida
    }
	paciente.Estado = entities.EstadoInternado
	paciente.Diagnostico = Diagnostico
	
	paciente.Tutor = tutor
	err := i.repo.Adicionar(paciente)
	if err != nil{
		return "",err
	}
	paciente.DatadeEntrada = time.Now()
	return fmt.Sprintf("Paciente internado com sucesso. Queixa: %s", paciente.Queixa), nil
}

func (s *ServicePaciente) ListarPacientesInternados() ([]*entities.Paciente, error) {
    return s.repo.ListarTodos()
}