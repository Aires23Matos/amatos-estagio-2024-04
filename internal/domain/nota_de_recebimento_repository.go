package domain

type NotaRecebimentoRepositorio interface {
	Buscar(id string) (NotaDeRecebimento, error)
	Guardar(n NotaDeRecebimento)
	Tamanho() int
}
