package infrastructure

import (
	"vet-clinic/domain/entities"
	"errors"
	
)

type RondaRepositoryImpl struct {
	rondas []*entities.Ronda
}

func NewRondaRepositoryImpl()*RondaRepositoryImpl{
	return &RondaRepositoryImpl{
		rondas: make([]*entities.Ronda, 0),
	}
}

func (r *RondaRepositoryImpl)Historico(ronda *entities.Ronda)error{
	// r.mu.Lock()
	// defer r.mu.Unlock()

	r.rondas = append(r.rondas, ronda)
	return nil
}

func (r *RondaRepositoryImpl) BuscarPacienteId(pacienteID string)([]*entities.Ronda, error){
	// r.mu.Lock()
	// defer r.mu.Unlock()

	// rondasDoPaciente := make([]*entities.Ronda, 0)

	// for _, ronda := range r.rondas {
	// 	if ronda.PacienteID == pacienteID{
	// 		rondasDoPaciente = append(rondasDoPaciente, ronda)
	// 	}
	// }

	// if len(rondasDoPaciente) == 0{
	// 	return nil, errors.New("rondas não encontradas para o paciente")
	// }
	// return rondasDoPaciente, nil
	var rondasDoPaciente []*entities.Ronda
    for _, ronda := range r.rondas {
        if ronda.PacienteID == pacienteID {
            rondasDoPaciente = append(rondasDoPaciente, ronda)
        }
    }
    if len(rondasDoPaciente) == 0 {
        return nil, errors.New("rondas não encontradas para o paciente")
    }
    return rondasDoPaciente, nil
}