package main

import (
	"fmt"
	"os"
	"time"

	"github.com/elthworth/Cactus-HUB/database"
)

func main() {
	state, err := database.NewStateFromDisk()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer state.Close()

	block0 := database.NewBlock(
		database.Hash{},
		uint64(time.Now().Unix()),
		[]database.Tx{
			database.NewTx("elthworth", "elthworth", 10, ""),
			database.NewTx("elthworth", "elthworth", 300, "reward"),
		},
	)

	state.AddBlock(block0)
	block0hash, _ := state.Persist()

	block1 := database.NewBlock(
		block0hash,
		uint64(time.Now().Unix()),
		[]database.Tx{
			database.NewTx("elthworth", "elroist", 2000, ""),
			database.NewTx("elthworth", "elthworth", 100, "reward"),
			database.NewTx("elroist", "elthworth", 1, ""),
			database.NewTx("elroist", "taka", 1000, ""),
			database.NewTx("elroist", "elthworth", 50, ""),
			database.NewTx("elthworth", "elthworth", 600, "reward"),
		},
	)
	state.AddBlock(block1)
	state.Persist()
}
