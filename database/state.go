package database

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type State struct {
	Balances  map[Account]uint
	TxMemPool []Tx
	DB        *os.File
}

func NewStateFromDisk() (*State, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	genFilePath := filepath.Join(cwd, "database", "genesis.json")
	gen, err := loadGenesis(genFilePath)
	if err != nil {
		return nil, err
	}

	balances := make(map[Account]uint)
	for account, balance := range gen.Balances {
		balances[account] = balance
	}

	txDbFilePath := filepath.Join(cwd, "database", "tx.db")
	f, err := os.OpenFile(txDbFilePath, os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	state := &State{balances, make([]Tx, 0), f}

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		var tx Tx
		json.Unmarshal(scanner.Bytes(), &tx)

		if err := state.Apply(tx); err != nil {
			return nil, err
		}
	}
	return state, nil
}

func (s *State) Add(tx Tx) error {
	if err := s.Apply(tx); err != nil {
		return err
	}

	s.TxMemPool = append(s.TxMemPool, tx)

	return nil
}

func (s *State) Persist() error {
	length := len(s.TxMemPool)
	mempool := make([]Tx, length)
	copy(mempool, s.TxMemPool)

	for i := 0; i < length; i++ {
		txJson, err := json.Marshal(mempool[i])
		if err != nil {
			return err
		}

		if _, err := s.DB.Write(append(txJson, '\n')); err != nil {
			return err
		}

		s.TxMemPool = s.TxMemPool[1:]
	}
	return nil
}

func (s *State) Apply(tx Tx) error {
	if tx.IsReward() {
		s.Balances[tx.To] += tx.Value
		return nil
	}

	if tx.Value > s.Balances[tx.From] {
		return fmt.Errorf("Insufficient Balance")
	}

	s.Balances[tx.From] -= tx.Value
	s.Balances[tx.To] += tx.Value
	return nil
}

func (s *State) Close() {
	s.DB.Close()
}
