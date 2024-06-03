package service_test

import (
	"Amc/internal/application/service"
	"testing"
)

func CriarGuiaDeRemessaTest(t *testing.T) {
	t.Run("Verifica se guia de remessa foi criada", func(t *testing.T) {
		guia := service.CriarGuiaDeRemessa()

		if guia == nil{
			t.Fatal("Erro ao Criar o Guia de Remmessa")
		}
		t.Logf("Guia de remessa criada com sucesso")
	})

	t.Run("Verificar se existe uma nova Guia de Remessa", func(t * testing.T){
		guia := service.CriarGuiaDeRemessa()

		guia.Criar()

		novaguia, err := guia.NovaGuiaDeRemessa()

		if err == nil{
			t.Fatalf("Esperava um erro ao criar nova guia de remessa antes de criar a guia de remessa")
		}

		if err != nil {
			t.Fatalf("Erro ao criar nova guia de remessa: %v", err)
		}

		if novaguia != "Nova Guia de Remessa criada"{
			t.Fatal("Resultado inesperado ao criar nova gia de remessa")
		}
		t.Logf("Nova guia de remessa criada com sucesso")
	})
}
