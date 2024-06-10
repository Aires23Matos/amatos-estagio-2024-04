package dominio

type Estoque interface {
	BuscarProdutos(ids ...string) []Produto
}
