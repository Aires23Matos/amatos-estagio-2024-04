package application

import (
	"fmt"
)

// CalculoRota representa a lógica para calcular a rota mais eficiente
type CalculoRota struct{}

// CalcularRota calcula a rota mais eficiente para a entrega
func (c *CalculoRota) CalcularRota(encomendas []Encomenda) {

	for _, encomenda := range encomendas {
		fmt.Printf("Calculando rota para a encomenda ID: %s\n", encomenda.EncomendaID)
		// Lógica real para calcular a rota iria aqui
	}

	fmt.Println("Rota calculada com sucesso para todas as encomendas!")
}
