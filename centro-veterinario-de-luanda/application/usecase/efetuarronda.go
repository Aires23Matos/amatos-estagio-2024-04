package usecase

import (
	"errors"
	"time"
	"vet-clinic/domain"
	"vet-clinic/domain/entities"
)

type RondaCaseUse struct {
	rondaRepo    domain.RondaRepository
	pacineteRepo domain.PacienteRepository
}

func NewRondaCaseUse(rondaRepo domain.RondaRepository, pacineteRepo domain.PacienteRepository) *RondaCaseUse {
	return &RondaCaseUse{
		rondaRepo:    rondaRepo,
		pacineteRepo: pacineteRepo,
	}
}

func (r *RondaCaseUse) EfetuarRonda(pacienteID string, exame entities.ExameGeral) error {
	paciente, err := r.pacineteRepo.BuscarId(pacienteID)
	if err != nil {
		return errors.New("paciente não encontrado")
	}

	now := time.Now()
	if now.Hour() < 8 || now.Hour() >= 22 {
		return errors.New("rondas só podem ser realizada entre 8h e 22h")
	}

	rondas, _ := r.rondaRepo.BuscarPacienteId(pacienteID)
	if len(rondas) >= 2 {
		for _, ronda := range rondas {
			if ronda.Timestamp.Day() == now.Day() {
				return errors.New("já foram realizadas duas rondas hoje")
			}
		}
	}

	ronda := &entities.Ronda{
		PacienteID: paciente.ID,
		Exame:      exame,
		Timestamp:  now,
	}
	err = r.rondaRepo.Historico(ronda)
	if err != nil {
		return err
	}
	return nil
}

func (r *RondaCaseUse) BuscarRondas(pacienteID string) ([]*entities.Ronda, error) {
	rondas, err := r.rondaRepo.BuscarPacienteId(pacienteID)
	if err != nil {
		return nil, err
	}
	return rondas, nil
}
