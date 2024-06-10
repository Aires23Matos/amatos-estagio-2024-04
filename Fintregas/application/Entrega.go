package application

import (
	"errors"
	"fmt"
)

// Entrega representa o processo de entrega de uma encomenda
type Entrega struct {
	EncomendaID  string
	Destino      string
	Destinatario string
}

// SolicitarEntrega solicita uma entrega para a encomenda especificada
func (e *Entrega) SolicitarEntrega(encomendaID, destino, destinatario string) error {

	if encomendaID == "" {
		return errors.New("ID da encomenda inválido")
	}

	e.EncomendaID = encomendaID
	e.Destino = destino
	e.Destinatario = destinatario

	fmt.Println("Entrega solicitada com sucesso!")
	return nil
}
