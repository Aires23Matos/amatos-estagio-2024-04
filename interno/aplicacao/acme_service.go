package application

import (
	"Acme/interno/dominio"
)

type GuiaDeRemessaServico struct {
	repo    dominio.GuiaDeRemessaRepositorio
	Estoque dominio.Estoque
}

func NovoGuiaDeRemessa(repo dominio.GuiaDeRemessaRepositorio, c dominio.Estoque) *GuiaDeRemessaServico {
	return &GuiaDeRemessaServico{
		repo:    repo,
		Estoque: c,
	}
}

func (g *GuiaDeRemessaServico) CriarGuiaDeRemessa() string {
	Guia := dominio.GuiaDeRemessa{ID: "qualquer-id", Mensagem: "Guia de Remessa Criada"}
	g.repo.Salvar(Guia)
	return Guia.Mensagem
}


func (g *GuiaDeRemessaServico) ExpedirMercadoria(p ...string) {
	produtos := g.Estoque.BuscarProdutos(p...)

	n := dominio.NovaGuiaDeRemessa("some-id", produtos)

	g.repo.Salvar(n)
}

func (g *GuiaDeRemessaServico) MercadoriaExpedida(idProduto string) bool {
	guias := g.repo.Todas()
	for _, guia := range guias {
		for _, linha := range guia.Linha() {
			if linha.IdProduto == idProduto {
				return true
			}
		}
	}
	return false
}