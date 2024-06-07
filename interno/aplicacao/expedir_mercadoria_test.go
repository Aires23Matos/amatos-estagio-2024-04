package application_test

import (
	servico "Acme/interno/aplicacao"
	"Acme/interno/dominio"

	"testing"
)

func TestExpedirMercadoria(t *testing.T) {

	t.Run("Deve criar uma Guia de Remessa", func(t *testing.T) {
		//Act
		catalogo := dominio.AcmeEstoque()
		repo := dominio.NovoExpedirMercadoriaRepositorio()
		guia := servico.NovoGuiaDeRemessa(repo, catalogo)
		//Assert
		if guia == nil {
			t.Error("Guia de Remessa Não foi Criado")
		}
	})

	t.Run("Verificar se existe Guia de Remessa", func(t *testing.T) {
		//Arrange
		catalogo := dominio.AcmeEstoque()
		repo := dominio.NovoExpedirMercadoriaRepositorio()
		guia := servico.NovoGuiaDeRemessa(repo, catalogo)

		//Act
		guia.CriarGuiaDeRemessa()

		//Assert
		if guia.CriarGuiaDeRemessa() != "Guia de Remessa Criada" {
			t.Error("Guia de Remessa Não foi Criada")
		}
	})

	t.Run("verificar se existe um repositorio do Guia de Remessa", func(t *testing.T) {
		//Arrange
		catalogo := dominio.AcmeEstoque()
		repo := dominio.NovoExpedirMercadoriaRepositorio()
		guia := servico.NovoGuiaDeRemessa(repo, catalogo)

		//Act
		guia.CriarGuiaDeRemessa()

		//Assert
		if repo.Tamanho() == 0 {
			t.Error("Repositorio Guia de Remessa vazio")
		}
	})

	t.Run("verificar se o repositorio guia de Remessa foi criado", func(t *testing.T) {
		//Arrange
		estoque := dominio.AcmeEstoque()
		repo := dominio.NovoExpedirMercadoriaRepositorio()
		guia := servico.NovoGuiaDeRemessa(repo, estoque)
		//Act

		guia.ExpedirMercadoria()
		//Assert
		if repo.Tamanho() != 1 {
			t.Error("Guia de Remessa não foi encontrado")
		}

	})

	t.Run("Deve criar linhas referentes aos produtos na Guai De Remessa", func(t *testing.T) {
		// Arrange
		catalogo := dominio.AcmeEstoque()
		repo := dominio.NovoExpedirMercadoriaRepositorio()
		servico := servico.NovoGuiaDeRemessa(repo, catalogo)

		// Act.
		servico.ExpedirMercadoria("P1", "P2", "P3")

		// Assert
		guia, _ := repo.Buscar("some-id")

		p1 := guia.Linha()[0]
		p2 := guia.Linha()[1]
		p3 := guia.Linha()[2]

		if len(guia.Linha()) != 3 {
			t.Errorf("Nota de recebimento com as linhas vazias %v", guia.Linha())
		}
		if p1.IdProduto != "P1" {
			t.Errorf("ID do produto errado %v", p1.IdProduto)
		}
		if p2.IdProduto != "P2" {
			t.Errorf("ID do produto errado %v", p2.IdProduto)
		}
		if p3.IdProduto != "P3" {
			t.Errorf("ID do produto errado %v", p3.IdProduto)
		}
	})
}
