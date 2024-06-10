package application

import (
	"github.com/google/uuid"
)

// Cliente representa um cliente que solicita uma entrega
type Cliente struct {
	Nome     string
	Endereco string
	Contacto string
}

// Encomenda representa uma encomenda a ser entregue
type Encomenda struct {
	EncomendaID string
	Largura     float64
	Altura      float64
	Peso        float64
}

func NovoCliente(nome, endereco, contacto string) *Cliente {
	return &Cliente{
		Nome:     nome,
		Endereco: endereco,
		Contacto: contacto,
	}
}

func NovaEncomenda(encomedaid string, largura, altura, peso float64) *Encomenda {
	return &Encomenda{
		Largura:     largura,
		Altura:      altura,
		Peso:        peso,
		EncomendaID: encomedaid,
	}
}
func (c *Cliente) SolicitarEntrega(encomenda *Encomenda) (string, error) {

	pedidoID := uuid.New().String()

	// Neste exemplo, estamos simplesmente retornando o identificador do pedido
	return pedidoID, nil
}
