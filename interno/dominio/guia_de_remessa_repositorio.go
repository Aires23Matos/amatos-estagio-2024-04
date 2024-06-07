package dominio

import "errors"

type GuiaDeRemessaRepositorio interface {
	Tamanho() int
	Salvar(Guia GuiaDeRemessa)
	Buscar(id string) (GuiaDeRemessa, error)
}

type GuiaDeRemessaRepositorioFalso struct {
	dadosguia map[string]GuiaDeRemessa
}

type FolhaDeEstoque struct {
	produtos map[string]Produto
}

func (f *FolhaDeEstoque) BuscarProdutos(ids ...string) []Produto {
	produtos := make([]Produto, 0)
	for _, id := range ids {
		produtos = append(produtos, f.produtos[id])
	}
	return produtos
}

func AcmeEstoque() *FolhaDeEstoque {
	produtos := make(map[string]Produto)
	produtos["P1"] = NovoProduto("P1", "Mesa")
	produtos["P2"] = NovoProduto("P2", "Cadeiras")
	produtos["P3"] = NovoProduto("P3", "Panos")
	return &FolhaDeEstoque{produtos: produtos}
}

func (g *GuiaDeRemessaRepositorioFalso) Buscar(id string) (GuiaDeRemessa, error) {
	for _, n := range g.dadosguia {
		if n.IDdados() == id {
			return n, nil
		}
	}
	return GuiaDeRemessaInvalido, errors.New("guia de remessa não encontrada")
}

func (r GuiaDeRemessaRepositorioFalso) Tamanho() int {
	return len(r.dadosguia)
}

func (r *GuiaDeRemessaRepositorioFalso) Salvar(Guia GuiaDeRemessa) {
	r.dadosguia[Guia.ID] = Guia
}

func NovoExpedirMercadoriaRepositorio() GuiaDeRemessaRepositorio {
	return &GuiaDeRemessaRepositorioFalso{
		dadosguia: make(map[string]GuiaDeRemessa),
	}
}
