package usecase

import (
	"time"
	"vet-clinic/domain"
	"vet-clinic/domain/entities"
)

type ServiceRonda struct {
	rondaRepo domain.RondaRepository
}

func NewServiceRonda(rondaRepo domain.RondaRepository) *ServiceRonda {
	return &ServiceRonda{
		rondaRepo: rondaRepo,
	}
}

func (r *ServiceRonda) EfetuarRonda(id string, horario time.Time, frequenciaCardiaca int, glicemia float64, batimento int, temperatura float64, pressaoArterial string) error {
	ronda := &entities.Ronda{
		ID:        id,
		Horario:   horario,
		Realizada: false,
	}
	err := ronda.NewEfetuarRonda(frequenciaCardiaca, glicemia, batimento, temperatura, pressaoArterial)
	if err != nil {
		return err
	}
	return r.rondaRepo.Salvar(ronda)
}

func (s *ServiceRonda) EfetuarDuasRondasPorDia(horario1, horario2 time.Time, frequenciaCardiaca1 int, glicemia1 float64, batimento1 int, temperatura1 float64, pressaoArterial1 string, frequenciaCardiaca2 int, glicemia2 float64, batimento2 int, temperatura2 float64, pressaoArterial2 string) error {
	err := s.EfetuarRonda("1", horario1, frequenciaCardiaca1, glicemia1, batimento1, temperatura1, pressaoArterial1)
	if err != nil {
		return err
	}
	err = s.EfetuarRonda("2", horario2, frequenciaCardiaca2, glicemia2, batimento2, temperatura2, pressaoArterial2)
	if err != nil {
		return err
	}
	return nil
}
