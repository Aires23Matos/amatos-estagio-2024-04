package entities

import "time"

type EstadoPaciente string
type Especie string

const (
	EstadoInternado EstadoPaciente = "Internado"
	Canino          Especie        = "Canino"
	Felino          Especie        = "Felino"
	Outros          Especie        = "Outros"
)

type Paciente struct {
	ID            string
	Nome          string
	Queixa        []string
	Estado        EstadoPaciente
	Diagnostico   string
	Tutor         *Tutor
	Especie       Especie
	DatadeEntrada time.Time
}

type Tutor struct {
	TutorID string
}
