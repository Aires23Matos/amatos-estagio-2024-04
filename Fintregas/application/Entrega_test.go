package application

import (
    "testing"
)

func TestEntregaSolicitarEntrega(t *testing.T) {
    t.Run("Solicitar entrega com encomenda válida", func(t *testing.T) {
        entrega := Entrega{}

        // Simulando uma encomenda válida
        encomendaID := "123"
        destino := "Rua A, 123"
        destinatario := "João"

        // Solicitar entrega
        err := entrega.SolicitarEntrega(encomendaID, destino, destinatario)

        // Verificar se não ocorreu erro
        if err != nil {
            t.Errorf("Erro ao solicitar entrega: %v", err)
        }

        // Verificar se os campos da entrega foram preenchidos corretamente
        if entrega.EncomendaID != encomendaID || entrega.Destino != destino || entrega.Destinatario != destinatario {
            t.Error("Os campos da entrega não foram preenchidos corretamente")
        }
    })

    t.Run("Solicitar entrega com ID de encomenda inválido", func(t *testing.T) {
        entrega := Entrega{}

        // Simulando uma encomenda inválida
        encomendaID := ""
        destino := "Rua A, 123"
        destinatario := "João"

        // Solicitar entrega
        err := entrega.SolicitarEntrega(encomendaID, destino, destinatario)

        // Verificar se ocorreu erro
        if err == nil {
            t.Error("Esperava-se um erro ao solicitar entrega com ID de encomenda inválido")
        }
    })
}
