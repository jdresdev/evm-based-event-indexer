package decoder

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Decoder interface {
	CanDecode(address common.Address, topic common.Hash) bool
	Decode(log types.Log) (DecodedEvent, error)
}

type DecodedEvent struct {
	EventName string
	Payload   map[string]any
}
