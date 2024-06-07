package dominio

type Produto struct {
	id        string
	descricao string
}

func (p Produto) ID() string {
	return p.id
}

func (p Produto) Descricao() string {
	return p.descricao
}

func NovoProduto(id, descricao string) Produto {
	return Produto{id: id, descricao: descricao}
}
