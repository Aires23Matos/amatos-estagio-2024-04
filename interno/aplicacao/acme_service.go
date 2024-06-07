package aplicacao

import "Acme/interno/dominio"

type GuiDeRemessaServico struct {
	repo dominio.GuiaDeRemessaRepositorio
}

func NovoGuiaDeRemessa(repo dominio.GuiaDeRemessaRepositorio) *GuiDeRemessaServico {
	return &GuiDeRemessaServico{
		repo: repo,
	}
}

func (g *GuiDeRemessaServico) CriarGuiaDeRemessa() string {
	Guia := dominio.GuiaDeRemessa{Mensagem: "Guia de Remessa Criada"}
	g.repo.Salvar(Guia)
	return Guia.Mensagem
}

func (g *GuiDeRemessaServico) ExpedirMercadoria() {
	Guia := dominio.GuiaDeRemessa{ID: "1"}
	g.repo.Salvar(Guia)
}
