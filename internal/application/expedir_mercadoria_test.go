package application_test

import (
	"Acme/internal/application"
	"testing"
)

func TestEXpedirMercadoria(t *testing.T){
	t.Run("",func(t *testing.T) {
		//Arrange
		guia := application.NovoGuiaDeRemessa()

		//Atc
		guia.ExpedirMercadoria()

		//Assert
		if guia == nil{
			t.Errorf("A Guia de Remessa esta vazia")
		}
	})
}