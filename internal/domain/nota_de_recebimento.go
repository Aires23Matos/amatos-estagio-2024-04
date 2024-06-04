package domain

var NotaDeRecebimentoInvalida = NotaDeRecebimento{}

type LinhaDeRecebimento struct {
	IdProduto string
	descricao string
}

type NotaDeRecebimento struct {
	id     string
	linhas []LinhaDeRecebimento
}

func (n *NotaDeRecebimento) ID() string {
	return n.id
}

func (n *NotaDeRecebimento) Linhas() []LinhaDeRecebimento {
	return n.linhas
}

func (n *NotaDeRecebimento) constroiLinhas(produtos []Produto) {
	for _, p := range produtos {

		l := LinhaDeRecebimento{p.ID(), p.Descricao()}

		n.linhas = append(n.linhas, l)
	}
}

func NovaNotaDeRecebimento(id string, produtos []Produto) NotaDeRecebimento {
	n := NotaDeRecebimento{id: id, linhas: make([]LinhaDeRecebimento, 0)}

	n.constroiLinhas(produtos)

	return n
}
