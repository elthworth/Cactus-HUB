package main

import (
	"context"
	"fmt"
	"time"

	"github.com/elthworth/Cactus-HUB/database"
	"github.com/elthworth/Cactus-HUB/node"
	"github.com/spf13/cobra"
)

var migrateCmd = func() *cobra.Command {
	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrates the blockchain database according to new business rules.",
		Run: func(cmd *cobra.Command, args []string) {
			miner, _ := cmd.Flags().GetString(flagMiner)
			ip, _ := cmd.Flags().GetString(flagIP)
			port, _ := cmd.Flags().GetUint64(flagPort)

			peer := node.NewPeerNode(
				"127.0.0.1",
				9000,
				true,
				database.NewAccount("elthworth"),
				false,
			)

			n := node.New(getDataDirFromCmd(cmd), ip, port, database.NewAccount(miner), peer)

			n.AddPendingTX(database.NewTx("elthworth", "elthworth", 3, ""), peer)
			n.AddPendingTX(database.NewTx("elthworth", "eroist", 2000, ""), peer)
			n.AddPendingTX(database.NewTx("eroist", "elthworth", 1, ""), peer)
			n.AddPendingTX(database.NewTx("eroist", "taka", 1000, ""), peer)
			n.AddPendingTX(database.NewTx("eroist", "elthworth", 50, ""), peer)

			ctx, closeNode := context.WithTimeout(context.Background(), time.Minute*15)

			go func() {
				ticker := time.NewTicker(time.Second * 10)

				for {
					select {
					case <-ticker.C:
						if !n.LatestBlockHash().IsEmpty() {
							closeNode()
							return
						}
					}
				}
			}()

			err := n.Run(ctx)
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	addDefaultRequiredFlags(migrateCmd)
	migrateCmd.Flags().String(flagMiner, node.DefaultMiner, "miner account of this node to receive block rewards")
	migrateCmd.Flags().String(flagIP, node.DefaultIP, "exposed IP for communication with peers")
	migrateCmd.Flags().Uint64(flagPort, node.DefaultHTTPort, "exposed HTTP port for communication with peers")

	return migrateCmd
}
