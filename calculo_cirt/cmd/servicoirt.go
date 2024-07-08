package cmd

import (
	"fmt"
	"servicoirt/application/usecase"

	"github.com/spf13/cobra"
)

var salariobase float64
var subsidiodealimentacao float64
var subsidiotransporte float64


var calcularCmd = &cobra.Command{
    Use:   "cirt",
    Short: "cirt",
    Run: func(cmd *cobra.Command, args []string) {
        
        service := usecase.NewServicoIRT()


        inss := service.CalcularDescontodeSegurancaSocial(salariobase, subsidiodealimentacao, subsidiotransporte)
		totalsalariobase := service.SalarioBase(salariobase)
		totalsubsidioalimentacao,_:=service.SubsidioDeAlimentacao(salariobase)
		totalsubsidiodetransporte,_:=service.SubsidioDeTransporte(salariobase)
		calcularirt,_ :=service.CalcularIRT(salariobase,subsidiodealimentacao,subsidiotransporte)
		calcularsalarioliquido:=service.CalcularoSalarioLiquido(salariobase,subsidiodealimentacao,subsidiotransporte)
      
       
		fmt.Println("== Incrementos ==")
        fmt.Printf("Sub Alimentação: %.2f\n", totalsubsidioalimentacao)
        fmt.Printf("Sub Transporte: %.2f\n", totalsubsidiodetransporte)
        // Incluir outros incrementos se necessário

        fmt.Println("\n== Descontos ==")
        fmt.Printf("Salário base: %.2f\n", totalsalariobase)
        fmt.Printf("Segurança Social: %.2f\n", inss)
        fmt.Printf("IRT: %.2f\n", calcularirt)

        fmt.Println("\nSalário Líquido:", calcularsalarioliquido)
    },
    
}

func init() {
    rootCmd.AddCommand(calcularCmd)
	calcularCmd.Flags().Float64Var(&salariobase, "salariobase", 0.0, "Salário Base Mensal")
    calcularCmd.Flags().Float64Var(&subsidiodealimentacao, "subsidiodealimentacao", 0.0, "Valor do Subsídio de Alimentação")
    calcularCmd.Flags().Float64Var(&subsidiotransporte, "subsidiotransporte", 0.0, "Valor do Subsídio de Transporte")

}