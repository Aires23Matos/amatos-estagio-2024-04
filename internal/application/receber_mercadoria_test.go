package application_test

import (
	"Amc/internal/application"
	"Amc/internal/domain"
	"errors"
	"testing"
)

func TestReceberMercadoria(t *testing.T) {

	t.Run("Deve criar uma nota de recebimento", func(t *testing.T) {
		// Arrange
		catalogo := AcmeCatalogo()
		repo := NotaRecebimentoRepositorioFalso()
		servico := application.AcmeServico(repo, catalogo)

		// Act.
		servico.ReceberMercadoria("P1")

		// Assert
		if repo.Tamanho() == 0 {
			t.Error("Repositorio de notas de recebimento vazio")
		}
	})

	t.Run("Verifica se nota de recebimento foi criada", func(t *testing.T) {
		// Arrange
		catalogo := AcmeCatalogo()
		repo := NotaRecebimentoRepositorioFalso()
		servico := application.AcmeServico(repo, catalogo)

		// Act.
		servico.ReceberMercadoria("P1")

		// Assert
		if repo.Tamanho() != 1 {
			t.Errorf("A nota não foi criada %v", repo.Tamanho())
		}
	})

	t.Run("Deve criar linhas referentes aos produtos recebidos", func(t *testing.T) {
		// Arrange
		catalogo := AcmeCatalogo()
		repo := NotaRecebimentoRepositorioFalso()
		servico := application.AcmeServico(repo, catalogo)

		// Act.
		servico.ReceberMercadoria("P1", "P2", "P3")

		// Assert
		nota, _ := repo.Buscar("some-id")

		p1 := nota.Linhas()[0]
		p2 := nota.Linhas()[1]
		p3 := nota.Linhas()[2]

		if len(nota.Linhas()) != 3 {
			t.Errorf("Nota de recebimento com as linhas vazias %v", nota.Linhas())
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

type notaRecebimentoRepositorioFalso struct {
	notas []domain.NotaDeRecebimento
}

type FolhaDeCatalogo struct {
	produtos map[string]domain.Produto
}

func (f *FolhaDeCatalogo) BuscarProdutos(ids ...string) []domain.Produto {
	produtos := make([]domain.Produto, 0)
	for _, id := range ids {
		produtos = append(produtos, f.produtos[id])
	}
	return produtos
}

func AcmeCatalogo() *FolhaDeCatalogo {
	produtos := make(map[string]domain.Produto)
	produtos["P1"] = domain.NovoProduto("P1", "Leite")
	produtos["P2"] = domain.NovoProduto("P2", "Manteiga")
	produtos["P3"] = domain.NovoProduto("P3", "Iogurte")
	return &FolhaDeCatalogo{produtos: produtos}
}

func (r *notaRecebimentoRepositorioFalso) Tamanho() int {
	return len(r.notas)
}

func (r *notaRecebimentoRepositorioFalso) Guardar(n domain.NotaDeRecebimento) {
	r.notas = append([]domain.NotaDeRecebimento{}, n)
}

func (r *notaRecebimentoRepositorioFalso) Buscar(id string) (domain.NotaDeRecebimento, error) {
	for _, n := range r.notas {
		if n.ID() == id {
			return n, nil
		}
	}
	return domain.NotaDeRecebimentoInvalida, errors.New("Nota de recebimento não encontrada")
}

func NotaRecebimentoRepositorioFalso() domain.NotaRecebimentoRepositorio {
	return &notaRecebimentoRepositorioFalso{
		make([]domain.NotaDeRecebimento, 0),
	}
}
