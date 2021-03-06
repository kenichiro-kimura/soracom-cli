package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(ProductsCmd)
}

// ProductsCmd defines 'products' subcommand
var ProductsCmd = &cobra.Command{
	Use:   "products",
	Short: TRCLI("cli.products.summary"),
	Long:  TRCLI(`cli.products.description`),
}
