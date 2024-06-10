package domain

import (
	"fmt"
)

// EntregaConcluida representa uma entrega concluída
type EntregaConcluida struct {
	EncomendaID    string
	Destinatario   string
	DataConclusao  string
}

// RegistrarEntregaConcluida registra uma entrega concluída
func (e *EntregaConcluida) RegistrarEntregaConcluida(encomendaID, destinatario, dataConclusao string) {
	e.EncomendaID = encomendaID
	e.Destinatario = destinatario
	e.DataConclusao = dataConclusao

	fmt.Println("Entrega concluída registrada com sucesso!")
}
