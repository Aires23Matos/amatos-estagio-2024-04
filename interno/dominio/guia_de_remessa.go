package dominio

var GuiaDeRemessaInvalido = GuiaDeRemessa{}

type LinhaDeGuiaDeRemessa struct {
	IdProduto string
	descricao string
}

type GuiaDeRemessa struct {
	ID       string
	Mensagem string
	Linhas   []LinhaDeGuiaDeRemessa
}

func (g *GuiaDeRemessa) IDdados() string {
	return g.ID
}

func (g *GuiaDeRemessa) Linha() []LinhaDeGuiaDeRemessa {
	return g.Linhas
}

func (g *GuiaDeRemessa) constroiLinhas(produtos []Produto) {
	for _, p := range produtos {

		l := LinhaDeGuiaDeRemessa{p.ID(), p.Descricao()}

		g.Linhas = append(g.Linhas, l)
	}
}

func NovaGuiaDeRemessa(id string, produtos []Produto) GuiaDeRemessa {
	n := GuiaDeRemessa{ID: id, Linhas: make([]LinhaDeGuiaDeRemessa, 0)}

	n.constroiLinhas(produtos)

	return n
}
