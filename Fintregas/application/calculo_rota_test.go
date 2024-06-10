package application

import (
	"testing"
)

func TestCalcularRota(t *testing.T) {
	t.Run("Calcular rota para uma lista de encomendas", func(t *testing.T) {
		calculoRota := CalculoRota{}

		// Simulando uma lista de encomendas
		encomendas := []Encomenda{
			{EncomendaID: "1", Largura: 10, Altura: 5, Peso: 2},
			{EncomendaID: "2", Largura: 8, Altura: 6, Peso: 3},
			{EncomendaID: "3", Largura: 12, Altura: 7, Peso: 4},
		}

		// Calcular rota para a lista de encomendas
		calculoRota.CalcularRota(encomendas)

	})
}
