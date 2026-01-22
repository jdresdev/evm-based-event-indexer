package config

import "math/big"

type ChainConfig struct {
	Name          string
	Confirmations uint64
	BatchSize     uint64
	StartBlock    *big.Int
}
