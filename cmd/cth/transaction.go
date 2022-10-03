package main

import (
	"fmt"
	"os"

	"github.com/elthworth/Cactus-HUB/database"
	"github.com/spf13/cobra"
)

const flagFrom = "from"
const flagTo = "to"
const flagValue = "value"
const flagData = "data"

func TxCmd() *cobra.Command {
	var txCmds = &cobra.Command{
		Use:   "tx",
		Short: "Interact with Transactions (add...).",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {},
	}
	txCmds.AddCommand(TxAddCmd())

	return txCmds
}

func TxAddCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "add",
		Short: "Add new Transaction to DB",
		Run: func(cmd *cobra.Command, args []string) {
			from, _ := cmd.Flags().GetString(flagFrom)
			to, _ := cmd.Flags().GetString(flagTo)
			value, _ := cmd.Flags().GetUint(flagValue)
			data, _ := cmd.Flags().GetString(flagData)

			fromAcc := database.NewAccount(from)
			toAcc := database.NewAccount(to)

			tx := database.NewTx(fromAcc, toAcc, value, data)

			state, err := database.NewStateFromDisk()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			defer state.Close()

			err = state.AddTx(tx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			_, err = state.Persist()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			fmt.Println("TX has successfully added to the Hub")
		},
	}

	cmd.Flags().String(flagFrom, "", "From what account to send tokens")
	cmd.MarkFlagRequired(flagFrom)
	cmd.Flags().String(flagTo, "", "To what account to send tokens")
	cmd.MarkFlagRequired(flagTo)
	cmd.Flags().Uint(flagValue, 0, "How many tokens to send")
	cmd.MarkFlagRequired(flagValue)
	cmd.Flags().String(flagData, "", "Possible values: 'reward'")
	return cmd

}
