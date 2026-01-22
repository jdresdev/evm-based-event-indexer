package chain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Client interface {
	HeadBlockNumber(ctx context.Context) (*big.Int, error)

	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)

	FilterLogs(ctx context.Context, query LogQuery) ([]types.Log, error)

	ChainID(ctx context.Context) (*big.Int, error)
}

type LogQuery struct {
	FromBlock *big.Int
	ToBlock   *big.Int
	Addresses []common.Address
	Topics    [][]common.Hash
}
