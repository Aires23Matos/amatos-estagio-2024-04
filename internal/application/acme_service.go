package application

import "Amc/internal/domain"

type acme struct {
	repo     domain.NotaRecebimentoRepositorio
	catalogo domain.Catalogo
}

func AcmeServico(r domain.NotaRecebimentoRepositorio, c domain.Catalogo) *acme {
	return &acme{repo: r, catalogo: c}
}

func (s acme) ReceberMercadoria(p ...string) {
	produtos := s.catalogo.BuscarProdutos(p...)

	n := domain.NovaNotaDeRecebimento("some-id", produtos)

	s.repo.Guardar(n)
}
