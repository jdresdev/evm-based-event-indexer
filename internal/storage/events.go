package storage

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type EventRepository interface {
	InsertBatch(ctx context.Context, events []Event) error
	DeleteFromBlock(ctx context.Context, blockNumber *big.Int) error
}

type Event struct {
	ChainID     *big.Int
	BlockNumber *big.Int
	BlockHash   common.Hash
	TxHash      common.Hash
	LogIndex    uint
	Contract    common.Address
	EventName   string
	Payload     map[string]any
}
