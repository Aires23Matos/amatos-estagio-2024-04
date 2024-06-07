package application_test

import (
	servico "Acme/internal/application"
	"Acme/internal/domain"
	"testing"
)

func TestExpedirMercadoria(t *testing.T) {

	t.Run("Deve criar uma Guia de Remessa",func(t *testing.T) {
		//Act
		repo := domain.NovoExpedirMercadoriaRepositorio()
		guia := servico.NovoGuiaDeRemessa(repo)
		//Assert
		if guia == nil{
			t.Error("Guia de Remessa Não foi Criado")
		}
	})

	t.Run("Verificar se existe Guia de Remessa",func(t *testing.T) {
		//Arrange
		repo := domain.NovoExpedirMercadoriaRepositorio()
		guia := servico.NovoGuiaDeRemessa(repo)

		//Act
		guia.CriarGuiaDeRemessa()

		//Assert
		if guia.CriarGuiaDeRemessa() != "Guia de Remessa Criada"{
			t.Error("Guia de Remessa Não foi Criada")
		}
	})

	t.Run("verificar se existe um repositorio do Guia de Remessa",func(t *testing.T) {
		//Arrange
		repo := domain.NovoExpedirMercadoriaRepositorio()
		guia := servico.NovoGuiaDeRemessa(repo)

		//Act
		guia.CriarGuiaDeRemessa()

		//Assert
		if repo.Tamanho() == 0{
			t.Error("Repositorio Guia de Remessa vazio")
		}
	})

	t.Run("verificar se o repositorio guia de Remessa foi criado", func(t *testing.T) {
		//Arrange
		repo := domain.NovoExpedirMercadoriaRepositorio()
		guia := servico.NovoGuiaDeRemessa(repo)
		//Act

		guia.ExpedirMercadoria()
		//Assert
		if !repo.Ultima("1"){
			t.Error("Guia de Remessa não foi encontrado")
		}

	})
}


