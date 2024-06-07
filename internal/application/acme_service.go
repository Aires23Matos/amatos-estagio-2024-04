package application

import "Acme/internal/domain"

type GuiDeRemessaServico struct {
	repo domain.GuiaDeRemessaRepositorio
}

func NovoGuiaDeRemessa(repo domain.GuiaDeRemessaRepositorio) *GuiDeRemessaServico {
	return &GuiDeRemessaServico{
		repo: repo,
	}
}

func (g *GuiDeRemessaServico) CriarGuiaDeRemessa() string {
	Guia := domain.GuiaDeRemessa{Mensagem: "Guia de Remessa Criada"}
	g.repo.Salvar(Guia)
	return Guia.Mensagem
}

func (g *GuiDeRemessaServico) ExpedirMercadoria() {
	Guia := domain.GuiaDeRemessa{ID: "1"}
	g.repo.Salvar(Guia)
}
