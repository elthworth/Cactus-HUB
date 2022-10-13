package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const flagDataDir = "datadir"

func main() {
	var cthCmd = &cobra.Command{
		Use:   "cth",
		Short: "Cactus HUB CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cthCmd.AddCommand(versionCmd)
	cthCmd.AddCommand(BalancesCmd())
	cthCmd.AddCommand(runCmd())
	err := cthCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addDefaultRequiredFlags(cmd *cobra.Command) {
	cmd.Flags().String(flagDataDir, "", "Absolute path to the node data dir where the DB will/is stored")
	cmd.MarkFlagRequired(flagDataDir)
}

func IncorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
