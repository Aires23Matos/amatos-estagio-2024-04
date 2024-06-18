package application_test

import (
	"cvl/application"
	"cvl/domain"
	"testing"
)

func TestInternarParciente(t *testing.T) {
	//Arrange
	repo := domain.NewFichaInternementorepository()
	service := application.CVLService(repo)

	//Act
	service.FichaDeInternamento()
	//Assert

	if repo.Tamanho() == 0 {
		t.Errorf("Ficha de Internamento está vazio")
	}
	if repo.Tamanho() != 1 {
		t.Errorf("Ficha de Internamento não criada %v",repo.Tamanho())
	}

	_, err := repo.Buscar("1"); 

	if err != nil {
        t.Errorf("Erro ao buscar a ficha de internamento criada: %v", err)
    }
}
