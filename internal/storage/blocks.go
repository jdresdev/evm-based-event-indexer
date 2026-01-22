package storage

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type BlockRepository interface {
	Insert(ctx context.Context, block Block) error
	GetHashByNumber(ctx context.Context, number *big.Int) common.Hash
	DeleteFromNumber(ctx context.Context, number *big.Int) error
	GetLastIndexed(ctx context.Context) (*big.Int, error)
	SetLastIndexed(ctx context.Context, number *big.Int) error
}

type Block struct {
	Number     *big.Int
	Hash       common.Hash
	ParentHash common.Hash
}
