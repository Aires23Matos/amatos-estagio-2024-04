package domain

type Catalogo interface {
	BuscarProdutos(ids ...string) []Produto
}
