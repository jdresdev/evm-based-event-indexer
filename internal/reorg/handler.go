package reorg

import (
	"context"
	"math/big"
)

type Handler interface {
	RollbackTo(ctx context.Context, blockNumber *big.Int) error
}
