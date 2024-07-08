package test_test

import (
	"servicoirt/application/usecase"
	"testing"
)

func TestCalularSalarioBaseAposFalta(t *testing.T) {
	t.Run("verificar se foi inserido o valor do salário base", func(t *testing.T) {
		//Arrange
		salario := 70.000
		servico := usecase.NewServicoIRT()
		//Act
		err := servico.SalarioBase(salario)
		//Assert
		if err <= 0 {
			t.Error("Não é permitido valores negativos")
		}
		if err < 1 {
			t.Errorf("Não foi digitado nenhum salário")
		}

	})

	t.Run("Deve aceitar um salário mínimo de 70 mil kz", func(t *testing.T) {
		//Arrange
		salario := 70.000
		servico := usecase.NewServicoIRT()
		//Act
		err := servico.SalarioBase(salario)
		//Assert
		if err < 70.000 {
			t.Errorf("salário inserido está a baixo do salário aceite. salario: %.3f", err)
		}
	})
	t.Run("Deve calcular o salário diario", func(t *testing.T) {
		//Arrange
		diastrabalho := 22
		servico := usecase.NewServicoIRT()
		//Act
		salariopordia := servico.SalarioDiario(float64(diastrabalho))
		//Assert
		if salariopordia < 1 {
			t.Errorf("Não foi  encontrado nenhum salário diario %.2f", salariopordia)
		}
	})

	t.Run("deve Descontar os dias após as faltas", func(t *testing.T) {
		//Arrange
		diasdetrabalho := 22
		diasdefaltas := 0
		servico := usecase.NewServicoIRT()
		//Act
		dias, zero := servico.DiasAposFalta(diasdetrabalho, diasdefaltas)
		//Assert
		if dias <= 0 {
			t.Error("Não foi encontrado dias com faltas", dias)
		}
		if diasdefaltas == diasdetrabalho {
			t.Error("Os dias de faltas é o mesmo dos dias de trabalho: ", zero)
		}
	})
	t.Run("Calcular salário base após a falta", func(t *testing.T) {
		//Arrange
		dias := 22
		faltas := 5
		salario := 70.000
		servico := usecase.NewServicoIRT()
		//Act
		salariobase := servico.SalarioBaseAposFalta(dias, faltas, salario)
		//Assert
		if salariobase < 0 {
			t.Errorf("O salário base após as faltas não encontrado. Salário: %.3f", salariobase)
		}
		if faltas == 0 {
			t.Errorf("Não foi inserido nenhuma falta. faltas: %v", faltas)
		}
	})
}

func TestCalcularSubsidioDeAlimentacao(t *testing.T) {
	var diasdetrabalho = 22
	var diasdefaltas = 5
	var subsidiodealimentacao = 35.000
	t.Run("deve inserir valor do subsídio de Alimentação", func(t *testing.T) {
		//Arrange
		subsidio := usecase.NewServicoIRT()
		//Act
		valordesubsidio, _ := subsidio.SubsidioDeAlimentacao(subsidiodealimentacao)
		//Assert
		if valordesubsidio < 1 {
			t.Errorf("O subsidio de Alimentação não foi inserido corretamente %.3f kz", valordesubsidio)
		}
	})

	t.Run("Deve dar erro quando é inserido um valor negativo", func(t *testing.T) {
		//Arrange
		subsidio := usecase.NewServicoIRT()
		//Act
		valor, mensagem := subsidio.SubsidioDeAlimentacao(subsidiodealimentacao)
		//Assert
		if valor == 0 {
			t.Errorf("Erro: %v", mensagem)
		}
	})

	t.Run("calcular o quanto dinheiro do subsidio pode receber por dia", func(t *testing.T) {
		//Arrange
		subsidio := usecase.NewServicoIRT()
		//Act
		subsidiopordia := subsidio.CalcularSubsidoPorDia(subsidiodealimentacao, diasdetrabalho)
		//Assert
		if subsidiopordia <= 0 {
			t.Errorf("Não teve nenhum dinheiro do subsídio de alimentação ganho por dia. subsídio por dia: %.3f", subsidiopordia)
		}
	})
	t.Run("Deve calcular o subsidio de Alimentação", func(t *testing.T) {
		//rrange
		subsidio := usecase.NewServicoIRT()
		//Act
		calculosubsidio := subsidio.CalcularSubsidiodeAlimentacaoNosDiasUteis(subsidiodealimentacao, diasdetrabalho, diasdefaltas)
		//Assert
		if calculosubsidio <= 0 {
			t.Errorf("Não foi adicionado nenhum subsidio de alimentação. valor do Subsídio: %.3f", calculosubsidio)
		}
	})
}
func TestCalcularSubsidiodeTransporte(t *testing.T) {
	diasdetrabalho := 22
	faltas := 5
	subsidiodetransporte := 70.000
	t.Run("Deve inserir o valor do subsídio de transporte", func(t *testing.T) {
		//Arrange
		transporte := usecase.NewServicoIRT()
		//Act
		subsidio, _ := transporte.SubsidioDeTransporte(subsidiodetransporte)
		//Assert
		if subsidio < 1 {
			t.Errorf("não foi inserido o valor do subídio de transporte. Valor : %.1f", subsidio)
		}
	})

	t.Run("Deve dar erro quando é inserido um valor negativo", func(t *testing.T) {
		//Arrange
		transporte := usecase.NewServicoIRT()
		//Act
		valor, mensagem := transporte.SubsidioDeTransporte(subsidiodetransporte)
		//Assert
		if valor == 0 {
			t.Errorf("Erro: %v", mensagem)
		}
	})

	t.Run("Calcular o dinheiro do subsidio de transporte ganho por dia", func(t *testing.T) {
		//Arrange
		transporte := usecase.NewServicoIRT()
		//Act
		subsidiodetransportepordia := transporte.CalcularSubsidioDeTransportePorDia(subsidiodetransporte, diasdetrabalho)
		//Assert
		if subsidiodetransportepordia <= 0 {
			t.Errorf("Não teve nenhum dinheiro do subsídio de transporte ganho por dia. valor por dia: %3f", subsidiodetransportepordia)
		}
	})
	t.Run("Deve calcular o subsídio de Transporte", func(t *testing.T) {
		//rrange
		transporte := usecase.NewServicoIRT()
		//Act
		calculosubsidiodetransporte := transporte.CalcularSubsidiodeTransporteNosDiasUteis(subsidiodetransporte, diasdetrabalho, faltas)
		//Assert
		if calculosubsidiodetransporte <= 0{
			t.Errorf("Não foi adicionado nenhum subsidio de Transporte. valor do Subsídio: %.3fkz", calculosubsidiodetransporte)
		}
	})
}
