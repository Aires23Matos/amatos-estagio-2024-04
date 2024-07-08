package usecase

import (
	"errors"
)

type ServiceIRT struct {
}

func NewServicoIRT() *ServiceIRT {
	return &ServiceIRT{}
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

func (s *ServiceIRT) CalcularSubsidiodeAlimentacao(subsidio float64, totaldosdiasdetrabalho, faltas int) float64 {
	valordosubsidio := subsidio
	diasuteis := totaldosdiasdetrabalho
	valordaalimentacao, _ := s.SubsidioDeAlimentacao(valordosubsidio)
	valorpordia := s.CalcularSubsidoPorDia(valordosubsidio, diasuteis)
	diasrealizados, _ := s.DiasAposFalta(diasuteis, faltas)
	calcularsubsidio := valordaalimentacao - (valorpordia * float64(diasrealizados))
	return calcularsubsidio
}

func (s *ServiceIRT)SubsidioDeTransporte(valordosubsidiodetransport float64)(float64, error){
	subsidiodetransporte := valordosubsidiodetransport
	if subsidiodetransporte < 0{
		return 0, errors.New("foi inserido valor negativo")
	}
	return subsidiodetransporte, nil
}

func(s *ServiceIRT) CalcularSubsidioDeTransportePorDia(){

}