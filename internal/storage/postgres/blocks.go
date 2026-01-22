package postgres

import (
	"context"
	"database/sql"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jdresdev/evm-based-event-indexer/internal/storage"
)

type BlockRepository struct {
	db *sql.DB
}

func NewBlockRepository(db *sql.DB) *BlockRepository {
	return &BlockRepository{db: db}
}

func (r *BlockRepository) Insert(ctx context.Context, block storage.Block) error {
	exec := r.db.ExecContext
	if tx := txFromContext(ctx); tx != nil {
		exec = tx.ExecContext
	}

	_, err := exec(ctx, `
		INSERT INTO blocks (number, hash, parent_hash)
		VALUES ($1, $2, $3)
		ON CONFLICT (number) DO NOTHING
	`,
		block.Number.Int64(),
		block.Hash.Hex(),
		block.ParentHash.Hex(),
	)

	return err
}
