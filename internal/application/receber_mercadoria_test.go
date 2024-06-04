package application_test

import (
	service "Amc/internal/application"
	"Amc/internal/domain"
	"testing"
)

func TestReceberMercadoria(t *testing.T) {

	t.Run("Deve criar uma nota de recebimento", func(t *testing.T) {
		// Arrange
		repo := NovoNotaRecebimentoRepositoryFalso()
		service := service.NovoAcmeService()

		// Act.
		service.ReceberMercadoria()

		// Assert
		if repo.Tamanho() == 0 {
			t.Error("Repositorio de notas de recebimento vazio")
		}
	})

	t.Run("", func(t *testing.T) {
		
	
	})

}

type NotaRecebimentoRepositoryFalso struct{}

func (r NotaRecebimentoRepositoryFalso) Tamanho() int {
	return 1
}

func NovoNotaRecebimentoRepositoryFalso() domain.NotaRecebimentoRepository {
	return NotaRecebimentoRepositoryFalso{}
}
