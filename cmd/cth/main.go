package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var cthCmd = &cobra.Command{
		Use:   "cth",
		Short: "Cactus HUB CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cthCmd.AddCommand(versionCmd)
	cthCmd.AddCommand(BalancesCmd())
	cthCmd.AddCommand(TxCmd())
	err := cthCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func IncorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
