package entities

import "time"

type ExameGeral struct {
	Glicemia           float64
	Temperatura        float64
	PressaoArterial    float64
	FrequenciaCardiaca int
	Batimento          int
}

type Ronda struct {
	PacienteID string
	Exame      ExameGeral
	Timestamp  time.Time
}