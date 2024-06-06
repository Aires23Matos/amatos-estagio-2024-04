package application_test

import (
	"Acme/internal/application"
	"testing"
)

func TestExpedirMercadoria(t *testing.T) {
	t.Run("Deve Criar uma Guia de Remessa", func(t *testing.T) {
		//Arrange
		guia := application.ServicoGuiaDeRemessa()

		//Atc
		guia.ExpedirMercadoria()

		//Assert
		if guia == nil {
			t.Errorf("A Guia de Remessa esta vazia")
		}
	})

	t.Run("Verificar se foi criado uma guia de Remessa", func(t *testing.T) {
		//Arrange
		guia := application.ServicoGuiaDeRemessa()

		//Act
		guia.NovoGuiaDeRemessa()

		//Assert
		if guia.NovoGuiaDeRemessa() != 1 {
			t.Errorf("A Guia não foi criada ")
		}
	})
	
}
