package application

import (
    "testing"
)

func TestClienteSolicitarEntrega(t *testing.T) {
	t.Run("Cliente solicita uma entrega", func(t *testing.T) {
		// Dado
		cliente := NovoCliente("João", "Rua A, 123", "joao@example.com")
		encomenda := NovaEncomenda("10", 20, 23, 12)

		// Quando
		pedidoID, err := cliente.SolicitarEntrega(encomenda)

		// Então
		if err != nil {
			t.Errorf("Erro ao solicitar entrega: %v", err)
		}

		if pedidoID == "" {
			t.Error("Identificador de pedido vazio")
		}
	})
   
}
