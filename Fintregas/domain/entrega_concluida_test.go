package domain

import (
	"testing"
)

func TestRegistrarEntregaConcluida(t *testing.T) {
	t.Run("Registrar entrega concluída com sucesso", func(t *testing.T) {
		entrega := EntregaConcluida{}

		// Dados da entrega concluída
		encomendaID := "123"
		destinatario := "João"
		dataConclusao := "2024-05-17 10:00:00"

		// Registrar entrega concluída
		entrega.RegistrarEntregaConcluida(encomendaID, destinatario, dataConclusao)

		// Verificar se os dados foram corretamente atribuídos
		if entrega.EncomendaID != encomendaID || entrega.Destinatario != destinatario || entrega.DataConclusao != dataConclusao {
			t.Errorf("Os dados da entrega concluída não foram corretamente atribuídos: esperado %s, %s, %s, recebido %s, %s, %s", encomendaID, destinatario, dataConclusao, entrega.EncomendaID, entrega.Destinatario, entrega.DataConclusao)
		}
	})
}




