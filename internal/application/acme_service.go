package application

type AcmeGuia struct {
	//repo domain.GuiaDeRemessarepositorio
}

func ServicoGuiaDeRemessa() *AcmeGuia {
	return &AcmeGuia{}
}

func (g *AcmeGuia) NovoGuiaDeRemessa() int {
	return 1
}

func (g *AcmeGuia) ExpedirMercadoria() {

}
