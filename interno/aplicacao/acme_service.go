package application

import (
	"Acme/interno/dominio"
)



type GuiDeRemessaServico struct {
	repo dominio.GuiaDeRemessaRepositorio
	Estoque dominio.Estoque
}

func NovoGuiaDeRemessa(repo dominio.GuiaDeRemessaRepositorio, c dominio.Estoque) *GuiDeRemessaServico {
	return &GuiDeRemessaServico{
		repo: repo,
		Estoque: c,
	}
}

func (g *GuiDeRemessaServico) CriarGuiaDeRemessa() string {
	Guia := dominio.GuiaDeRemessa{Mensagem: "Guia de Remessa Criada"}
	g.repo.Salvar(Guia)
	return Guia.Mensagem
}

func (g *GuiDeRemessaServico) ExpedirMercadoria(p ...string) {
	produtos := g.Estoque.BuscarProdutos(p...)

	n := dominio.NovaNotaDeRecebimento("some-id", produtos)

	g.repo.Salvar(n)
}
