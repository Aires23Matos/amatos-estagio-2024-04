package dominio

type GuiaDeRemessaRepositorio interface {
	Tamanho() int
	Ultima(id string) bool
	Salvar(Guia GuiaDeRemessa)
}

type GuiaDeRemessaRepositorioFalso struct {
	dadosguia map[string]GuiaDeRemessa
}

func (r GuiaDeRemessaRepositorioFalso) Tamanho() int {
	return len(r.dadosguia)
}

func (r GuiaDeRemessaRepositorioFalso) Ultima(id string) bool {
	_, existe := r.dadosguia[id]
	return existe
}

func (r *GuiaDeRemessaRepositorioFalso) Salvar(Guia GuiaDeRemessa) {
	r.dadosguia[Guia.ID] = Guia
}

func NovoExpedirMercadoriaRepositorio() GuiaDeRemessaRepositorio {
	return &GuiaDeRemessaRepositorioFalso{
		dadosguia: make(map[string]GuiaDeRemessa),
	}
}
