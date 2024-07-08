package usecase

import (
	"errors"
)

type TabelaIrt struct {
	Escalao      int
	ValorInicial float64
	ValorFinal   float64
	ParcelaFixa  float64
	Taxa         float64
	Excesso      float64
}

type ServiceIRT struct {
}

func NewServicoIRT() *ServiceIRT {
	return &ServiceIRT{}
}
func Dados() []TabelaIrt {
	return []TabelaIrt{
		{ ValorInicial: 0, ValorFinal: 100000, ParcelaFixa: 0, Taxa: 0, Excesso: 0},
		{ ValorInicial: 100001, ValorFinal: 150000, ParcelaFixa: 0, Taxa: 0.13, Excesso: 100001},
		{ ValorInicial: 150001, ValorFinal: 200000, ParcelaFixa: 12.500, Taxa: 0.16, Excesso: 150001},
		{ ValorInicial: 200001, ValorFinal: 300000, ParcelaFixa: 31.250, Taxa: 0.18, Excesso: 200001},
		{ ValorInicial: 300001, ValorFinal: 500000, ParcelaFixa: 49.250, Taxa: 0.19, Excesso: 300001},
		{ ValorInicial: 500001, ValorFinal: 1000000, ParcelaFixa: 87.250, Taxa: 0.20, Excesso: 500001},
		{ ValorInicial: 1000001, ValorFinal: 1500000, ParcelaFixa: 187.249, Taxa: 0.21, Excesso: 1000001},
		{ ValorInicial: 1500001, ValorFinal: 2000000, ParcelaFixa: 292.249, Taxa: 0.22, Excesso: 1500001},
		{ ValorInicial: 2000001, ValorFinal: 2500000, ParcelaFixa: 402.249, Taxa: 0.23, Excesso: 2000001},
		{ ValorInicial: 2500001, ValorFinal: 5000000, ParcelaFixa: 517.249, Taxa: 0.24, Excesso: 2500001},
		{ ValorInicial: 5000001, ValorFinal: 10000000, ParcelaFixa: 1117.249, Taxa: 24.5, Excesso: 5000001},
		{ ValorInicial: 10000001, ValorFinal: 0, ParcelaFixa: 2342248, Taxa: 25.0, Excesso: 10000001},
	}
}
func (s *ServiceIRT) SalarioBase(salariomensal float64) float64 {
	var salario = salariomensal
	return salario
}

func (s *ServiceIRT) SalarioDiario(dias float64) float64 {
	diasdetarbalho := dias
	salariodiario := s.SalarioBase(70.000) / float64(diasdetarbalho)
	return salariodiario
}
func (s *ServiceIRT) DiasAposFalta(dias, faltas int) (int, error) {
	diasdetarbalho := dias
	diasdefalta := faltas
	if diasdetarbalho == diasdefalta {
		return diasdefalta - diasdetarbalho, errors.New("não compareceu nos ultimos 22 dias")
	}
	diasdetabalhoaposfalta := diasdetarbalho - diasdefalta
	return diasdetabalhoaposfalta, nil
}

func (s *ServiceIRT) SalarioBaseAposFalta(dias, faltas int, salario float64) float64 {
	salariobase := s.SalarioBase(salario)
	salariopordia := s.SalarioDiario(float64(dias))
	salarioaposfalta := salariobase - (salariopordia * float64(faltas))
	return salarioaposfalta
}

func (s *ServiceIRT) SubsidioDeAlimentacao(valordosubsidiodealimentacao float64) (float64, error) {
	subsidio := valordosubsidiodealimentacao
	if subsidio < 0 {
		return 0, errors.New("foi inserido um valor negativo")
	}
	return subsidio, nil
}

func (s *ServiceIRT) CalcularSubsidoPorDia(subsidio float64, dias int) float64 {
	subsidioalimentacao := subsidio
	diasdetarbalho := dias
	subsidiodiario, _ := s.SubsidioDeAlimentacao(subsidioalimentacao)
	calculosubsidiopordia := subsidiodiario / float64(diasdetarbalho)
	return calculosubsidiopordia
}

func (s *ServiceIRT) CalcularSubsidiodeAlimentacaoNosDiasUteis(subsidio float64, totaldosdiasdetrabalho, faltas int) float64 {
	valordosubsidio := subsidio
	diasuteis := totaldosdiasdetrabalho
	valordaalimentacao, _ := s.SubsidioDeAlimentacao(valordosubsidio)
	valorpordia := s.CalcularSubsidoPorDia(valordosubsidio, diasuteis)
	diasrealizados, _ := s.DiasAposFalta(diasuteis, faltas)
	calcularsubsidio := valordaalimentacao - (valorpordia * float64(diasrealizados))
	return calcularsubsidio
}

func (s *ServiceIRT) SubsidioDeTransporte(valordosubsidiodetransport float64) (float64, error) {
	subsidiodetransporte := valordosubsidiodetransport
	if subsidiodetransporte < 0 {
		return 0, errors.New("foi inserido valor negativo")
	}
	return subsidiodetransporte, nil
}

func (s *ServiceIRT) CalcularSubsidioDeTransportePorDia(subsidio float64, dias int) float64 {
	sudsidiodetransport := subsidio
	diasdetrabalho := dias
	subsidiodiario, _ := s.SubsidioDeTransporte(sudsidiodetransport)
	calcularsubsidiotransportpordia := subsidiodiario / float64(diasdetrabalho)
	return calcularsubsidiotransportpordia
}

func (s *ServiceIRT) CalcularSubsidiodeTransporteNosDiasUteis(subsidio float64, totaldosdiasdetrabalho, faltas int) float64 {
	valordosubsidio := subsidio
	diasuteis := totaldosdiasdetrabalho
	valordotransporte, _ := s.SubsidioDeTransporte(valordosubsidio)
	valorpordia := s.CalcularSubsidioDeTransportePorDia(valordosubsidio, diasuteis)
	diasrealizados, _ := s.DiasAposFalta(diasuteis, faltas)
	calcularsubsidio := valordotransporte - (valorpordia * float64(diasrealizados))
	return calcularsubsidio
}

func (s *ServiceIRT) CalculoDoSalarioBruto(salariobase, subsidiodealimentacao, subsidiotransporte float64) float64 {
	totaldosalariobase := salariobase
	totaldosubsidiodealimentacao := subsidiodealimentacao
	totaldosubsidiodetransporte := subsidiotransporte
	alimentacao, _ := s.SubsidioDeAlimentacao(totaldosubsidiodealimentacao)
	transporte, _ := s.SubsidioDeTransporte(totaldosubsidiodetransporte)
	salariobruto := s.SalarioBase(totaldosalariobase) + alimentacao + transporte
	return salariobruto
}

func (s *ServiceIRT) CalcularDescontodeSegurancaSocial(salariobase, subsidioalimentacao, subsidiotransporte float64) float64 {
	totaldosalariobase := salariobase
	totaldosubsidiodealimentacao := subsidioalimentacao
	totaldosubsidiodetransporte := subsidiotransporte

	descontosdesegurancasocial := s.CalculoDoSalarioBruto(totaldosalariobase, totaldosubsidiodealimentacao, totaldosubsidiodetransporte) * 0.03
	return descontosdesegurancasocial
}

func (s *ServiceIRT) CalcularExcessodoSubsidioAlimentacao(subsidioalimentacao float64) (float64, error) {
	limitesubsidio := 30.000
	totaldosutotaldosubsidiodealimentacao := subsidioalimentacao
	alimentacao, _ := s.SubsidioDeAlimentacao(totaldosutotaldosubsidiodealimentacao)
	excessoalimentacao := alimentacao - limitesubsidio
	if excessoalimentacao < 0 {
		converter := -1 * (excessoalimentacao)
		return converter, nil
	}
	if alimentacao < limitesubsidio && alimentacao > 0 {
		return alimentacao, nil
	}
	if excessoalimentacao == 0 {
		return alimentacao, errors.New("não passou o limite então não está sujeito a irt")
	}
	return excessoalimentacao, nil
}

func (s *ServiceIRT) CalcularExcessodoSubsidioTransporte(subsidiotransporte float64) (float64, error) {
	limitesubsidio := 30.000
	totaldosubsidiodetransporte := subsidiotransporte
	transporte, _ := s.SubsidioDeTransporte(totaldosubsidiodetransporte)
	excessotransporte := limitesubsidio - transporte
	if excessotransporte < 0 {
		converter := -1 * (excessotransporte)
		return converter, nil
	}
	if transporte < limitesubsidio && transporte > 0 {
		return transporte, nil
	}
	if excessotransporte == 0 {
		return transporte, errors.New("não passou o limite então não está sujeito a irt")
	}
	return excessotransporte, nil
}

func (s *ServiceIRT) CalcularTotalSujeitoIRT(subsidioalimentacao, subsidiotransporte float64) float64 {
	totalexcessodealimentacao := subsidioalimentacao
	totalexcessodetransporte := subsidiotransporte
	excessoalimentacao, _ := s.CalcularExcessodoSubsidioAlimentacao(totalexcessodealimentacao)
	excessotransporte, _ := s.CalcularExcessodoSubsidioTransporte(totalexcessodetransporte)
	totalsujeitoirt := excessoalimentacao + excessotransporte
	return totalsujeitoirt
}
func (s *ServiceIRT) CalcularMaterialColetado(salariobase, subsidioalimentacao, subsidiotransporte float64) float64 {
	totaldosalariobase := salariobase
	totalexcessodealimentacao := subsidioalimentacao
	totalexcessodetransporte := subsidiotransporte
	salario := s.SalarioBase(totaldosalariobase)
	sujeitoirt := s.CalcularTotalSujeitoIRT(totalexcessodealimentacao, totalexcessodetransporte)
	inss := s.CalcularDescontodeSegurancaSocial(totaldosalariobase, totalexcessodealimentacao, totalexcessodetransporte)
	mc := (salario + sujeitoirt) - inss
	return mc
}

func (s *ServiceIRT) CalcularIRT(salariobase, subsidioalimentacao, subsidiotransporte float64) (float64, error) {
	totaldosalariobase := salariobase
	totalexcessodealimentacao := subsidioalimentacao
	totalexcessodetransporte := subsidiotransporte

	salario := s.CalcularMaterialColetado(totaldosalariobase, totalexcessodealimentacao, totalexcessodetransporte)
	tabelas := Dados()

	for _, tabela := range tabelas {
		if (tabela.ValorFinal == 0 && salario >= tabela.ValorInicial) || (salario >= tabela.ValorInicial && salario <= tabela.ValorFinal) {
			calculoIrt := (tabela.ParcelaFixa + (salario - tabela.Excesso)) * tabela.Taxa 
			return calculoIrt, nil
		}
	}

	return 0, errors.New("salário fora da faixa de imposto")
}

func (s *ServiceIRT)CalcularoSalarioLiquido(salariobase, subsidioalimentacao, subsidiotransporte float64)float64{
	totaldosalariobase := salariobase
	totalexcessodealimentacao := subsidioalimentacao
	totalexcessodetransporte := subsidiotransporte

	salarioirt,_ := s.CalcularIRT(totaldosalariobase, totalexcessodealimentacao,totalexcessodetransporte)
	salarioinss := s.CalcularDescontodeSegurancaSocial(totaldosalariobase, totalexcessodealimentacao,totalexcessodetransporte)
	total := salarioirt + salarioinss

	salarioliquido := s.CalculoDoSalarioBruto(totaldosalariobase, totalexcessodealimentacao, totalexcessodetransporte) + total
	return salarioliquido
}
