package cmd

import (
	aplicacao "Acme/interno/aplicacao"
	"Acme/interno/dominio"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var expedirCmd = &cobra.Command{
	Use:   "expedir [IDs dos produtos]",
	Short: "Expede mercadorias",
	Long:  `Expede mercadorias fornecendo os IDs dos produtos.`,
	Run: func(cmd *cobra.Command, args []string) {

		estoque := dominio.AcmeEstoque()
		repo := dominio.NovoExpedirMercadoriaRepositorio()
		servico := aplicacao.NovoGuiaDeRemessa(repo, estoque)

		if len(args) < 3 {
			fmt.Println("Você precisa inserir pelo menos 3 IDs de produtos.")
			return
		}

		ids := strings.Join(args, ", ")
		fmt.Printf("Expedindo mercadorias: %s\n", ids)
		servico.ExpedirMercadoria(args...)

		fmt.Println("Mercadorias expedidas:")
		guias := repo.Todas()
		for _, guia := range guias {
			fmt.Printf("Guia ID: %s\n", guia.IDdados())
			for _, linha := range guia.Linha() {
				fmt.Printf(" - Produto ID: %s\n", linha.IdProduto)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(expedirCmd)
}