package service

import "errors"


type GuiaRemessa struct{
	criada bool
}

func CriarGuiaDeRemessa() *GuiaRemessa{
	return &GuiaRemessa{}
}

func (g * GuiaRemessa) NovaGuiaDeRemessa() (string, error){
	if !g.criada {
		return "", errors.New("crie uma guia de remmesa antes de criar uma nova guia ")
	}
	return "Nova Guia de Remessa Criada", nil
}

func (g * GuiaRemessa) Criar()error{
	g.criada = true
	return nil
}