package entities

import "time"

type Paciente struct {
	ID          string
	Nome        string
	
	DataEntrada time.Time
}

func NewPaciente(id, nome string) *Paciente {
	return &Paciente{
		ID:          id,
		Nome:        nome,
		DataEntrada: time.Now(),
	}
}
