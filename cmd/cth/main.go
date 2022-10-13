package main

import (
	"fmt"
	"os"

	"github.com/elthworth/Cactus-HUB/fs"
	"github.com/spf13/cobra"
)

const flagDataDir = "datadir"
const flagIP = "ip"
const flagPort = "port"

func main() {
	var cthCmd = &cobra.Command{
		Use:   "cth",
		Short: "Cactus Hub CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cthCmd.AddCommand(migrateCmd())
	cthCmd.AddCommand(versionCmd)
	cthCmd.AddCommand(runCmd())
	cthCmd.AddCommand(balancesCmd())

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

func getDataDirFromCmd(cmd *cobra.Command) string {
	dataDir, _ := cmd.Flags().GetString(flagDataDir)

	return fs.ExpandPath(dataDir)
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
