package entities

import (
	"errors"
	"time"
)


type Ronda struct {
    ID       string
    Horario  time.Time
    Realizada bool
	FrequenciaCardiaca int
    Glicemia         float64
    Batimento        int
    Temperatura      float64
    PressaoArterial  string
}

func (r *Ronda) EfetuarRonda(frequenciaCardiaca int, glicemia float64, batimento int, temperatura float64, pressaoArterial string) error {
    if !r.estaNoHorarioPermitido() {
        return errors.New("a ronda só pode ser efetuada entre 8 e 22 horas")
    }
    r.FrequenciaCardiaca = frequenciaCardiaca
    r.Glicemia = glicemia
    r.Batimento = batimento
    r.Temperatura = temperatura
    r.PressaoArterial = pressaoArterial
    r.Realizada = true
    return nil
}

func (r *Ronda) estaNoHorarioPermitido() bool {
    hora := r.Horario.Hour()
    return hora >= 8 && hora < 22
}